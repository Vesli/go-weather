/* TODO mock API */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/vesli/go-weather/config"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type weather struct {
	Main        string `json:"main" bson:"main"`
	Description string `json:"description" bson:"description"`
	Icon        string `json:"icon" bson:"icon"`
}

type mainWeather struct {
	Temp     float32 `json:"temp" bson:"temp"`
	Humidity float32 `json:"humidity" bson:"humidity"`
	TempMin  float32 `json:"temp_min" bson:"temp_min"`
	TempMax  float32 `json:"temp_max" bson:"temp_max"`
}

type cityWeather struct {
	Name       string      `json:"name" bson:"name"`
	LastUpdate time.Time   `json:"last_update" bson:"last_update"`
	Weather    []weather   `json:"weather" bson:"weather"`
	Main       mainWeather `json:"main" bson:"main"`
}

var errOutDated = errors.New("Update weather from API")

func updateWeatherInDB(w *cityWeather, c *mgo.Collection) error {
	w.LastUpdate = time.Now()
	_, err := c.Upsert(bson.M{"name": w.Name}, w)
	return err
}

func getWeatherFromDB(city string, c *mgo.Collection, conf *config.Config) (*cityWeather, error) {
	w := new(cityWeather)
	err := c.Find(bson.M{"name": city}).One(w)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	diff := now.Sub(w.LastUpdate)

	if diff.Hours() > 1 {
		return nil, errOutDated
	}
	return w, nil
}

func getCityWeatherFromAPI(city string, c *mgo.Collection, conf *config.Config) (*cityWeather, error) {
	urlParams := make(url.Values)
	urlParams.Add("q", city)
	urlParams.Add("units", conf.Units)
	urlParams.Add("appid", conf.AppID)

	retURL, err := http.Get(fmt.Sprintf("%s?%s", conf.URL, urlParams.Encode()))
	if err != nil {
		return nil, err
	}
	defer retURL.Body.Close()

	body, err := ioutil.ReadAll(retURL.Body)
	if err != nil {
		return nil, err
	}

	w := &cityWeather{}
	if err = json.Unmarshal(body, w); err != nil {
		return nil, err
	}

	err = updateWeatherInDB(w, c)
	if err != nil {
		fmt.Println("An error occured with the DB: ", err)
	}

	return w, nil
}

func getCityWeather(city string, c *mgo.Collection, conf *config.Config) (*cityWeather, error) {
	w, err := getWeatherFromDB(city, c, conf)
	if err == nil {
		return w, nil
	}
	log.Println(err)
	return getCityWeatherFromAPI(city, c, conf)
}
