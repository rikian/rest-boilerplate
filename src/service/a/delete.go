package a

import (
	"net/http"

	"golang-test/common"
	"golang-test/src/helper"

	"github.com/julienschmidt/httprouter"
)

func (s *aServiceImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := &common.WebRequestDelete{}
	err := helper.ReadFromRequestBody(r.Body, data)
	if err != nil {
		s.logger.Error(err.Error())
		panic(common.Http400{})
	}
	err = s.validation.ValidationRequestDelete(data)
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
			// if commit succes, we can add caching here. Redis etc...
		}
	}()
	book := &common.Book{}
	s.repo.Delete(r.Context(), tx, book, data.ID)

	//create response
	s.response.ToJson(w, r, http.StatusOK, &common.WebResponseDelete{
		Status: &common.WebBaseResponse{
			Msg:  "OK",
			Code: http.StatusOK,
			Err:  "",
		},
		Data: data.ID,
	})
}
