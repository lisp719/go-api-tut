package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ping", func() {
	Describe("/ping", func() {
		var (
			response *httptest.ResponseRecorder
		)

		BeforeEach(func() {
			router := setupRouter()
			response = httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/ping", nil)
			router.ServeHTTP(response, req)
		})

		It("returns success response", func() {
			Expect(response.Code).To((Equal(200)))
		})

		It("checks response body", func() {
			Expect(response.Body.String()).To((Equal("pong")))
		})
	})
})
