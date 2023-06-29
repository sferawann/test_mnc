package usecase

import (
	"github.com/sferawann/test_mnc/config"
	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
	"github.com/sferawann/test_mnc/token"
	"github.com/sferawann/test_mnc/utils"
)

type AuthUsecaseImpl struct {
	userRepo repository.UserRepo
	sesRepo  repository.SessionRepo
}

// Login implements AuthUsecase
func (u *AuthUsecaseImpl) Login(username, password string) (string, error) {
	// Cari user berdasarkan username
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	verify_error := utils.VerifyPassword(user.Password, password)
	if verify_error != nil {
		return "", err
	}

	config, _ := config.LoadConfig(".")
	// Buat token JWT
	payload := user.ID
	tokenStr, err := token.GenerateToken(config.TokenExpiresIn, payload, config.TokenSecret)
	if err != nil {
		return "", err
	}

	//create session
	userid, err := u.userRepo.FindById(payload)
	if err != nil {
		return "", err
	}
	session := model.Session{
		UserID: user.ID,
		User:   userid,
		Token:  tokenStr,
	}
	_, err = u.sesRepo.Save(session)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Logout implements AuthUsecase
func (u *AuthUsecaseImpl) Logout(token string) (model.Session, error) {
	return u.sesRepo.DeleteByToken(token)
}

func NewAuthUsecaseImpl(userRepo repository.UserRepo, sesRepo repository.SessionRepo) AuthUsecase {
	return &AuthUsecaseImpl{
		userRepo: userRepo,
		sesRepo:  sesRepo,
	}
}
