package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoWeather(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWeather Suite")
}
