package repository_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
)

const (
	testFilePathUser = "user.json" // Path to test data file
)

var (
	testUser = model.User{
		ID:        1,
		Username:  "testuser",
		Password:  "testpassword",
		Email:     "testuser@example.com",
		CreatedAt: time.Now(),
	}
)

func setupUser(t *testing.T) (repository.UserRepo, func()) {
	// Create a temporary test data file
	file, err := os.Create(testFilePathUser)
	if err != nil {
		t.Fatalf("failed to create test data file: %v", err)
	}

	// Write test user to the file
	users := []model.User{testUser}
	err = json.NewEncoder(file).Encode(users)
	if err != nil {
		t.Fatalf("failed to write test user to file: %v", err)
	}

	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("failed to close test data file: %v", err)
	}

	// Create the repository with the test data file
	repo := repository.NewUserRepoImpl(testFilePathUser)

	// Return the repository and a cleanup function
	return repo, func() {
		// Remove the test data file
		err := os.Remove(testFilePathUser)
		if err != nil {
			t.Fatalf("failed to remove test data file: %v", err)
		}
	}
}

func TestSave(t *testing.T) {
	repo, cleanup := setupUser(t)
	defer cleanup()

	// Save a new user
	newUser := model.User{
		Username: "newuser",
		Password: "newpassword",
		Email:    "newuser@example.com",
	}
	savedUser, err := repo.Save(newUser)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	// Verify the saved user
	if savedUser.ID == 0 {
		t.Error("saved user ID should not be zero")
	}
	if savedUser.Username != newUser.Username {
		t.Errorf("saved user username does not match: got %s, want %s", savedUser.Username, newUser.Username)
	}
	if savedUser.Password != newUser.Password {
		t.Errorf("saved user password does not match: got %s, want %s", savedUser.Password, newUser.Password)
	}
	if savedUser.Email != newUser.Email {
		t.Errorf("saved user email does not match: got %s, want %s", savedUser.Email, newUser.Email)
	}
	if savedUser.CreatedAt.IsZero() {
		t.Error("saved user created at should not be zero")
	}
}

func TestFindAll(t *testing.T) {
	repo, cleanup := setupUser(t)
	defer cleanup()

	// Retrieve all users
	users, err := repo.FindAll()
	if err != nil {
		t.Fatalf("failed to retrieve users: %v", err)
	}

	// Verify the number of users
	if len(users) != 1 {
		t.Errorf("incorrect number of users: got %d, want %d", len(users), 1)
	}

	// Verify the retrieved user
	retrievedUser := users[0]
	if retrievedUser.ID != testUser.ID {
		t.Errorf("retrieved user ID does not match: got %d, want %d", retrievedUser.ID, testUser.ID)
	}
	if retrievedUser.Username != testUser.Username {
		t.Errorf("retrieved user username does not match: got %s, want %s", retrievedUser.Username, testUser.Username)
	}
	if retrievedUser.Password != testUser.Password {
		t.Errorf("retrieved user password does not match: got %s, want %s", retrievedUser.Password, testUser.Password)
	}
	if retrievedUser.Email != testUser.Email {
		t.Errorf("retrieved user email does not match: got %s, want %s", retrievedUser.Email, testUser.Email)
	}
	if !retrievedUser.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved user created at does not match: got %s, want %s", retrievedUser.CreatedAt, testUser.CreatedAt)
	}
}

func TestFindById(t *testing.T) {
	repo, cleanup := setupUser(t)
	defer cleanup()

	// Retrieve a user by ID
	retrievedUser, err := repo.FindById(testUser.ID)
	if err != nil {
		t.Fatalf("failed to retrieve user by ID: %v", err)
	}

	// Verify the retrieved user
	if retrievedUser.ID != testUser.ID {
		t.Errorf("retrieved user ID does not match: got %d, want %d", retrievedUser.ID, testUser.ID)
	}
	if retrievedUser.Username != testUser.Username {
		t.Errorf("retrieved user username does not match: got %s, want %s", retrievedUser.Username, testUser.Username)
	}
	if retrievedUser.Password != testUser.Password {
		t.Errorf("retrieved user password does not match: got %s, want %s", retrievedUser.Password, testUser.Password)
	}
	if retrievedUser.Email != testUser.Email {
		t.Errorf("retrieved user email does not match: got %s, want %s", retrievedUser.Email, testUser.Email)
	}
	if !retrievedUser.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved user created at does not match: got %s, want %s", retrievedUser.CreatedAt, testUser.CreatedAt)
	}
}

