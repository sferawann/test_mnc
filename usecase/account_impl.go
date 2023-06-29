package usecase

import (
	"errors"
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
)

type AccountUsecaseImpl struct {
	AccountRepo repository.AccountRepo
	UserRepo    repository.UserRepo
}

// Delete implements AccountUsecase
func (u *AccountUsecaseImpl) Delete(id int64) (model.Account, error) {
	return u.AccountRepo.Delete(id)
}

// FindAll implements AccountUsecase
func (u *AccountUsecaseImpl) FindAll() ([]model.Account, error) {
	return u.AccountRepo.FindAll()
}

// FindById implements AccountUsecase
func (u *AccountUsecaseImpl) FindById(id int64) (model.Account, error) {
	return u.AccountRepo.FindById(id)
}

// FindByUserId implements AccountUsecase
func (u *AccountUsecaseImpl) FindByUserId(userID int64) ([]model.Account, error) {
	return u.AccountRepo.FindByUserId(userID)
}

// Save implements AccountUsecase
func (u *AccountUsecaseImpl) Save(newAccount model.Account) (model.Account, error) {
	if newAccount.Balance <= 0 {
		return model.Account{}, errors.New("balance must be greater than 0")
	}

	user, err := u.UserRepo.FindById(newAccount.UserID)
	if err != nil {
		return model.Account{}, err
	}

	newAccount.User = user
	return u.AccountRepo.Save(newAccount)
}

// Update implements AccountUsecase
func (u *AccountUsecaseImpl) Update(updatedAccount model.Account) (model.Account, error) {

	// Mendapatkan entitas Account sebelumnya dari AccountRepo berdasarkan ID
	previousAccount, err := u.AccountRepo.FindById(updatedAccount.ID)
	if err != nil {
		return model.Account{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousUserID := previousAccount.UserID
	previousBalance := previousAccount.Balance
	previousCreatedAt := previousAccount.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedAccount.UserID == 0 {
		updatedAccount.UserID = previousUserID
	}
	if updatedAccount.Balance == 0 {
		updatedAccount.Balance = previousBalance
	}
	if updatedAccount.CreatedAt == (time.Time{}) {
		updatedAccount.CreatedAt = previousCreatedAt
	}

	return u.AccountRepo.Update(updatedAccount)
}

func NewAccountUsecaseImpl(AccountRepo repository.AccountRepo, UserRepo repository.UserRepo) AccountUsecase {
	return &AccountUsecaseImpl{
		AccountRepo: AccountRepo,
		UserRepo:    UserRepo,
	}
}
