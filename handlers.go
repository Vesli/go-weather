package main

import (
	"net/http"
	"strings"

	"github.com/pressly/chi"
	"github.com/vesli/go-weather/config"
	mgo "gopkg.in/mgo.v2"
)

/*
	TODO create const for key context + collection name
	'DB NAME' in config file
*/

const weatherCollection = "dailyweather"

func getWeeklyWeather(w http.ResponseWriter, r *http.Request) {
	conf := r.Context().Value(configuration).(*config.Config)
	city := strings.Title(chi.URLParam(r, "city"))

	sessionC := r.Context().Value(mgoSession).(*mgo.Session)
	c := sessionC.DB(conf.DBName).C(weatherCollection)

	cw, err := getCityWeather(city, c, conf)
	if err != nil {
		writeJSON(w, err)
		return
	}
	writeJSON(w, cw)
}
