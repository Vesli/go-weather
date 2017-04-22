package main

/*TODO
config file
flags
*/
import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pressly/chi"
	//"github.com/pressly/chi/middleware"
	"flag"
	"net/http"

	"gopkg.in/mgo.v2"
)

var flagConfig = flag.String("config", "", "Yolo t'as fait un flag poto")
var toto = flag.Bool("test", false, "GG BOOL")

func main() {
	flag.Parse()
	fmt.Println(*toto, "lol", *flagConfig)
	r := chi.NewRouter()

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	r.Use(SessionMiddleware(session))
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

	http.ListenAndServe(":8080", r)
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
