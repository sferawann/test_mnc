package usecase

import "github.com/sferawann/test_mnc/model"

type HistoryUsecase interface {
	Save(newHistory model.History) (model.History, error)
	Update(updatedHistory model.History) (model.History, error)
	Delete(id int64) (model.History, error)
	FindById(id int64) (model.History, error)
	FindAll() ([]model.History, error)
}
