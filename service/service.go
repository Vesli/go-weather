package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-weather/config"
	mgo "gopkg.in/mgo.v2"
)

type Service struct {
	Config  *config.Config
	Session *mgo.Session
	router  *chi.Mux
}

func NewService(pathToConfig string) (*Service, error) {
	conf, err := config.LoadConfig(pathToConfig)
	if err != nil {
		return nil, err
	}

	session, err := mgo.Dial(conf.DbURL)
	if err != nil {
		return nil, err
	}

	service := &Service{
		Config:  conf,
		Session: session,
		router:  chi.NewRouter(),
	}

	ret := initMiddleware(service)
	service.router.Use(ret)

	service.router.Mount("/progimage", registerRoutes(service))

	return service, nil
}

func (s *Service) Run() {
	fmt.Println("Project started correctly and now running on port:", s.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Port), s.router))
}

func (s *Service) Close() {
	s.Session.Close()
}
