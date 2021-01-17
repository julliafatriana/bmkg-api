package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yudhasubki/forecast/forecasting"
)

type Router struct {
	ForecastingHandler *forecasting.Handler
}

func NewRouter(router Router) *mux.Router {
	r := mux.NewRouter()
	handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/health", router.ForecastingHandler.HealthCheck).Methods("GET")
	r.HandleFunc("/v1/forecasting/province", router.ForecastingHandler.ResolveProvinces).Methods("GET")
	r.HandleFunc("/v1/forecasting/area", router.ForecastingHandler.ResolveAreas).Methods("POST")
	r.HandleFunc("/v1/forecasting/search/province", router.ForecastingHandler.ResolveForecastingByProvince).Methods("POST")
	r.HandleFunc("/v1/forecasting/search/area", router.ForecastingHandler.ResolveForecastingByProvinceAndArea).Methods("POST")

	return r
}
