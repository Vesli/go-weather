package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func registerRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Main route. Visit /weather/daily/{cityName} to visualise data."))
	})
	r.Route("/weather", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Nothing to display here. Want to go to daily/{cityName} to visualise data?"))
		})

		r.Route("/daily", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Please precise a city name"))
			})

			/* Logic start here */
			r.Get("/:city", getWeeklyWeather)
		})
	})
	return r
}
