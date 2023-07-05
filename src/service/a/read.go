package a

import (
	"net/http"

	"golang-test/common"
	"golang-test/src/helper"

	"github.com/julienschmidt/httprouter"
)

func (s *aServiceImpl) Read(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := &common.WebRequestRead{}
	err := helper.ReadFromRequestBody(r.Body, data)
	if err != nil {
		s.logger.Error(err.Error())
		panic(common.Http400{})
	}
	err = s.validation.ValidationRequestRead(data)
	if err != nil {
		s.logger.Error(err.Error())
		panic(common.Http400{})
	}
	// begin transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		s.logger.Error(tx.Error.Error())
		panic(common.Http500{})
	}
	defer func() {
		err := recover()
		if err != nil {
			rbErr := s.repo.Rollback(tx)
			if rbErr != nil {
				s.logger.Error(rbErr.Error())
				panic(common.Http500{})
			}
		} else {
			cmErr := s.repo.Commit(tx)
			if cmErr != nil {
				s.logger.Error(cmErr.Error())
				panic(common.Http500{})
			}
			// we can add caching here. Redis etc...
		}
	}()
	books := s.repo.Read(r.Context(), tx, data.Offset)
	//create response
	s.response.ToJson(w, r, http.StatusOK, &common.WebResponseRead{
		Status: &common.WebBaseResponse{
			Msg:  "OK",
			Code: http.StatusOK,
			Err:  "",
		},
		Data: books,
	})
}
