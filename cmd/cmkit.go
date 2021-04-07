package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	"cmkit/pkg/auth"
	"cmkit/pkg/fileupload"
	"cmkit/pkg/hello"
	"cmkit/pkg/home"
	"cmkit/pkg/res"
	"cmkit/pkg/sys"
	"cmkit/pkg/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

func main() {
	fmt.Printf("v%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)
	var (
		serviceHost = flag.String("service.host", "", "service ip address")
		servicePort = flag.String("service.port", "8089", "service port")
		dbHost      = flag.String("db.host", "127.0.0.1", "db ip address")
		dbPort      = flag.Int("db.port", 5432, "db port")
		dbUser      = flag.String("db.user", "postgres", "db user")
		dbPasswd    = flag.String("db.passwd", "123456", "db password")
		dbName      = flag.String("db.name", "cmkit", "db name")
		webDir      = flag.String("web.dir", "../web/public/assets/pictures/", "web dir")
	)
	flag.Parse()

	hostPort := *serviceHost + ":" + *servicePort
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestamp)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		*dbHost, *dbPort, *dbUser, *dbPasswd, *dbName)
	// logger.Log("database string", psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		logger.Log("database error", err)
		return
	}
	defer db.Close()

	// instrumenting
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "CMKIT",
		Subsystem: "default",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "CMKIT",
		Subsystem: "default",
		Name:      "request_latency",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// service
	// 权限
	var authSvc auth.Service
	authSvc = auth.AuthService{
		DB: db,
	}

	authSvc = auth.NewLoggingMiddleware(log.With(logger, "component", "auth"), authSvc)
	authSvc = auth.NewInstrumentingMiddleware(requestCount, requestLatency, authSvc)
	authEndpoints := auth.CreateEndpoints(authSvc)

	// 系统基础
	var sysSvc sys.Service
	sysSvc = sys.SysService{
		DB:     db,
		WebDir: *webDir,
	}

	sysSvc = sys.NewLoggingMiddleware(log.With(logger, "component", "sys"), sysSvc)
	sysSvc = sys.NewInstrumentingMiddleware(requestCount, requestLatency, sysSvc)
	sysEndpoints := sys.CreateEndpoints(sysSvc)

	// 资源管理
	var resSvc res.Service
	resSvc = res.ResService{
		DB: db,
	}

	resSvc = res.NewLoggingMiddleware(log.With(logger, "component", "res"), resSvc)
	resSvc = res.NewInstrumentingMiddleware(requestCount, requestLatency, resSvc)
	resEndpoints := res.CreateEndpoints(resSvc)

	// 文件上传
	var fileuploadSvc fileupload.Service
	fileuploadSvc = fileupload.FileUploadService{
		WebDir: *webDir,
	}

	fileuploadSvc = fileupload.NewLoggingMiddleware(log.With(logger, "component", "fileupload"), fileuploadSvc)
	fileuploadSvc = fileupload.NewInstrumentingMiddleware(requestCount, requestLatency, fileuploadSvc)
	fileUploadEndpoints := fileupload.CreateEndpoints(fileuploadSvc)

	// 首页
	var homeSvc home.Service
	homeSvc = home.HomeService{
		DB: db,
	}

	homeSvc = home.NewLoggingMiddleware(log.With(logger, "component", "home"), homeSvc)
	homeSvc = home.NewInstrumentingMiddleware(requestCount, requestLatency, homeSvc)
	homeEndpoints := home.CreateEndpoints(homeSvc)

	// 测试
	var helloSvc hello.Service
	helloSvc = hello.HelloService{}

	helloSvc = hello.NewLoggingMiddleware(log.With(logger, "component", "hello"), helloSvc)
	helloSvc = hello.NewInstrumentingMiddleware(requestCount, requestLatency, helloSvc)

	helloEndpoint := hello.MakeHelloEndpoint(helloSvc)
	helloEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(helloEndpoint)

	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()

	mux.Handle("/auth/", auth.MakeHandler(authEndpoints, httpLogger))
	mux.Handle("/hello/", hello.MakeHandler(helloEndpoint, httpLogger))
	mux.Handle("/sys/", sys.MakeHandler(sysEndpoints, httpLogger))
	mux.Handle("/res/", res.MakeHandler(resEndpoints, httpLogger))
	mux.Handle("/home/", home.MakeHandler(homeEndpoints, httpLogger))
	http.Handle("/file/", fileupload.MakeHandler(fileUploadEndpoints, httpLogger))
	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", hostPort, "msg", "listening")
		errs <- http.ListenAndServe(hostPort, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func upload(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			return
		}

		r.ParseMultipartForm(32 << 20) // max memory is set to 32MB
		clientfd, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			msg, _ := json.Marshal(utils.CmkitResponse{Code: 40003, Success: false, Message: "upload failed."})
			w.Write(msg)
			return
		}
		defer clientfd.Close()
		destLocalPath := "./files/"
		localpath := fmt.Sprintf("%s%s", destLocalPath, handler.Filename)
		localfd, err := os.OpenFile(localpath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			msg, _ := json.Marshal(utils.CmkitResponse{Code: 40003, Success: false, Message: "upload failed."})
			w.Write(msg)
			return
		}
		defer localfd.Close()

		// 利用io.TeeReader在读取文件内容时计算hash值
		fhash := sha1.New()
		io.Copy(localfd, io.TeeReader(clientfd, fhash))
		hstr := hex.EncodeToString(fhash.Sum(nil))
		msg, _ := json.Marshal(utils.CmkitResponse{Code: 20000, Success: true, Message: "upload finish.", Data: hstr})
		w.Write(msg)

		h.ServeHTTP(w, r)
	})
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
