package r

import (
	"net/http"

	"go.uber.org/zap"
)

func NewResponseHandler(logger *zap.Logger) ResponseHandler {
	return &responseHandlerImpl{
		logger: logger,
	}
}

type ResponseHandler interface {
	ToJson(w http.ResponseWriter, r *http.Request, code int, data interface{})
}

type responseHandlerImpl struct {
	logger *zap.Logger
}
