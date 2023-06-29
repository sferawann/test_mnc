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
	testFilePathAccount = "account.json"
)

var (
	testUser = model.User{
		ID:        1,
		Username:  "testuser",
		Password:  "testpassword",
		Email:     "testuser@example.com",
		CreatedAt: time.Now(),
	}

	testAccount = model.Account{
		ID:        1,
		UserID:    1,
		User:      testUser,
		Balance:   1000.0,
		CreatedAt: time.Now(),
	}
)

func setupAccount(t *testing.T) (repository.AccountRepo, func()) {
	// Create a temporary test data file
	file, err := os.Create(testFilePathAccount)
	if err != nil {
		t.Fatalf("failed to create test data file: %v", err)
	}

	// Write test Account to the file
	Accounts := []model.Account{testAccount}
	err = json.NewEncoder(file).Encode(Accounts)
	if err != nil {
		t.Fatalf("failed to write test Account to file: %v", err)
	}

	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("failed to close test data file: %v", err)
	}

	// Create the repository with the test data file
	repo := repository.NewAccountRepoImpl(testFilePathAccount)

	// Return the repository and a cleanup function
	return repo, func() {
		// Remove the test data file
		err := os.Remove(testFilePathAccount)
		if err != nil {
			t.Fatalf("failed to remove test data file: %v", err)
		}
	}
}

func TestSaveAccount(t *testing.T) {
	repo, cleanup := setupAccount(t)
	defer cleanup()

	// Save a new account
	newAccount := model.Account{
		UserID:    2,
		User:      model.User{ID: 2, Username: "newuser", Password: "newpassword", Email: "newuser@example.com", CreatedAt: time.Now()},
		Balance:   2000.0,
		CreatedAt: time.Now(),
	}
	savedAccount, err := repo.Save(newAccount)
	if err != nil {
		t.Fatalf("failed to save account: %v", err)
	}

	// Verify the saved account
	if savedAccount.ID == 0 {
		t.Error("saved account ID should not be zero")
	}
	if savedAccount.UserID != newAccount.UserID {
		t.Errorf("saved account user ID does not match: got %d, want %d", savedAccount.UserID, newAccount.UserID)
	}
	if savedAccount.User != newAccount.User {
		t.Errorf("saved account user does not match: got %+v, want %+v", savedAccount.User, newAccount.User)
	}
	if savedAccount.Balance != newAccount.Balance {
		t.Errorf("saved account balance does not match: got %f, want %f", savedAccount.Balance, newAccount.Balance)
	}
	if savedAccount.CreatedAt.IsZero() {
		t.Error("saved account created at should not be zero")
	}
}

func TestFindAllAccount(t *testing.T) {
	repo, cleanup := setupAccount(t)
	defer cleanup()

	// Retrieve all Account
	Account, err := repo.FindAll()
	if err != nil {
		t.Fatalf("failed to retrieve Account: %v", err)
	}

	// Verify the number of Account
	if len(Account) != 1 {
		t.Errorf("incorrect number of Account: got %d, want %d", len(Account), 1)
	}

	// Verify the retrieved Account
	retrievedAccount := Account[0]
	if retrievedAccount.ID != testAccount.ID {
		t.Errorf("retrieved account ID does not match: got %d, want %d", retrievedAccount.ID, testUser.ID)
	}
	if retrievedAccount.UserID != testAccount.UserID {
		t.Errorf("retrieved account user ID does not match: got %d, want %d", retrievedAccount.UserID, testAccount.UserID)
	}
	if retrievedAccount.Balance != testAccount.Balance {
		t.Errorf("retrieved account balance does not match: got %f, want %f", retrievedAccount.Balance, testAccount.Balance)
	}
	if !retrievedAccount.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved account created at does not match: got %s, want %s", retrievedAccount.CreatedAt, testUser.CreatedAt)
	}
}

func TestFindByIdAccount(t *testing.T) {
	repo, cleanup := setupAccount(t)
	defer cleanup()

	// Retrieve Account by ID
	retrievedAccount, err := repo.FindById(testAccount.ID)
	if err != nil {
		t.Fatalf("failed to retrieve Account: %v", err)
	}

	//verify the retrieved Account
	if retrievedAccount.ID != testAccount.ID {
		t.Errorf("retrieved account ID does not match: got %d, want %d", retrievedAccount.ID, testUser.ID)
	}
	if retrievedAccount.UserID != testAccount.UserID {
		t.Errorf("retrieved account user ID does not match: got %d, want %d", retrievedAccount.UserID, testAccount.UserID)
	}
	if retrievedAccount.Balance != testAccount.Balance {
		t.Errorf("retrieved account balance does not match: got %f, want %f", retrievedAccount.Balance, testAccount.Balance)
	}
	if !retrievedAccount.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved account created at does not match: got %s, want %s", retrievedAccount.CreatedAt, testUser.CreatedAt)
	}
}

func TestUpdateAccount(t *testing.T) {
	repo, cleanup := setupAccount(t)
	defer cleanup()

	//update a account
	updatedAccount := model.Account{
		ID:        testAccount.ID,
		UserID:    testAccount.UserID,
		User:      testUser,
		Balance:   1000.0,
		CreatedAt: testAccount.CreatedAt,
	}

	updatedAccount, err := repo.Update(updatedAccount)
	if err != nil {
		t.Fatalf("failed to update account: %v", err)
	}

	// Retrieve Account by ID
	retrievedAccount, err := repo.FindById(testAccount.ID)
	if err != nil {
		t.Fatalf("failed to retrieve Account: %v", err)
	}

	//verify the retrieved Account
	if retrievedAccount.ID != testAccount.ID {
		t.Errorf("retrieved account ID does not match: got %d, want %d", retrievedAccount.ID, updatedAccount.ID)
	}
	if retrievedAccount.UserID != testAccount.UserID {
		t.Errorf("retrieved account user ID does not match: got %d, want %d", retrievedAccount.UserID, updatedAccount.UserID)
	}
	if retrievedAccount.Balance != testAccount.Balance {
		t.Errorf("retrieved account balance does not match: got %f, want %f", retrievedAccount.Balance, updatedAccount.Balance)
	}
	if !retrievedAccount.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved account created at does not match: got %s, want %s", retrievedAccount.CreatedAt, updatedAccount.CreatedAt)
	}
}

func TestDeleteAccount(t *testing.T) {
	repo, cleanup := setupAccount(t)
	defer cleanup()

	//delete a account
	deletedAccount, err := repo.Delete(testAccount.ID)
	if err != nil {
		t.Fatalf("failed to delete account: %v", err)
	}

	//verify the deleted Account
	if deletedAccount.ID != testAccount.ID {
		t.Errorf("retrieved account ID does not match: got %d, want %d", deletedAccount.ID, testAccount.ID)
	}
	if deletedAccount.UserID != testAccount.UserID {
		t.Errorf("retrieved account user ID does not match: got %d, want %d", deletedAccount.UserID, testAccount.UserID)
	}
	if deletedAccount.Balance != testAccount.Balance {
		t.Errorf("retrieved account balance does not match: got %f, want %f", deletedAccount.Balance, testAccount.Balance)
	}
	if !deletedAccount.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved account created at does not match: got %s, want %s", deletedAccount.CreatedAt, testAccount.CreatedAt)
	}

	// Verify the user is deleted
	_, err = repo.FindById(testAccount.ID)
	if err == nil {
		t.Error("deleted user still exists")
	}
}
