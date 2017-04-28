package main

/*TODO
config file
flags
*/
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/pressly/chi"
	//"github.com/pressly/chi/middleware"

	"net/http"

	"gopkg.in/mgo.v2"
)

/*
	This is a simple test for flags. Now I see how it works
	and will change it to proper external communication
*/
type config struct {
	Url   string
	Units string
	AppID string
	DbUrl string
	Port  int
}

// var flagConfig = flag.String("config", "", "Yolo t'as fait un flag poto")
// var toto = flag.Bool("test", false, "GG BOOL")

var pathToConfig = flag.String("config", "configuration/weatherConfig.json", "Config file path")

var conf *config

func loadConfig(pathToConfig string) (*config, error) {
	data, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		return nil, fmt.Errorf("File config error: %s", err)
	}

	var conf config
	if err = json.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func main() {
	flag.Parse()

	var err error
	conf, err = loadConfig(*pathToConfig)
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial(conf.DbUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	r := chi.NewRouter()
	/* export middleware to other functions if plural */
	r.Use(SessionMiddleware(session))

	/* Mount in anticipation of wild routing */
	r.Mount("/", registerRoutes())

	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}

func SessionMiddleware(session *mgo.Session) func(next http.Handler) http.Handler {
	//copy session here
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionC := session.Copy()
			defer sessionC.Close()
			ctx := context.WithValue(r.Context(), "MGOsession", sessionC)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WriteJson(w http.ResponseWriter, obj interface{}) {
	jData, err := json.Marshal(obj)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func getWeeklyWeather(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	sessionC := r.Context().Value("MGOsession").(*mgo.Session)
	c := sessionC.DB("goweather").C("dailyweather")

	cw, err := GetCityWeather(city, c)
	if err != nil {
		WriteJson(w, err)
		return
	}
	WriteJson(w, cw)
}
