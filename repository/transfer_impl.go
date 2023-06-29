package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sferawann/test_mnc/model"
)

type TransferRepoImpl struct {
	filePath string
}

// Delete implements TransferRepo
func (r *TransferRepoImpl) Delete(id int64) (model.Transfer, error) {
	Transfers, err := r.FindAll()
	if err != nil {
		return model.Transfer{}, err
	}

	var deletedTransfer model.Transfer
	for i, Transfer := range Transfers {
		if Transfer.ID == id {
			deletedTransfer = Transfer
			Transfers = append(Transfers[:i], Transfers[i+1:]...)
			break
		}
	}

	err = r.writeTransfersToFile(Transfers)
	if err != nil {
		return model.Transfer{}, err
	}

	return deletedTransfer, nil
}

// FindAll implements TransferRepo
func (r *TransferRepoImpl) FindAll() ([]model.Transfer, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Transfer{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var Transfers []model.Transfer
	err = json.NewDecoder(file).Decode(&Transfers)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return Transfers, nil
}

// FindById implements TransferRepo
func (r *TransferRepoImpl) FindById(id int64) (model.Transfer, error) {
	Transfers, err := r.FindAll()
	if err != nil {
		return model.Transfer{}, err
	}

	for _, Transfer := range Transfers {
		if Transfer.ID == id {
			return Transfer, nil
		}
	}

	return model.Transfer{}, fmt.Errorf("transfer by id: %d not found", id)
}

// Save implements TransferRepo
func (r *TransferRepoImpl) Save(newTransfer model.Transfer) (model.Transfer, error) {
	Transfers, err := r.FindAll()
	if err != nil {
		return model.Transfer{}, err
	}

	newTransfer.ID = generateUniqueIDTransfer(Transfers)
	newTransfer.CreatedAt = time.Now()

	Transfers = append(Transfers, newTransfer)

	err = r.writeTransfersToFile(Transfers)
	if err != nil {
		return model.Transfer{}, err
	}

	return newTransfer, nil
}

// Update implements TransferRepo
func (r *TransferRepoImpl) Update(updatedTransfer model.Transfer) (model.Transfer, error) {
	Transfers, err := r.FindAll()
	if err != nil {
		return model.Transfer{}, err
	}

	var found bool
	for i, Transfer := range Transfers {
		if Transfer.ID == updatedTransfer.ID {
			Transfers[i] = updatedTransfer
			found = true
			break
		}
	}

	if !found {
		return model.Transfer{}, fmt.Errorf("transfer by id: %d not found", updatedTransfer.ID)
	}

	err = r.writeTransfersToFile(Transfers)
	if err != nil {
		return model.Transfer{}, err
	}

	return updatedTransfer, nil
}

func (r *TransferRepoImpl) writeTransfersToFile(Transfers []model.Transfer) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(Transfers)
	if err != nil {
		return err
	}

	return nil
}

func generateUniqueIDTransfer(Transfers []model.Transfer) int64 {
	var maxID int64
	for _, Transfer := range Transfers {
		if Transfer.ID > maxID {
			maxID = Transfer.ID
		}
	}
	return maxID + 1
}

func NewTransferRepoImpl(filePath string) TransferRepo {
	return &TransferRepoImpl{
		filePath: filePath,
	}
}
