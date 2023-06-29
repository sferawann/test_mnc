package repository

import "github.com/sferawann/test_mnc/model"

type AccountRepo interface {
	Save(newAccount model.Account) (model.Account, error)
	Update(updatedAccount model.Account) (model.Account, error)
	Delete(id int64) (model.Account, error)
	FindById(id int64) (model.Account, error)
	FindByUserId(userID int64) ([]model.Account, error)
	FindAll() ([]model.Account, error)
}
