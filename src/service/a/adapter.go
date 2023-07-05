package a

import (
	"golang-test/src/helper"
	"golang-test/src/repositories"
	"golang-test/src/service/r"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewAService(logger *zap.Logger, repo repositories.Repo, db *gorm.DB, validation helper.Validation, serviceR r.ResponseHandler) AService {
	return &aServiceImpl{
		logger:     logger,
		repo:       repo,
		db:         db,
		validation: validation,
		response:   serviceR,
	}
}

type AService interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Read(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type aServiceImpl struct {
	logger     *zap.Logger
	repo       repositories.Repo
	db         *gorm.DB
	validation helper.Validation
	response   r.ResponseHandler
}
