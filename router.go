package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func registerRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Main route. Visit /weather/weekly/{cityName} to see returned data."))
	})
	r.Route("/weather", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Nothing to display here. Want to go to weekly/{cityName} to see returned data?"))
		})

		r.Route("/weekly", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Please precise a city name"))
			})
			r.Get("/:city", getWeeklyWeather)
		})
	})
}
