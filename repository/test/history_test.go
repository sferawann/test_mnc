package repository

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
)

const (
	testFilePathHistory = "history.json" // Path to test data file
)

var (
	testHistory = model.History{
		ID:        1,
		AccountID: 1,
		Amount:    100,
		CreatedAt: time.Now(),
	}
)

func setupHistory(t *testing.T) (repository.HistoryRepo, func()) {
	// Create a temporary test data file
	file, err := os.Create(testFilePathHistory)
	if err != nil {
		t.Fatalf("failed to create test data file: %v", err)
	}

	// Write test history to the file
	histories := []model.History{testHistory}
	err = json.NewEncoder(file).Encode(histories)
	if err != nil {
		t.Fatalf("failed to write test history to file: %v", err)
	}

	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("failed to close test data file: %v", err)
	}

	// Create the repository with the test data file
	repo := repository.NewHistoryRepoImpl(testFilePathHistory)

	// Return the repository and a cleanup function
	return repo, func() {
		// Remove the test data file
		err := os.Remove(testFilePathHistory)
		if err != nil {
			t.Fatalf("failed to remove test data file: %v", err)
		}
	}
}

func TestSaveHistory(t *testing.T) {
	repo, cleanup := setupHistory(t)
	defer cleanup()

	// Save a new history
	newHistory := model.History{
		AccountID: 2,
		Amount:    200,
	}
	savedHistory, err := repo.Save(newHistory)
	if err != nil {
		t.Fatalf("failed to save history: %v", err)
	}

	// Verify the saved history
	if savedHistory.ID == 0 {
		t.Error("saved history ID should not be zero")
	}
	if savedHistory.AccountID != newHistory.AccountID {
		t.Errorf("saved history user ID does not match: got %d, want %d", savedHistory.AccountID, newHistory.AccountID)
	}
	if savedHistory.Amount != newHistory.Amount {
		t.Errorf("saved history amount does not match: got %f, want %f", savedHistory.Amount, newHistory.Amount)
	}
	if savedHistory.CreatedAt.IsZero() {
		t.Error("saved history created at should not be zero")
	}
}

func TestFindAllHistory(t *testing.T) {
	repo, cleanup := setupHistory(t)
	defer cleanup()

	// Retrieve all histories
	histories, err := repo.FindAll()
	if err != nil {
		t.Fatalf("failed to retrieve histories: %v", err)
	}

	// Verify the number of histories
	if len(histories) != 1 {
		t.Errorf("incorrect number of histories: got %d, want %d", len(histories), 1)
	}

	// Verify the retrieved history
	retrievedHistory := histories[0]
	if retrievedHistory.ID != testHistory.ID {
		t.Errorf("retrieved history ID does not match: got %d, want %d", retrievedHistory.ID, testHistory.ID)
	}
	if retrievedHistory.AccountID != testHistory.AccountID {
		t.Errorf("retrieved history user ID does not match: got %d, want %d", retrievedHistory.AccountID, testHistory.AccountID)
	}
	if retrievedHistory.Amount != testHistory.Amount {
		t.Errorf("retrieved history amount does not match: got %f, want %f", retrievedHistory.Amount, testHistory.Amount)
	}
	if !retrievedHistory.CreatedAt.Equal(testHistory.CreatedAt) {
		t.Errorf("retrieved history created at does not match: got %s, want %s", retrievedHistory.CreatedAt, testHistory.CreatedAt)
	}
}

func TestFindByIDHistory(t *testing.T) {
	repo, cleanup := setupHistory(t)
	defer cleanup()

	// Retrieve history by ID
	history, err := repo.FindById(testHistory.ID)
	if err != nil {
		t.Fatalf("failed to retrieve history by ID: %v", err)
	}

	// Verify the retrieved history
	if history.ID != testHistory.ID {
		t.Errorf("retrieved history ID does not match: got %d, want %d", history.ID, testHistory.ID)
	}
	if history.AccountID != testHistory.AccountID {
		t.Errorf("retrieved history user ID does not match: got %d, want %d", history.AccountID, testHistory.AccountID)
	}
	if history.Amount != testHistory.Amount {
		t.Errorf("retrieved history amount does not match: got %f, want %f", history.Amount, testHistory.Amount)
	}
	if !history.CreatedAt.Equal(testHistory.CreatedAt) {
		t.Errorf("retrieved history created at does not match: got %v, want %v", history.CreatedAt, testHistory.CreatedAt)
	}
}

func TestDeleteHistory(t *testing.T) {
	repo, cleanup := setupHistory(t)
	defer cleanup()

	// Delete history by ID
	_, err := repo.Delete(testHistory.ID)
	if err != nil {
		t.Fatalf("failed to delete history: %v", err)
	}

	// Verify that the history is deleted
	_, err = repo.FindById(testHistory.ID)
	if err == nil {
		t.Error("history should be deleted, but it still exists")
	}
}

func TestUpdateHistory(t *testing.T) {
	repo, cleanup := setupHistory(t)
	defer cleanup()

	//update a history
	updatedHistory := model.History{
		ID:        testHistory.ID,
		AccountID: testHistory.AccountID,
		Amount:    testHistory.Amount,
		CreatedAt: testHistory.CreatedAt,
	}

	updatedHistory, err := repo.Update(updatedHistory)
	if err != nil {
		t.Fatalf("failed to update History: %v", err)
	}

	// Retrieve History by ID
	retrievedHistory, err := repo.FindById(testHistory.ID)
	if err != nil {
		t.Fatalf("failed to retrieve History: %v", err)
	}

	//verify the retrieved History
	if retrievedHistory.ID != testHistory.ID {
		t.Errorf("retrieved History ID does not match: got %d, want %d", retrievedHistory.ID, updatedHistory.ID)
	}
	if retrievedHistory.AccountID != testHistory.AccountID {
		t.Errorf("retrieved History Account ID does not match: got %d, want %d", retrievedHistory.AccountID, updatedHistory.AccountID)
	}
	if retrievedHistory.Amount != testHistory.Amount {
		t.Errorf("retrieved History Amount does not match: got %f, want %f", retrievedHistory.Amount, updatedHistory.Amount)
	}
	if !retrievedHistory.CreatedAt.Equal(testHistory.CreatedAt) {
		t.Errorf("retrieved History created at does not match: got %s, want %s", retrievedHistory.CreatedAt, updatedHistory.CreatedAt)
	}
}
