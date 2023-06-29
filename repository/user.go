package repository

import "github.com/sferawann/test_mnc/model"

type UserRepo interface {
	Save(newUser model.User) (model.User, error)
	Update(updatedUser model.User) (model.User, error)
	Delete(id int64) (model.User, error)
	FindById(id int64) (model.User, error)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
}