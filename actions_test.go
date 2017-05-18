package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Actions", func() {
	Describe("Test get city weather function", func() {
		Context("When the DB is empty", func() {
			It("Should call the API", func() {
				w, err := getWeatherFromDB("rome", dailyWeatherC, confTest)
				Expect(err).To(HaveOccurred())
				Expect(w).To(BeNil())

				w, err = getCityWeatherFromAPI("rome", dailyWeatherC, confTest)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).NotTo(BeNil())
				Expect(w.Name).To(Equal("Rome"))

				w, err = getCityWeather("rome", dailyWeatherC, confTest)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).NotTo(BeNil())
			})
		})
		Context("When the entity is in the DB", func() {
			It("Should return from DB", func() {
				w, err := getWeatherFromDB("Rome", dailyWeatherC, confTest)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).NotTo(BeNil())
				Expect(w.Name).To(Equal("Rome"))

				w, err = getCityWeather("Rome", dailyWeatherC, confTest)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).NotTo(BeNil())
			})

			/*
				This is the only place to copy session.
				London should NOT be updated in DB, to keep the test correct.
			*/
			It("Should update from API after one hour", func() {
				dbSessionCopy := dbSession.Copy()
				defer dbSessionCopy.Close()

				dailyWeatherC = dbSessionCopy.DB(confTest.DBName).C("permanent-dailyweather")
				w, err := getWeatherFromDB("London", dailyWeatherC, confTest)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(errOutDated))
				Expect(w).To(BeNil())
			})
		})
	})
})
