package app

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/isauran/go-pkg/logger"
	"github.com/isauran/gokit-microservice-profile/internal/config"
	customerService "github.com/isauran/gokit-microservice-profile/internal/service/customer"
	profileService "github.com/isauran/gokit-microservice-profile/internal/service/profile"
)

type App struct {
	log          *slog.Logger
	gokitLogger  log.Logger
	handler      http.Handler
	serverConfig config.ServerConfig
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{
		log:         logger.SlogJSONLogger(os.Stdout, slog.LevelDebug),
		gokitLogger: logger.GoKitJSONLogger(os.Stderr),
	}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	a.runHTTPServer()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initEnvironment,
		a.initHTTPHandler,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initEnvironment(_ context.Context) error {
	if !config.EnvExists("APP.SERVER.PORT") {
		err := config.EnvLoad(".env")
		if err != nil {
			return err
		}
	}
	a.log.Info("app", "environment", config.EnvInfo())
	return nil
}

func (a *App) initHTTPHandler(_ context.Context) error {

	if a.serverConfig == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			a.log.Error("failed to get server config", "error", err.Error())
			os.Exit(1)
		}
		a.serverConfig = cfg
	}

	r := mux.NewRouter()

	var profileSvc profileService.Service
	profileSvc = profileService.NewInmemService()
	profileSvc = profileService.LoggingMiddleware(a.gokitLogger)(profileSvc)
	a.handler = profileService.MakeHTTPHandler(r, profileSvc, log.With(a.gokitLogger, "component", "HTTP"))

	var customerSvc customerService.Service
	customerSvc = customerService.NewInmemService()
	customerSvc = customerService.LoggingMiddleware(a.gokitLogger)(customerSvc)
	a.handler = customerService.MakeHTTPHandler(r, customerSvc, log.With(a.gokitLogger, "component", "HTTP"))

	return nil
}

func (a *App) runHTTPServer() {
	var (
		httpAddr = flag.String("http.addr", a.serverConfig.ServerPort(), "HTTP listen address")
	)
	flag.Parse()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		a.gokitLogger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, a.handler)
	}()

	a.gokitLogger.Log("exit", <-errs)
}
