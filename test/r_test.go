package test_test

import (
	"encoding/json"
	"golang-test/cmd/migrations"
	"golang-test/common"
	"golang-test/config"
	h "golang-test/src/handler"
	"golang-test/src/helper"
	"golang-test/src/repositories"
	"golang-test/src/router"
	"golang-test/src/service/a"
	"golang-test/src/service/e"
	"golang-test/src/service/r"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestUserService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CRUD Service")
}

var (
	host     string
	handler  http.Handler
	recorder *httptest.ResponseRecorder
	jwt      helper.JWT
	db       *gorm.DB
)

var _ = BeforeEach(func() {
	// initial di
	// load env file
	config.LoadEnvFile()
	// build zap logger
	logger := config.BuildLogger()
	// create databse connection
	db = config.ConnectDB()
	// initial service
	jwt = helper.NewJwtImpl([]byte(os.Getenv("JWT_SECRET")))
	validate := helper.NewValidation(validator.New())
	repo := repositories.NewRepo(logger)
	serviceR := r.NewResponseHandler(logger)
	serviceE := e.NewErrorHandler(serviceR, logger)
	serviceA := a.NewAService(logger, repo, db, validate, serviceR)
	handler = h.NewHandler(logger, serviceE, serviceA, jwt)
	recorder = httptest.NewRecorder()

	// setup host
	host = "http://localhost:" + os.Getenv("PORT")
})

var _ = Describe("Test Handler for CRUD", func() {
	Describe("authentication for every request", func() {
		It("has valid token, will return status code 200", func() {
			validToken, _ := jwt.GenerateJWT("", 3600)
			data := &common.WebRequestRead{
				Offset: 0,
			}
			jb, _ := json.Marshal(data)
			migrations.RunMigration(db)
			request := httptest.NewRequest(http.MethodGet, host+router.Read, strings.NewReader(string(jb)))
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("Authorization", "Bearer "+validToken)
			handler.ServeHTTP(recorder, request)
			response := recorder.Result()
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("has invalid token, will return status code 401 unauthorized", func() {
			invalidToken := "invalid token"
			request := httptest.NewRequest(http.MethodGet, host+router.Read, nil)
			request.Header.Add("Authorization", "Bearer "+invalidToken)
			handler.ServeHTTP(recorder, request)
			response := recorder.Result()
			Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
		})
	})

	Describe("Read Request", func() {
		It("will return data with status code 200", func() {
			validToken, _ := jwt.GenerateJWT("", 3600)
			data := &common.WebRequestRead{
				Offset: 0,
			}
			jb, _ := json.Marshal(data)
			migrations.RunMigration(db)
			request := httptest.NewRequest(http.MethodGet, host+router.Read, strings.NewReader(string(jb)))
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("Authorization", "Bearer "+validToken)

			handler.ServeHTTP(recorder, request)
			response := recorder.Result()

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			readResponse := &common.WebResponseRead{}
			err := json.NewDecoder(response.Body).Decode(readResponse)
			Expect(err).To(BeNil())
			Expect(readResponse.Status.Msg).To(Equal("OK"))
			Expect(readResponse.Status.Code).To(Equal(200))
			Expect(readResponse.Status.Err).Should(BeEmpty())
			Expect(len(readResponse.Data)).To(Equal(3)) // limit 3
			Expect(readResponse.Data[0].ID).ShouldNot(BeEmpty())
		})
	})
	Describe("Create Request", func() {
		It("will return data with status code 200", func() {

		})
	})

	Describe("Update Request", func() {
		It("will return data with status code 200", func() {

		})
	})

	Describe("Delete Request", func() {
		It("will return data with status code 200", func() {

		})
	})
})
