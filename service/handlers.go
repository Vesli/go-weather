package service

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-weather/config"
	"github.com/go-weather/weatherCore"

	mgo "gopkg.in/mgo.v2"
)

/*
	TODO create const for key context + collection name
	'DB NAME' in config file
*/

const weatherCollection = "dailyweather"

func (s *Service) GetWeeklyWeather(w http.ResponseWriter, r *http.Request) {
	conf := r.Context().Value(configuration).(*config.Config)
	city := strings.Title(chi.URLParam(r, "city"))

	sessionC := r.Context().Value(mgoSession).(*mgo.Session)
	c := sessionC.DB(conf.DBName).C(weatherCollection)

	cw, err := weatherCore.GetCityWeather(city, c, conf)
	if err != nil {
		writeJSON(w, err)
		return
	}
	writeJSON(w, cw)
}
