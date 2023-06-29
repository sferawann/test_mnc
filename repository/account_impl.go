package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sferawann/test_mnc/model"
)

type AccountRepoImpl struct {
	filePath string
}

// Delete implements AccountRepo
func (r *AccountRepoImpl) Delete(id int64) (model.Account, error) {
	Accounts, err := r.FindAll()
	if err != nil {
		return model.Account{}, err
	}

	var deletedAccount model.Account
	for i, Account := range Accounts {
		if Account.ID == id {
			deletedAccount = Account
			Accounts = append(Accounts[:i], Accounts[i+1:]...)
			break
		}
	}

	err = r.writeAccountsToFile(Accounts)
	if err != nil {
		return model.Account{}, err
	}

	return deletedAccount, nil
}

// FindAll implements AccountRepo
func (r *AccountRepoImpl) FindAll() ([]model.Account, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Account{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var Accounts []model.Account
	err = json.NewDecoder(file).Decode(&Accounts)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return Accounts, nil
}

// FindById implements AccountRepo
func (r *AccountRepoImpl) FindById(id int64) (model.Account, error) {
	Accounts, err := r.FindAll()
	if err != nil {
		return model.Account{}, err
	}

	for _, Account := range Accounts {
		if Account.ID == id {
			return Account, nil
		}
	}

	return model.Account{}, fmt.Errorf("account by id: %d not found", id)
}

// FindByUserID implements AccountRepo
func (r *AccountRepoImpl) FindByUserId(userID int64) ([]model.Account, error) {
	Accounts, err := r.FindAll()
	if err != nil {
		return nil, err
	}

	var AccByUserID []model.Account
	for _, Account := range Accounts {
		if Account.UserID == userID {
			AccByUserID = append(AccByUserID, Account)
		}
	}

	if len(AccByUserID) == 0 {
		return nil, fmt.Errorf("no accounts found for userID: %d", userID)
	}

	return AccByUserID, nil
}

// Save implements AccountRepo
func (r *AccountRepoImpl) Save(newAccount model.Account) (model.Account, error) {
	Accounts, err := r.FindAll()
	if err != nil {
		return model.Account{}, err
	}

	newAccount.ID = generateUniqueIDAccount(Accounts)
	newAccount.CreatedAt = time.Now()

	Accounts = append(Accounts, newAccount)

	err = r.writeAccountsToFile(Accounts)
	if err != nil {
		return model.Account{}, err
	}

	return newAccount, nil
}

// Update implements AccountRepo
func (r *AccountRepoImpl) Update(updatedAccount model.Account) (model.Account, error) {
	Accounts, err := r.FindAll()
	if err != nil {
		return model.Account{}, err
	}

	var found bool
	for i, Account := range Accounts {
		if Account.ID == updatedAccount.ID {
			Accounts[i] = updatedAccount
			found = true
			break
		}
	}

	if !found {
		return model.Account{}, fmt.Errorf("account by id: %d not found", updatedAccount.ID)
	}

	err = r.writeAccountsToFile(Accounts)
	if err != nil {
		return model.Account{}, err
	}

	return updatedAccount, nil
}

func (r *AccountRepoImpl) writeAccountsToFile(Accounts []model.Account) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(Accounts)
	if err != nil {
		return err
	}

	return nil
}

func generateUniqueIDAccount(Accounts []model.Account) int64 {
	var maxID int64
	for _, Account := range Accounts {
		if Account.ID > maxID {
			maxID = Account.ID
		}
	}
	return maxID + 1
}

func NewAccountRepoImpl(filePath string) AccountRepo {
	return &AccountRepoImpl{
		filePath: filePath,
	}
}
