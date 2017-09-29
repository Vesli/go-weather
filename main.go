/* TODO don't flag as global => *config */

package main

import (
	"flag"
	"log"

	"github.com/go-weather/service"
)

func main() {
	pathToConfig := flag.String("config", "config.json", "Config file path")
	flag.Parse()

	service, err := service.NewService(*pathToConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()

	service.Run()
}
