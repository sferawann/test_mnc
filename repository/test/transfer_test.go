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
	testFilePathTransfer = "transfers.json" // Path to test data file
)

var (
	testTransfer = model.Transfer{
		ID:            1,
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        1000.0,
		CreatedAt:     time.Now(),
	}
)

func setupTransfer(t *testing.T) (repository.TransferRepo, func()) {
	// Create a temporary test data file
	file, err := os.Create(testFilePathTransfer)
	if err != nil {
		t.Fatalf("failed to create test data file: %v", err)
	}

	// Write test transfer to the file
	transfers := []model.Transfer{testTransfer}
	err = json.NewEncoder(file).Encode(transfers)
	if err != nil {
		t.Fatalf("failed to write test transfer to file: %v", err)
	}

	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("failed to close test data file: %v", err)
	}

	// Create the repository with the test data file
	repo := repository.NewTransferRepoImpl(testFilePathTransfer)

	// Return the repository and a cleanup function
	return repo, func() {
		// Remove the test data file
		err := os.Remove(testFilePathTransfer)
		if err != nil {
			t.Fatalf("failed to remove test data file: %v", err)
		}
	}
}

func TestSaveTransfer(t *testing.T) {
	repo, cleanup := setupTransfer(t)
	defer cleanup()

	// Save a new transfer
	newTransfer := model.Transfer{
		FromAccountID: 2,
		ToAccountID:   3,
		Amount:        500.0,
	}
	savedTransfer, err := repo.Save(newTransfer)
	if err != nil {
		t.Fatalf("failed to save transfer: %v", err)
	}

	// Verify the saved transfer
	if savedTransfer.ID == 0 {
		t.Error("saved transfer ID should not be zero")
	}
	if savedTransfer.FromAccountID != newTransfer.FromAccountID {
		t.Errorf("saved transfer from account ID does not match: got %d, want %d", savedTransfer.FromAccountID, newTransfer.FromAccountID)
	}
	if savedTransfer.ToAccountID != newTransfer.ToAccountID {
		t.Errorf("saved transfer to account ID does not match: got %d, want %d", savedTransfer.ToAccountID, newTransfer.ToAccountID)
	}
	if savedTransfer.Amount != newTransfer.Amount {
		t.Errorf("saved transfer amount does not match: got %f, want %f", savedTransfer.Amount, newTransfer.Amount)
	}
	if savedTransfer.CreatedAt.IsZero() {
		t.Error("saved transfer created at should not be zero")
	}
}

func TestFindAllTransfer(t *testing.T) {
	repo, cleanup := setupTransfer(t)
	defer cleanup()

	// Retrieve all transfers
	transfers, err := repo.FindAll()
	if err != nil {
		t.Fatalf("failed to retrieve transfers: %v", err)
	}

	// Verify the number of transfers
	if len(transfers) != 1 {
		t.Errorf("incorrect number of transfers: got %d, want %d", len(transfers), 1)
	}

	// Verify the retrieved transfer
	retrievedTransfer := transfers[0]
	if retrievedTransfer.ID != testTransfer.ID {
		t.Errorf("retrieved transfer ID does not match: got %d, want %d", retrievedTransfer.ID, testTransfer.ID)
	}
	if retrievedTransfer.FromAccountID != testTransfer.FromAccountID {
		t.Errorf("retrieved transfer from account ID does not match: got %d, want %d", retrievedTransfer.FromAccountID, testTransfer.FromAccountID)
	}
	if retrievedTransfer.ToAccountID != testTransfer.ToAccountID {
		t.Errorf("retrieved transfer to account ID does not match: got %d, want %d", retrievedTransfer.ToAccountID, testTransfer.ToAccountID)
	}
	if retrievedTransfer.Amount != testTransfer.Amount {
		t.Errorf("retrieved transfer amount does not match: got %f, want %f", retrievedTransfer.Amount, testTransfer.Amount)
	}
	if !retrievedTransfer.CreatedAt.Equal(testTransfer.CreatedAt) {
		t.Errorf("retrieved transfer created at does not match: got %s, want %s", retrievedTransfer.CreatedAt, testTransfer.CreatedAt)
	}
}

