package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Weather struct {
	Main        string
	Description string
	Icon        string
}

type MainWeather struct {
	Temp     float32
	Humidity float32
	Temp_min float32
	Temp_max float32
}

type CityWeather struct {
	Name       string
	LastUpdate time.Time
	Weather    []Weather
	Main       MainWeather
}

func GetWeatherFromDB(city string, c *mgo.Collection) (*CityWeather, error) {
	w := new(CityWeather)
	err := c.Find(bson.M{"name": city}).One(&w)

	if err != nil {
		return nil, err
	}

	now := time.Now()
	diff := now.Sub(w.LastUpdate)

	if diff.Hours() > 1 {
		w, err = GetCityWeatherFromApi(city, c)
		if err != nil {
			return nil, err
		}
	}
	return w, nil
}

func UpdateWeatherInDB(w *CityWeather, c *mgo.Collection) error {
	w.LastUpdate = time.Now()
	_, err := c.Upsert(bson.M{"name": w.Name}, w)
	return err
}

func GetCityWeatherFromApi(city string, c *mgo.Collection) (*CityWeather, error) {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=a3f80eebf75efbe3e8ce1c1192c5d48f"

	retUrl, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer retUrl.Body.Close()

	body, err := ioutil.ReadAll(retUrl.Body)
	if err != nil {
		return nil, err
	}

	w := &CityWeather{}
	if err = json.Unmarshal(body, w); err != nil {
		return nil, err
	}

	err = UpdateWeatherInDB(w, c)
	if err != nil {
		fmt.Println("An error occured with the DB: ", err)
	}

	return w, nil
}

func GetCityWeather(city string, c *mgo.Collection) (*CityWeather, error) {
	w, err := GetWeatherFromDB(city, c)
	if err != nil {
		log.Println(err)
	} else {
		return w, nil
	}

	return GetCityWeatherFromApi(city, c)
}