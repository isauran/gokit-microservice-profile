package service

import (
	"net/http"

	gokitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
	customerService "github.com/isauran/gokit-microservice-profile/internal/service/customer"
	profileService "github.com/isauran/gokit-microservice-profile/internal/service/profile"
)

func MakeHTTPHandler(gokitlogger gokitlog.Logger) http.Handler {
	r := mux.NewRouter()

	var profileSvc profileService.Service
	profileSvc = profileService.NewInmemService()
	profileSvc = profileService.LoggingMiddleware(gokitlogger)(profileSvc)
	profileService.MakeHTTPHandler(r, profileSvc, gokitlog.With(gokitlogger, "component", "HTTP"))

	var customerSvc customerService.Service
	customerSvc = customerService.NewInmemService()
	customerSvc = customerService.LoggingMiddleware(gokitlogger)(customerSvc)
	customerService.MakeHTTPHandler(r, customerSvc, gokitlog.With(gokitlogger, "component", "HTTP"))

	return r
}
