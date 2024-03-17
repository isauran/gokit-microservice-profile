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

	"github.com/isauran/gokit-microservice-profile/internal/config"
	"github.com/isauran/gokit-microservice-profile/internal/app/provider"
)

type App struct {
	log             *slog.Logger
	serviceProvider *provider.ServiceProvider
}

func NewApp(ctx context.Context, l *slog.Logger) (*App, error) {
	a := &App{log: l}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	return a.ListenAndServe()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initEnvironment,
		a.initServiceProvider,
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

	if config.CountWithPrefix("app.") == 0 {

		err := config.Load(".env")
		if err != nil {
			return err
		}
	}

	a.log.Info("app", "environment", config.GetWithPrefix("app.", "password", "secret"))
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = provider.NewServiceProvider(a.log)
	return nil
}

func (a *App) ListenAndServe() error {
	var (
		httpAddr = flag.String("http.addr", a.serviceProvider.ServerConfig().ServerPort(), "HTTP listen address")
	)
	flag.Parse()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		a.log.Info("server", "transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, a.serviceProvider.HTTPHandler())
	}()

	a.log.Error("server", "exit", <-errs)
	return <-errs
}
