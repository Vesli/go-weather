package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func registerRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Main route. Visit /weather/weekly/{cityName} to visualise data."))
	})
	r.Route("/weather", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Nothing to display here. Want to go to weekly/{cityName} to visualise data?"))
		})

		r.Route("/weekly", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Please precise a city name"))
			})
			r.Get("/:city", getWeeklyWeather)
		})
	})
	return r
}
