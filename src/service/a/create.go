package a

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"

	"golang-test/common"
	"golang-test/src/helper"
)

func (s *aServiceImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := &common.WebRequestCreate{}
	err := helper.ReadFromRequestBody(r.Body, data)
	if err != nil {
		s.logger.Error(err.Error())
		panic(common.Http400{})
	}
	err = s.validation.ValidationRequestCreate(data)
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
	dataBooks := &common.Book{
		ID:          uuid.New().String(),
		JudulBuku:   data.JudulBuku,
		Description: data.Description,
		KategoriID:  uint(data.KategoriID),
		KeywordID:   uint(data.KeywordID),
		Harga:       fmt.Sprintf("Rp %v ,-", data.Harga),
		Penerbit:    data.Penerbit,
	}
	// inser data book to db
	s.repo.Create(r.Context(), tx, dataBooks)
	//create response
	s.response.ToJson(w, r, http.StatusOK, &common.WebResponseCreate{
		Status: &common.WebBaseResponse{
			Msg:  "OK",
			Code: http.StatusOK,
			Err:  "",
		},
		Data: dataBooks.ID,
	})
}
