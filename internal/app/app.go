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

	gokitlog "github.com/go-kit/log"
	"github.com/isauran/gokit-microservice-profile/internal/config"
	"github.com/isauran/gokit-microservice-profile/internal/service"
	"github.com/isauran/gokitlogger"
	"github.com/isauran/slogger"
)

type App struct {
	log          *slog.Logger
	gokitLogger  gokitlog.Logger
	handler      http.Handler
	serverConfig config.ServerConfig
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
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

func (a *App) initConfig(_ context.Context) error {

	a.log = slogger.NewLogger(os.Stdout, slogger.WithJSON(true))
	a.gokitLogger = gokitlogger.NewLogger(os.Stdout, gokitlogger.WithJSON(true))

	if config.CountWithPrefix("app.") == 0 {

		a.log = slogger.NewLogger(os.Stdout)
		a.gokitLogger = gokitlogger.NewLogger(os.Stdout)

		err := config.Load(".env")
		if err != nil {
			return err
		}
	}

	if a.serverConfig == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			a.log.Error("failed to get server config", "error", err.Error())
			os.Exit(1)
		}
		a.serverConfig = cfg
	}

	a.log.Info("app", "environment", config.GetWithPrefix("app.", "password", "secret"))
	return nil
}

func (a *App) initHTTPHandler(_ context.Context) error {
	a.handler = service.MakeHTTPHandler(a.gokitLogger)
	return nil
}

func (a *App) runHTTPServer() error {
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

	return nil
}
