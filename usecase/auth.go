package usecase

import "github.com/sferawann/test_mnc/model"

type AuthUsecase interface {
	Login(username, password string) (string, error)
	Logout(token string) (model.Session, error)
}
