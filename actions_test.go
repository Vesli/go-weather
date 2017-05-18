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
				w, err := getWeatherFromDB("roma", dailyWeatherC, conf)
				Expect(err).To(HaveOccurred())
				Expect(w).To(BeNil())

				w, err = getCityWeatherFromAPI("roma", dailyWeatherC, conf)
				Expect(err).NotTo(HaveOccurred())
				Expect(w).NotTo(BeNil())
				Expect(w.Name).To(Equal("Roma"))
			})
		})
	})
})
