package main

import (
	"log"
	"net/http"

	"github.com/yudhasubki/forecast/config"
	"github.com/yudhasubki/forecast/forecasting"
	"github.com/yudhasubki/forecast/router"
)

func main() {
	cfg := config.Get()
	repository := forecasting.NewRepository(*cfg)
	service := forecasting.NewService(repository)
	handler := forecasting.NewHandler(service)
	routerHandler := router.Router{ForecastingHandler: handler}
	router := router.NewRouter(routerHandler)

	port := cfg.Port
	if port == "" {
		port = "80"
	}
	log.Println("Running on Port : " + port)
	http.ListenAndServe(":"+port, router)
}
