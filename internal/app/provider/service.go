package provider

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/isauran/gokit-microservice-profile/internal/config"
	"github.com/isauran/gokit-microservice-profile/internal/service/customer"
	"github.com/isauran/gokit-microservice-profile/internal/service/profile"
)

type ServiceProvider struct {
	log       *slog.Logger
	serverCfg config.ServerConfig
	handler   http.Handler
}

func NewServiceProvider(l *slog.Logger) *ServiceProvider {
	return &ServiceProvider{log: l}
}

func (s *ServiceProvider) ServerConfig() config.ServerConfig {
	if s.serverCfg == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			s.log.Error("failed to get server config", "error", err.Error())
			os.Exit(1)
		}
		s.serverCfg = cfg
	}
	return s.serverCfg
}

func (s *ServiceProvider) HTTPHandler() http.Handler {
	if s.handler == nil {
		r := mux.NewRouter()

		var profileSvc profile.Service
		profileSvc = profile.NewInmemService()
		profileSvc = profile.LoggingMiddleware(s.log)(profileSvc)
		profile.MakeHTTPHandler(r, profileSvc)

		var customerSvc customer.Service
		customerSvc = customer.NewInmemService()
		customerSvc = customer.LoggingMiddleware(s.log)(customerSvc)
		customer.MakeHTTPHandler(r, customerSvc)

		s.handler = r
	}
	return s.handler
}
