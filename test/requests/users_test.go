package requests_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-tut/core"
	"go-api-tut/models"
	"go-api-tut/router"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test/Requests/Users", func() {
	BeforeEach(func() {
		core.SetupDb()
	})

	Describe("GET /users", func() {
		user := models.User{Name: "foo"}

		BeforeEach(func() {
			core.Db.Create(&user)
		})

		AfterEach(func() {
			core.Db.Delete(&user)
		})

		It("returns ok response", func() {
			r := router.SetupRouter()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users", nil)
			r.ServeHTTP(w, req)

			users := []models.User{user}
			expected, _ := json.Marshal(users)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body).To(MatchJSON(expected))
		})
	})

	Describe("POST /users", func() {
		AfterEach(func() {
			var user models.User
			core.Db.Last(&user)
			core.Db.Delete(&user)
		})

		It("returns created response", func() {
			r := router.SetupRouter()

			params, _ := json.Marshal(map[string]interface{}{"name": "foo"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(params))
			r.ServeHTTP(w, req)

			var user models.User
			core.Db.Last(&user)
			expected, _ := json.Marshal(user)

			Expect(w.Code).To(Equal(http.StatusCreated))
			Expect(w.Body).To(MatchJSON(expected))
		})

		It("changes users count", func() {
			r := router.SetupRouter()

			params, _ := json.Marshal(map[string]interface{}{"name": "foo"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(params))
			r.ServeHTTP(w, req)

			var users []models.User
			result := core.Db.Find(&users)

			Expect(result.RowsAffected).To(Equal(int64(1)))
		})
	})

	Describe("GET /users/:id", func() {
		user := models.User{Name: "foo"}

		BeforeEach(func() {
			core.Db.Create(&user)
		})

		AfterEach(func() {
			core.Db.Delete(&user)
		})

		It("returns ok response", func() {
			r := router.SetupRouter()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
			r.ServeHTTP(w, req)

			expected, _ := json.Marshal(user)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body).To(MatchJSON(expected))
		})
	})

	Describe("PUT /users/:id", func() {
		user := models.User{Name: "foo"}

		BeforeEach(func() {
			core.Db.Create(&user)
		})

		AfterEach(func() {
			core.Db.Delete(&user)
		})

		It("returns ok response", func() {
			r := router.SetupRouter()

			params, _ := json.Marshal(map[string]interface{}{"name": "updated"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", user.ID), bytes.NewBuffer(params))
			r.ServeHTTP(w, req)

			var reload models.User
			core.Db.Find(&reload, user.ID)
			expected, _ := json.Marshal(reload)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body).To(MatchJSON(expected))
		})
	})

	Describe("DELETE /users/:id", func() {
		user := models.User{Name: "foo"}

		BeforeEach(func() {
			core.Db.Create(&user)
		})

		AfterEach(func() {
			core.Db.Delete(&user)
		})

		It("returns ok response", func() {
			r := router.SetupRouter()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusNoContent))
			Expect(w.Body).To(Equal(bytes.NewBuffer(nil)))
		})

		It("deletes the user", func() {
			r := router.SetupRouter()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
			r.ServeHTTP(w, req)

			var users []models.User
			result := core.Db.Find(&users)
			Expect(result.RowsAffected).To(Equal(int64(0)))
		})
	})
})