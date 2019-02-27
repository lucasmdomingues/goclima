package main

import (
	"goclima/app/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/locale", handlers.LocaleHandler).Methods("GET")
	r.HandleFunc("/locale/city", handlers.LocaleHandler).Methods("GET")
	r.HandleFunc("/climate", handlers.ClimateHandler).Methods("GET")
	r.HandleFunc("/weather", handlers.WeatherHandler).Methods("GET")

	http.ListenAndServe(":"+port, r)

}