func TestFindByIDTransfer(t *testing.T) {
	repo, cleanup := setupTransfer(t)
	defer cleanup()

	// Retrieve transfer by ID
	transfer, err := repo.FindById(testTransfer.ID)
	if err != nil {
		t.Fatalf("failed to retrieve transfer by ID: %v", err)
	}

	// Verify the retrieved transfer
	if transfer.ID != testTransfer.ID {
		t.Errorf("retrieved transfer ID does not match: got %d, want %d", transfer.ID, testTransfer.ID)
	}
	if transfer.FromAccountID != testTransfer.FromAccountID {
		t.Errorf("retrieved transfer from account ID does not match: got %d, want %d", transfer.FromAccountID, testTransfer.FromAccountID)
	}
	if transfer.ToAccountID != testTransfer.ToAccountID {
		t.Errorf("retrieved transfer to account ID does not match: got %d, want %d", transfer.ToAccountID, testTransfer.ToAccountID)
	}
	if transfer.Amount != testTransfer.Amount {
		t.Errorf("retrieved transfer amount does not match: got %f, want %f", transfer.Amount, testTransfer.Amount)
	}
	if !transfer.CreatedAt.Equal(testTransfer.CreatedAt) {
		t.Errorf("retrieved transfer created at does not match: got %v, want %v", transfer.CreatedAt, testTransfer.CreatedAt)
	}
}

func TestDeleteTransfer(t *testing.T) {
	repo, cleanup := setupTransfer(t)
	defer cleanup()

	// Delete transfer by ID
	_, err := repo.Delete(testTransfer.ID)
	if err != nil {
		t.Fatalf("failed to delete transfer: %v", err)
	}

	// Verify that the transfer is deleted
	_, err = repo.FindById(testTransfer.ID)
	if err == nil {
		t.Error("transfer should be deleted, but it still exists")
	}
}

func TestUpdateTransfer(t *testing.T) {
	repo, cleanup := setupTransfer(t)
	defer cleanup()

	// Update a transfer
	updatedTransfer := model.Transfer{
		ID:            testTransfer.ID,
		FromAccountID: testTransfer.FromAccountID,
		ToAccountID:   testTransfer.ToAccountID,
		Amount:        testTransfer.Amount,
		CreatedAt:     testTransfer.CreatedAt,
	}

	updatedTransfer, err := repo.Update(updatedTransfer)
	if err != nil {
		t.Fatalf("failed to update transfer: %v", err)
	}

	// Retrieve transfer by ID
	retrievedTransfer, err := repo.FindById(testTransfer.ID)
	if err != nil {
		t.Fatalf("failed to retrieve transfer: %v", err)
	}

	// Verify the retrieved transfer
	if retrievedTransfer.ID != testTransfer.ID {
		t.Errorf("retrieved transfer ID does not match: got %d, want %d", retrievedTransfer.ID, updatedTransfer.ID)
	}
	if retrievedTransfer.FromAccountID != testTransfer.FromAccountID {
		t.Errorf("retrieved transfer from account ID does not match: got %d, want %d", retrievedTransfer.FromAccountID, updatedTransfer.FromAccountID)
	}
	if retrievedTransfer.ToAccountID != testTransfer.ToAccountID {
		t.Errorf("retrieved transfer to account ID does not match: got %d, want %d", retrievedTransfer.ToAccountID, updatedTransfer.ToAccountID)
	}
	if retrievedTransfer.Amount != testTransfer.Amount {
		t.Errorf("retrieved transfer amount does not match: got %f, want %f", retrievedTransfer.Amount, updatedTransfer.Amount)
	}
	if !retrievedTransfer.CreatedAt.Equal(testTransfer.CreatedAt) {
		t.Errorf("retrieved transfer created at does not match: got %s, want %s", retrievedTransfer.CreatedAt, updatedTransfer.CreatedAt)
	}
}
