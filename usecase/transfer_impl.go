package usecase

import (
	"errors"
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
)

type TransferUsecaseImpl struct {
	TransferRepo repository.TransferRepo
	AccRepo      repository.AccountRepo
	UserRepo     repository.UserRepo
	HisRepo      repository.HistoryRepo
}

// Delete implements TransferUsecase
func (u *TransferUsecaseImpl) Delete(id int64) (model.Transfer, error) {
	return u.TransferRepo.Delete(id)
}

// FindAll implements TransferUsecase
func (u *TransferUsecaseImpl) FindAll() ([]model.Transfer, error) {
	return u.TransferRepo.FindAll()
}

// FindById implements TransferUsecase
func (u *TransferUsecaseImpl) FindById(id int64) (model.Transfer, error) {
	return u.TransferRepo.FindById(id)
}

// Save implements TransferUsecase
func (u *TransferUsecaseImpl) Save(newTransfer model.Transfer) (model.Transfer, error) {

	fromacc, err := u.AccRepo.FindById(newTransfer.FromAccountID)
	if err != nil {
		return model.Transfer{}, err
	}

	newTransfer.FromAccount = fromacc

	toacc, err := u.AccRepo.FindById(newTransfer.ToAccountID)
	if err != nil {
		return model.Transfer{}, err
	}

	fromuser, err := u.UserRepo.FindById(fromacc.UserID)
	if err != nil {
		return model.Transfer{}, err
	}

	touser, err := u.UserRepo.FindById(toacc.UserID)
	if err != nil {
		return model.Transfer{}, err
	}

	newTransfer.ToAccount = toacc
	newTransfer.FromAccount.User = fromuser
	newTransfer.ToAccount.User = touser

	if fromacc.Balance < newTransfer.Amount {
		return model.Transfer{}, errors.New(err.Error())
	}

	fromacc.Balance -= newTransfer.Amount
	toacc.Balance += newTransfer.Amount

	_, err = u.AccRepo.Update(fromacc)
	if err != nil {
		return model.Transfer{}, errors.New(err.Error())
	}

	_, err = u.AccRepo.Update(toacc)
	if err != nil {
		return model.Transfer{}, errors.New(err.Error())
	}

	fromaccHis := model.History{
		AccountID: fromacc.ID,
		Account:   fromacc,
		Amount:    -newTransfer.Amount,
	}

	toaccHis := model.History{
		AccountID: toacc.ID,
		Account:   toacc,
		Amount:    +newTransfer.Amount,
	}

	_, err = u.HisRepo.Save(fromaccHis)
	if err != nil {
		return model.Transfer{}, errors.New(err.Error())
	}

	_, err = u.HisRepo.Save(toaccHis)
	if err != nil {
		return model.Transfer{}, errors.New(err.Error())
	}

	return u.TransferRepo.Save(newTransfer)
}

// Update implements TransferUsecase
func (u *TransferUsecaseImpl) Update(updatedTransfer model.Transfer) (model.Transfer, error) {

	// Mendapatkan entitas Transfer sebelumnya dari TransferRepo berdasarkan ID
	previousTransfer, err := u.TransferRepo.FindById(updatedTransfer.ID)
	if err != nil {
		return model.Transfer{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousFromAccID := previousTransfer.FromAccountID
	previousToAccID := previousTransfer.ToAccountID
	previousAmount := previousTransfer.Amount
	previousCreatedAt := previousTransfer.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedTransfer.FromAccountID == 0 {
		updatedTransfer.FromAccountID = previousFromAccID
	}

	if updatedTransfer.ToAccountID == 0 {
		updatedTransfer.ToAccountID = previousToAccID
	}

	if updatedTransfer.Amount == 0 {
		updatedTransfer.Amount = previousAmount
	}

	if updatedTransfer.CreatedAt == (time.Time{}) {
		updatedTransfer.CreatedAt = previousCreatedAt
	}

	fromacc, err := u.AccRepo.FindById(updatedTransfer.FromAccountID)
	if err != nil {
		return model.Transfer{}, err
	}

	updatedTransfer.FromAccount = fromacc

	fromuser, err := u.UserRepo.FindById(fromacc.UserID)
	if err != nil {
		return model.Transfer{}, err
	}

	updatedTransfer.FromAccount.User = fromuser

	toacc, err := u.AccRepo.FindById(updatedTransfer.ToAccountID)
	if err != nil {
		return model.Transfer{}, err
	}

	updatedTransfer.ToAccount = toacc

	touser, err := u.UserRepo.FindById(toacc.UserID)
	if err != nil {
		return model.Transfer{}, err
	}

	updatedTransfer.ToAccount.User = touser

	return u.TransferRepo.Update(updatedTransfer)
}

func NewTransferUsecaseImpl(TransferRepo repository.TransferRepo, UserRepo repository.UserRepo, AccountRepo repository.AccountRepo, HisRepo repository.HistoryRepo) TransferUsecase {
	return &TransferUsecaseImpl{
		TransferRepo: TransferRepo,
		UserRepo:     UserRepo,
		AccRepo:      AccountRepo,
		HisRepo:      HisRepo,
	}
}
