package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sferawann/test_mnc/model"
)

type HistoryRepoImpl struct {
	filePath string
}

// Delete implements HistoryRepo
func (r *HistoryRepoImpl) Delete(id int64) (model.History, error) {
	Historys, err := r.FindAll()
	if err != nil {
		return model.History{}, err
	}

	var deletedHistory model.History
	for i, History := range Historys {
		if History.ID == id {
			deletedHistory = History
			Historys = append(Historys[:i], Historys[i+1:]...)
			break
		}
	}

	err = r.writeHistorysToFile(Historys)
	if err != nil {
		return model.History{}, err
	}

	return deletedHistory, nil
}

// FindAll implements HistoryRepo
func (r *HistoryRepoImpl) FindAll() ([]model.History, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.History{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var Historys []model.History
	err = json.NewDecoder(file).Decode(&Historys)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return Historys, nil
}

// FindById implements HistoryRepo
func (r *HistoryRepoImpl) FindById(id int64) (model.History, error) {
	Historys, err := r.FindAll()
	if err != nil {
		return model.History{}, err
	}

	for _, History := range Historys {
		if History.ID == id {
			return History, nil
		}
	}

	return model.History{}, fmt.Errorf("history by id: %d not found", id)
}

// Save implements HistoryRepo
func (r *HistoryRepoImpl) Save(newHistory model.History) (model.History, error) {
	Historys, err := r.FindAll()
	if err != nil {
		return model.History{}, err
	}

	newHistory.ID = generateUniqueIDHistory(Historys)
	newHistory.CreatedAt = time.Now()

	Historys = append(Historys, newHistory)

	err = r.writeHistorysToFile(Historys)
	if err != nil {
		return model.History{}, err
	}

	return newHistory, nil
}

// Update implements HistoryRepo
func (r *HistoryRepoImpl) Update(updatedHistory model.History) (model.History, error) {
	Historys, err := r.FindAll()
	if err != nil {
		return model.History{}, err
	}

	var found bool
	for i, History := range Historys {
		if History.ID == updatedHistory.ID {
			Historys[i] = updatedHistory
			found = true
			break
		}
	}

	if !found {
		return model.History{}, fmt.Errorf("history by id: %d not found", updatedHistory.ID)
	}

	err = r.writeHistorysToFile(Historys)
	if err != nil {
		return model.History{}, err
	}

	return updatedHistory, nil
}

func (r *HistoryRepoImpl) writeHistorysToFile(Historys []model.History) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(Historys)
	if err != nil {
		return err
	}

	return nil
}

func generateUniqueIDHistory(Historys []model.History) int64 {
	var maxID int64
	for _, History := range Historys {
		if History.ID > maxID {
			maxID = History.ID
		}
	}
	return maxID + 1
}

func NewHistoryRepoImpl(filePath string) HistoryRepo {
	return &HistoryRepoImpl{
		filePath: filePath,
	}
}