func TestFindByUsername(t *testing.T) {
	repo, cleanup := setupUser(t)
	defer cleanup()

	// Retrieve a user by username
	retrievedUser, err := repo.FindByUsername(testUser.Username)
	if err != nil {
		t.Fatalf("failed to retrieve user by username: %v", err)
	}

	// Verify the retrieved user
	if retrievedUser.ID != testUser.ID {
		t.Errorf("retrieved user ID does not match: got %d, want %d", retrievedUser.ID, testUser.ID)
	}
	if retrievedUser.Username != testUser.Username {
		t.Errorf("retrieved user username does not match: got %s, want %s", retrievedUser.Username, testUser.Username)
	}
	if retrievedUser.Password != testUser.Password {
		t.Errorf("retrieved user password does not match: got %s, want %s", retrievedUser.Password, testUser.Password)
	}
	if retrievedUser.Email != testUser.Email {
		t.Errorf("retrieved user email does not match: got %s, want %s", retrievedUser.Email, testUser.Email)
	}
	if !retrievedUser.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved user created at does not match: got %s, want %s", retrievedUser.CreatedAt, testUser.CreatedAt)
	}
}

func TestUpdate(t *testing.T) {
	repo, cleanup := setupUser(t)
	defer cleanup()

	// Update a user
	updatedUser := model.User{
		ID:        testUser.ID,
		Username:  "updateduser",
		Password:  "updatedpassword",
		Email:     "updateduser@example.com",
		CreatedAt: testUser.CreatedAt,
	}
	updatedUser, err := repo.Update(updatedUser)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}

	// Retrieve the updated user
	retrievedUser, err := repo.FindById(testUser.ID)
	if err != nil {
		t.Fatalf("failed to retrieve user by ID: %v", err)
	}

	// Verify the updated user
	if retrievedUser.ID != updatedUser.ID {
		t.Errorf("updated user ID does not match: got %d, want %d", retrievedUser.ID, updatedUser.ID)
	}
	if retrievedUser.Username != updatedUser.Username {
		t.Errorf("updated user username does not match: got %s, want %s", retrievedUser.Username, updatedUser.Username)
	}
	if retrievedUser.Password != updatedUser.Password {
		t.Errorf("updated user password does not match: got %s, want %s", retrievedUser.Password, updatedUser.Password)
	}
	if retrievedUser.Email != updatedUser.Email {
		t.Errorf("updated user email does not match: got %s, want %s", retrievedUser.Email, updatedUser.Email)
	}
	if !retrievedUser.CreatedAt.Equal(updatedUser.CreatedAt) {
		t.Errorf("updated user created at does not match: got %s, want %s", retrievedUser.CreatedAt, updatedUser.CreatedAt)
	}
}

func TestDelete(t *testing.T) {
	repo, cleanup := setupUser(t)
	defer cleanup()

	// Delete a user
	deletedUser, err := repo.Delete(testUser.ID)
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	// Verify the deleted user
	if deletedUser.ID != testUser.ID {
		t.Errorf("deleted user ID does not match: got %d, want %d", deletedUser.ID, testUser.ID)
	}
	if deletedUser.Username != testUser.Username {
		t.Errorf("deleted user username does not match: got %s, want %s", deletedUser.Username, testUser.Username)
	}
	if deletedUser.Password != testUser.Password {
		t.Errorf("deleted user password does not match: got %s, want %s", deletedUser.Password, testUser.Password)
	}
	if deletedUser.Email != testUser.Email {
		t.Errorf("deleted user email does not match: got %s, want %s", deletedUser.Email, testUser.Email)
	}
	if !deletedUser.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("deleted user created at does not match: got %s, want %s", deletedUser.CreatedAt, testUser.CreatedAt)
	}

	// Verify the user is deleted
	_, err = repo.FindById(testUser.ID)
	if err == nil {
		t.Error("deleted user still exists")
	}
}
