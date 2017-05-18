package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vesli/go-weather/config"
	"gopkg.in/mgo.v2"

	"testing"
)

var (
	conf          *config.Config
	dbSession     *mgo.Session
	dailyWeatherC *mgo.Collection
)

func TestGoWeather(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWeather Suite")
}

var _ = BeforeSuite(func() {
	var err error
	conf, err = config.LoadConfig("config-test.json")
	Expect(err).NotTo(HaveOccurred())

	dbSession, err = mgo.Dial(conf.DbURL)
	Expect(err).NotTo(HaveOccurred())

	dailyWeatherC = dbSession.DB("goweather").C("dailyweather")
	Expect(dailyWeatherC).NotTo(BeNil())

})

var _ = AfterSuite(func() {
	fmt.Println("Closing!")
	dbSession.Close()
})
