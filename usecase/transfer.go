package usecase

import "github.com/sferawann/test_mnc/model"

type TransferUsecase interface {
	Save(newTransfer model.Transfer) (model.Transfer, error)
	Update(updatedTransfer model.Transfer) (model.Transfer, error)
	Delete(id int64) (model.Transfer, error)
	FindById(id int64) (model.Transfer, error)
	FindAll() ([]model.Transfer, error)
}
