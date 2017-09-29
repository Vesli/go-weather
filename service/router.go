package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

func registerRoutes(service *Service) *chi.Mux {
	service.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Main route. Visit /weather/daily/{cityName} to visualise data."))
	})
	service.router.Route("/weather", func(r chi.Router) {
		service.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Nothing to display here. Want to go to daily/{cityName} to visualise data?"))
		})

		service.router.Route("/daily", func(r chi.Router) {
			service.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Please precise a city name"))
			})

			/* Logic start here */
			service.router.Get("/:city", service.GetWeeklyWeather)
		})
	})
	return service.router
}
