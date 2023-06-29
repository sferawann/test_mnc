package usecase

import "github.com/sferawann/test_mnc/model"

type SessionUsecase interface {
	Save(newSession model.Session) (model.Session, error)
	Update(updatedSession model.Session) (model.Session, error)
	Delete(id int64) (model.Session, error)
	FindById(id int64) (model.Session, error)
	FindAll() ([]model.Session, error)
}
