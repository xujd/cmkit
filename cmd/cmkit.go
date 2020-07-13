package main

import (
	"context"
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
	"cmkit/pkg/db"
	"cmkit/pkg/hello"

	_ "github.com/lib/pq"
)

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "vdcdb"
)

func main() {
	fmt.Printf("v%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION)
	var (
		serviceHost = flag.String("service.host", "127.0.0.1", "service ip address")
		servicePort = flag.String("service.port", "8089", "service port")
	)
	flag.Parse()

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	hostPort := *serviceHost + ":" + *servicePort
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestamp)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

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

	// database
	dbLogger := log.With(logger, "component", "db")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	pool := db.Init(psqlInfo, dbLogger)
	defer pool.Close()
	db.Ping(ctx, dbLogger)
	db.Query(ctx, dbLogger)

	// service
	var authSvc auth.Service
	authSvc = auth.AuthService{}

	authSvc = auth.NewLoggingMiddleware(log.With(logger, "component", "auth"), authSvc)
	authSvc = auth.NewInstrumentingMiddleware(requestCount, requestLatency, authSvc)
	loginEndpoint := auth.MakeLoginEndpoint(authSvc)
	renewEndpoint := auth.MakeRenewEndpoint(authSvc)
	renewEndpoint = kitjwt.NewParser(auth.JwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(renewEndpoint)

	authEndpoints := auth.AuthEndpoints{
		LoginEndpoint: loginEndpoint,
		RenewEndpoint: renewEndpoint,
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
