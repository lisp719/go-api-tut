package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-tut/core"
	"go-api-tut/model"
	"go-api-tut/router"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test/Requests/Users", func() {
	BeforeEach(func() {
		os.Setenv("GORM_DSN", "root:@tcp(mysql:3306)/go_api_tut_test?charset=utf8&parseTime=True&loc=Local")
		core.SetupDb()
		core.Db = core.Db.Begin()
	})

	AfterEach(func() {
		core.Db.Rollback()
	})

	Describe("GET /users", func() {
		Context("with query params", func() {
			It("returns ok response", func() {
				for i := 0; i < 10; i++ {
					user := model.User{Name: "foo"}
					core.Db.Create(&user)
				}
				user := model.User{Name: "foo"}
				another := model.User{Name: "bar"}
				core.Db.Create(&user)
				core.Db.Create(&another)

				r := router.SetupRouter()

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/users?q=fo&page=2", nil)
				r.ServeHTTP(w, req)

				users := []model.User{user}
				expected, _ := json.Marshal(users)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body).To(MatchJSON(expected))
			})
		})

		Context("without query params", func() {
			It("returns ok response", func() {
				user := model.User{Name: "foo"}
				core.Db.Create(&user)

				r := router.SetupRouter()

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/users", nil)
				r.ServeHTTP(w, req)

				users := []model.User{user}
				expected, _ := json.Marshal(users)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body).To(MatchJSON(expected))
			})
		})
	})

	Describe("POST /users", func() {
		It("returns created response", func() {
			r := router.SetupRouter()

			params, _ := json.Marshal(map[string]interface{}{"name": "foo"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(params))
			r.ServeHTTP(w, req)

			var user model.User
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

			var users []model.User
			result := core.Db.Find(&users)

			Expect(result.RowsAffected).To(Equal(int64(1)))
		})
	})

	Describe("GET /users/:id", func() {
		It("returns ok response", func() {
			user := model.User{Name: "foo"}
			core.Db.Create(&user)

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
		It("returns ok response", func() {
			user := model.User{Name: "foo"}
			core.Db.Create(&user)
			r := router.SetupRouter()

			params, _ := json.Marshal(map[string]interface{}{"name": "updated"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/%d", user.ID), bytes.NewBuffer(params))
			r.ServeHTTP(w, req)

			var reload model.User
			core.Db.Find(&reload, user.ID)
			expected, _ := json.Marshal(reload)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body).To(MatchJSON(expected))
		})
	})

	Describe("DELETE /users/:id", func() {
		user := model.User{Name: "foo"}

		BeforeEach(func() {
			core.Db.Create(&user)
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

			var users []model.User
			result := core.Db.Find(&users)
			Expect(result.RowsAffected).To(Equal(int64(0)))
		})
	})
})
