package main

import (
	"flag"
	"fmt"
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
	"cmkit/pkg/hello"

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
		serviceHost = flag.String("service.host", "127.0.0.1", "service ip address")
		servicePort = flag.String("service.port", "8089", "service port")
		dbHost      = flag.String("db.host", "127.0.0.1", "db ip address")
		dbPort      = flag.Int("db.port", 5432, "db port")
		dbUser      = flag.String("db.user", "postgres", "db user")
		dbPasswd    = flag.String("db.passwd", "123456", "db password")
		dbName      = flag.String("db.name", "cmkit", "db name")
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
	logger.Log("database string", psqlInfo)
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
	var authSvc auth.Service
	authSvc = auth.AuthService{
		DbHandler: db,
	}

	authSvc = auth.NewLoggingMiddleware(log.With(logger, "component", "auth"), authSvc)
	authSvc = auth.NewInstrumentingMiddleware(requestCount, requestLatency, authSvc)
	loginEndpoint := auth.MakeLoginEndpoint(authSvc)
	renewvalEndpoint := auth.MakeRenewvalEndpoint(authSvc)
	renewvalEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(renewvalEndpoint)

	authEndpoints := auth.AuthEndpoints{
		LoginEndpoint:    loginEndpoint,
		RenewvalEndpoint: renewvalEndpoint,
	}

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
