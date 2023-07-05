package e

import (
	"fmt"
	"golang-test/common"
	"net/http"
)

func (s *errorHandlerImpl) PanicHandler(w http.ResponseWriter, r *http.Request) {
	err := recover()
	if err != nil {
		var code int
		switch e := err.(type) {
		case common.Http401:
			code = http.StatusUnauthorized
			s.res.ToJson(w, r, code, s.CreateResponseError(code, "Invalid Authentication"))
			return
		case common.Http400:
			code = http.StatusBadRequest
			s.res.ToJson(w, r, code, s.CreateResponseError(code, "Invalid Format request"))
			return
		case common.Http404:
			code = http.StatusNotFound
			s.res.ToJson(w, r, code, s.CreateResponseError(code, "Data Not Found"))
			return
		case common.Http500:
			code = http.StatusInternalServerError
			s.res.ToJson(w, r, code, s.CreateResponseError(code, "Sorry we are under maintenance"))
			return
		default:
			// if error not found, like panic from memory address or nil pointer
			s.logger.Error(fmt.Sprint(e))
			code = http.StatusInternalServerError
			s.res.ToJson(w, r, code, s.CreateResponseError(code, "Sorry we are under maintenance"))
			return
		}
	}
}

func (s *errorHandlerImpl) CreateResponseError(code int, err string) *common.WebResponseError {
	return &common.WebResponseError{
		Status: &common.WebBaseResponse{
			Msg:  "NOT OK",
			Code: code,
			Err:  err,
		},
		Data: nil,
	}
}
