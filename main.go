/* TODO don't flag as global => *config */

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/pressly/chi"
)

func main() {
	pathToConfig := flag.String("config", "config.json", "Config file path")
	flag.Parse()

	service, err := newService(*pathToConfig)
	if err != nil {
		panic(err)
	}
	defer service.Close()

	r := chi.NewRouter()
	ret := initMiddleware(service)
	r.Use(ret)

	/* Mount in anticipation of wild routing */
	r.Mount("/", registerRoutes())

	http.ListenAndServe(fmt.Sprintf(":%d", service.Config.Port), r)
}
