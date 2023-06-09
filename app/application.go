package app

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sanitize/sanitize"
	"github.com/microcosm-cc/bluemonday"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/johannes-kuhfuss/aoip-svc/config"
	"github.com/johannes-kuhfuss/aoip-svc/handler"
	"github.com/johannes-kuhfuss/aoip-svc/repository"
	"github.com/johannes-kuhfuss/aoip-svc/service"
	"github.com/johannes-kuhfuss/services_utils/date"
	"github.com/johannes-kuhfuss/services_utils/logger"
)

var (
	cfg           config.AppConfig
	server        http.Server
	appEnd        chan os.Signal
	ctx           context.Context
	cancel        context.CancelFunc
	deviceRepo    repository.DeviceRepositoryMem
	deviceSvc     service.DefaultDeviceService
	deviceHandler handler.DeviceHandler
)

func StartApp() {
	logger.Info("Starting application")

	getCmdLine()
	err := config.InitConfig(config.EnvFile, &cfg)
	if err != nil {
		panic(err)
	}
	initRouter()
	initServer()
	initMetrics()
	wireApp()
	mapUrls()
	registerForOsSignals()
	createSanitizers()

	go startDeviceDiscovery()
	go startServer()

	<-appEnd
	cleanUp()

	if srvErr := server.Shutdown(ctx); err != nil {
		logger.Error("Graceful shutdown failed", srvErr)
	} else {
		logger.Info("Graceful shutdown finished")
	}
}

func getCmdLine() {
	flag.StringVar(&config.EnvFile, "config.file", ".env", "Specify location of config file. Default is .env")
	flag.Parse()
}

func initRouter() {
	gin.SetMode(cfg.Gin.Mode)
	gin.DefaultWriter = logger.GetLogger()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(AddRequestId())
	router.SetTrustedProxies(nil)
	globPath := filepath.Join(cfg.Gin.TemplatePath, "*.tmpl")
	router.LoadHTMLGlob(globPath)

	cfg.RunTime.Router = router
}

func initServer() {
	var tlsConfig tls.Config

	if cfg.Server.UseTls {
		tlsConfig = tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		}
	}
	if cfg.Server.UseTls {
		cfg.RunTime.ListenAddr = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.TlsPort)
	} else {
		cfg.RunTime.ListenAddr = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	}

	server = http.Server{
		Addr:              cfg.RunTime.ListenAddr,
		Handler:           cfg.RunTime.Router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    0,
	}
	if cfg.Server.UseTls {
		server.TLSConfig = &tlsConfig
		server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler))
	}
}

func initMetrics() {
	prometheusRegister()
}

func wireApp() {
	deviceRepo = repository.NewDeviceRepositoryMem(&cfg)
	deviceSvc = service.NewDeviceService(&cfg, &deviceRepo)
	deviceHandler = handler.NewDeviceHandler(&cfg, deviceSvc)
}

func mapUrls() {
	api := cfg.RunTime.Router.Group("/devices", validateAuth(), prometheusMetrics())
	{
		api.GET("/", deviceHandler.GetAllDevices)
		api.GET("/:device_id", nil)
	}
	ui := cfg.RunTime.Router.Group("/")
	{
		ui.GET("/about", nil)
	}
	cfg.RunTime.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func registerForOsSignals() {
	appEnd = make(chan os.Signal, 1)
	signal.Notify(appEnd, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}

func createSanitizers() {
	sani, err := sanitize.New()
	if err != nil {
		logger.Error("Error creating sanitizer", err)
		panic(err)
	}
	cfg.RunTime.Sani = sani
	cfg.RunTime.BmPolicy = bluemonday.UGCPolicy()
}

func startServer() {
	logger.Info(fmt.Sprintf("Listening on %v", cfg.RunTime.ListenAddr))
	cfg.RunTime.StartDate = date.GetNowUtc()
	if cfg.Server.UseTls {
		if err := server.ListenAndServeTLS(cfg.Server.CertFile, cfg.Server.KeyFile); err != nil && err != http.ErrServerClosed {
			logger.Error("Error while starting https server", err)
			panic(err)
		}
	} else {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error while starting http server", err)
			panic(err)
		}
	}
}

func startDeviceDiscovery() {
	cfg.RunTime.RunDiscover = true
	deviceSvc.Run()
}

func cleanUp() {
	shutdownTime := time.Duration(cfg.Server.GracefulShutdownTime) * time.Second
	ctx, cancel = context.WithTimeout(context.Background(), shutdownTime)
	defer func() {
		logger.Info("Cleaning up")
		cfg.RunTime.RunDiscover = false
		logger.Info("Done cleaning up")
		cancel()
	}()
}
