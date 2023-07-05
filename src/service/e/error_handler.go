package e

import (
	"golang-test/src/service/r"
	"net/http"

	"go.uber.org/zap"
)

func NewErrorHandler(res r.ResponseHandler, logger *zap.Logger) ErrorHandler {
	return &errorHandlerImpl{
		res:    res,
		logger: logger,
	}
}

type ErrorHandler interface {
	PanicHandler(w http.ResponseWriter, r *http.Request)
}

type errorHandlerImpl struct {
	res    r.ResponseHandler
	logger *zap.Logger
}
