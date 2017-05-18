package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Actions", func() {
	// s := dbSession.Copy()
	// defer s.Close()
	Describe("Test get city weather function", func() {
		Context("When the DB is empty", func() {
			It("Should call the API", func() {
				w, err := getCityWeather("london", dailyWeatherC, conf)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).NotTo(BeNil())
				Expect(w.Name).To(Equal("London"))
			})
		})
	})
})
