package handler

import (
	"golang-test/src/helper"
	"golang-test/src/middleware"
	"golang-test/src/router"
	"golang-test/src/service/a"
	"golang-test/src/service/e"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func NewHandler(logger *zap.Logger, serviceE e.ErrorHandler, serviceA a.AService, jwt helper.JWT) http.Handler {
	r := httprouter.New()
	r.GET(router.Read, serviceA.Read)
	r.POST(router.Create, serviceA.Create)
	r.PATCH(router.Update, serviceA.Update)
	r.DELETE(router.Delete, serviceA.Delete)
	return middleware.NewAuthMiddleware(logger, r, serviceE, serviceA, jwt)
}
