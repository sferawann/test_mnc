package usecase

import (
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
)

type HistoryUsecaseImpl struct {
	HistoryRepo repository.HistoryRepo
	AccRepo     repository.AccountRepo
	UserRepo    repository.UserRepo
}

// Delete implements HistoryUsecase
func (u *HistoryUsecaseImpl) Delete(id int64) (model.History, error) {
	return u.HistoryRepo.Delete(id)
}

// FindAll implements HistoryUsecase
func (u *HistoryUsecaseImpl) FindAll() ([]model.History, error) {
	return u.HistoryRepo.FindAll()
}

// FindById implements HistoryUsecase
func (u *HistoryUsecaseImpl) FindById(id int64) (model.History, error) {
	return u.HistoryRepo.FindById(id)
}

// Save implements HistoryUsecase
func (u *HistoryUsecaseImpl) Save(newHistory model.History) (model.History, error) {

	acc, err := u.AccRepo.FindById(newHistory.AccountID)
	if err != nil {
		return model.History{}, err
	}

	newHistory.Account = acc

	user, err := u.UserRepo.FindById(acc.UserID)
	if err != nil {
		return model.History{}, err
	}

	newHistory.Account.User = user

	return u.HistoryRepo.Save(newHistory)
}

// Update implements HistoryUsecase
func (u *HistoryUsecaseImpl) Update(updatedHistory model.History) (model.History, error) {

	// Mendapatkan entitas History sebelumnya dari HistoryRepo berdasarkan ID
	previousHistory, err := u.HistoryRepo.FindById(updatedHistory.ID)
	if err != nil {
		return model.History{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousAccountID := previousHistory.AccountID
	previousAmount := previousHistory.Amount
	previousCreatedAt := previousHistory.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedHistory.AccountID == 0 {
		updatedHistory.AccountID = previousAccountID
	}
	if updatedHistory.Amount == 0 {
		updatedHistory.Amount = previousAmount
	}

	if updatedHistory.CreatedAt == (time.Time{}) {
		updatedHistory.CreatedAt = previousCreatedAt
	}

	acc, err := u.AccRepo.FindById(updatedHistory.AccountID)
	if err != nil {
		return model.History{}, err
	}

	updatedHistory.Account = acc

	user, err := u.UserRepo.FindById(acc.UserID)
	if err != nil {
		return model.History{}, err
	}

	updatedHistory.Account.User = user

	return u.HistoryRepo.Update(updatedHistory)
}

func NewHistoryUsecaseImpl(HistoryRepo repository.HistoryRepo, UserRepo repository.UserRepo, AccountRepo repository.AccountRepo) HistoryUsecase {
	return &HistoryUsecaseImpl{
		HistoryRepo: HistoryRepo,
		UserRepo:    UserRepo,
		AccRepo:     AccountRepo,
	}
}
