package middleware

import (
	"context"
	"encoding/json"
	"golang-test/common"
	"golang-test/src/helper"
	"net/http"
	"strings"

	"golang-test/src/service/a"
	"golang-test/src/service/e"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	logger   *zap.Logger
	router   http.Handler
	serviceE e.ErrorHandler
	serviceA a.AService
	jwt      helper.JWT
}

func NewAuthMiddleware(logger *zap.Logger, router http.Handler, serviceE e.ErrorHandler, serviceA a.AService, jwt helper.JWT) *AuthMiddleware {
	return &AuthMiddleware{
		logger:   logger,
		serviceE: serviceE,
		serviceA: serviceA,
		jwt:      jwt,
		router:   router,
	}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/* all request start here*/
	var err error
	// add request info to logger
	logReq := &common.LogRequest{
		ID:     uuid.New().String(),
		Method: r.Method,
		Path:   r.URL.Path,
		Ip:     r.Header.Get("x-real-ip"),
		Ua:     r.Header.Get("user-agent"),
	}
	reqInfo, _ := json.Marshal(logReq)
	m.logger.Info(string(reqInfo))

	// Create a new context with the request ID if need track use activity
	reqVal := &common.RequestType{ID: logReq.ID}
	ctx := context.WithValue(r.Context(), common.RequestKey, reqVal)
	r = r.WithContext(ctx)

	// handle panic request
	defer m.serviceE.PanicHandler(w, r)

	// authentication
	auth := r.Header.Get("authorization")
	bearer := strings.Split(auth, " ")
	if strings.ToLower(bearer[0]) != "bearer" || len(bearer) != 2 {
		m.logger.Error("invalid token bearer")
		panic(common.Http401{})
	}
	_, err = m.jwt.ClaimJWT(bearer[1])
	if err != nil {
		m.logger.Error(err.Error())
		panic(common.Http401{})
	}
	m.router.ServeHTTP(w, r)
}
