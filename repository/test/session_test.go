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
	testFilePathSession = "session.json"
)

var (
	testSession = model.Session{
		ID:        1,
		UserID:    1,
		User:      testUser,
		Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgwNDE2MjksImlhdCI6MTY4ODAzOTgyOSwibmJmIjoxNjg4MDM5ODI5LCJzdWIiOjF9.qhVyu3oIv407m6KgqvLQ0aOYbq0zHMVYlQpWgDdCa40",
		CreatedAt: time.Now(),
	}
)

func setupSession(t *testing.T) (repository.SessionRepo, func()) {
	// Create a temporary test data file
	file, err := os.Create(testFilePathSession)
	if err != nil {
		t.Fatalf("failed to create test data file: %v", err)
	}

	// Write test Session to the file
	Sessions := []model.Session{testSession}
	err = json.NewEncoder(file).Encode(Sessions)
	if err != nil {
		t.Fatalf("failed to write test Session to file: %v", err)
	}

	// Close the file
	err = file.Close()
	if err != nil {
		t.Fatalf("failed to close test data file: %v", err)
	}

	// Create the repository with the test data file
	repo := repository.NewSessionRepoImpl(testFilePathSession)

	// Return the repository and a cleanup function
	return repo, func() {
		// Remove the test data file
		err := os.Remove(testFilePathSession)
		if err != nil {
			t.Fatalf("failed to remove test data file: %v", err)
		}
	}
}

func TestSaveSession(t *testing.T) {
	repo, cleanup := setupSession(t)
	defer cleanup()

	// Save a new Session
	newSession := model.Session{
		UserID:    2,
		User:      model.User{ID: 2, Username: "newuser", Password: "newpassword", Email: "newuser@example.com", CreatedAt: time.Now()},
		Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgwNDE2MjksImlhdCI6MTY4ODAzOTgyOSwibmJmIjoxNjg4MDM5ODI5LCJzdWIiOjF9.qhVyu3oIv407m6KgqvLQ0aOYbq0zHMVYlQpWgDdCa40",
		CreatedAt: time.Now(),
	}
	savedSession, err := repo.Save(newSession)
	if err != nil {
		t.Fatalf("failed to save Session: %v", err)
	}

	// Verify the saved Session
	if savedSession.ID == 0 {
		t.Error("saved Session ID should not be zero")
	}
	if savedSession.UserID != newSession.UserID {
		t.Errorf("saved Session user ID does not match: got %d, want %d", savedSession.UserID, newSession.UserID)
	}
	if savedSession.User != newSession.User {
		t.Errorf("saved Session user does not match: got %+v, want %+v", savedSession.User, newSession.User)
	}
	if savedSession.Token != newSession.Token {
		t.Errorf("saved Session Token does not match: got %s, want %s", savedSession.Token, newSession.Token)
	}
	if savedSession.CreatedAt.IsZero() {
		t.Error("saved Session created at should not be zero")
	}
}

func TestFindAllSession(t *testing.T) {
	repo, cleanup := setupSession(t)
	defer cleanup()

	// Retrieve all Session
	Session, err := repo.FindAll()
	if err != nil {
		t.Fatalf("failed to retrieve Session: %v", err)
	}

	// Verify the number of Session
	if len(Session) != 1 {
		t.Errorf("incorrect number of Session: got %d, want %d", len(Session), 1)
	}

	// Verify the retrieved Session
	retrievedSession := Session[0]
	if retrievedSession.ID != testSession.ID {
		t.Errorf("retrieved Session ID does not match: got %d, want %d", retrievedSession.ID, testUser.ID)
	}
	if retrievedSession.UserID != testSession.UserID {
		t.Errorf("retrieved Session user ID does not match: got %d, want %d", retrievedSession.UserID, testSession.UserID)
	}
	if retrievedSession.Token != testSession.Token {
		t.Errorf("retrieved Session Token does not match: got %s, want %s", retrievedSession.Token, testSession.Token)
	}
	if !retrievedSession.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved Session created at does not match: got %s, want %s", retrievedSession.CreatedAt, testUser.CreatedAt)
	}
}

func TestFindByIdSession(t *testing.T) {
	repo, cleanup := setupSession(t)
	defer cleanup()

	// Retrieve Session by ID
	retrievedSession, err := repo.FindById(testSession.ID)
	if err != nil {
		t.Fatalf("failed to retrieve Session: %v", err)
	}

	//verify the retrieved Session
	if retrievedSession.ID != testSession.ID {
		t.Errorf("retrieved Session ID does not match: got %d, want %d", retrievedSession.ID, testUser.ID)
	}
	if retrievedSession.UserID != testSession.UserID {
		t.Errorf("retrieved Session user ID does not match: got %d, want %d", retrievedSession.UserID, testSession.UserID)
	}
	if retrievedSession.Token != testSession.Token {
		t.Errorf("retrieved Session Token does not match: got %s, want %s", retrievedSession.Token, testSession.Token)
	}
	if !retrievedSession.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved Session created at does not match: got %s, want %s", retrievedSession.CreatedAt, testUser.CreatedAt)
	}
}

func TestUpdateSession(t *testing.T) {
	repo, cleanup := setupSession(t)
	defer cleanup()

	//update a Session
	updatedSession := model.Session{
		ID:        testSession.ID,
		UserID:    testSession.UserID,
		User:      testUser,
		Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgwNDE2MjksImlhdCI6MTY4ODAzOTgyOSwibmJmIjoxNjg4MDM5ODI5LCJzdWIiOjF9.qhVyu3oIv407m6KgqvLQ0aOYbq0zHMVYlQpWgDdCa40",
		CreatedAt: testSession.CreatedAt,
	}

	updatedSession, err := repo.Update(updatedSession)
	if err != nil {
		t.Fatalf("failed to update Session: %v", err)
	}

	// Retrieve Session by ID
	retrievedSession, err := repo.FindById(testSession.ID)
	if err != nil {
		t.Fatalf("failed to retrieve Session: %v", err)
	}

	//verify the retrieved Session
	if retrievedSession.ID != testSession.ID {
		t.Errorf("retrieved Session ID does not match: got %d, want %d", retrievedSession.ID, updatedSession.ID)
	}
	if retrievedSession.UserID != testSession.UserID {
		t.Errorf("retrieved Session user ID does not match: got %d, want %d", retrievedSession.UserID, updatedSession.UserID)
	}
	if retrievedSession.Token != testSession.Token {
		t.Errorf("retrieved Session Token does not match: got %s, want %s", retrievedSession.Token, updatedSession.Token)
	}
	if !retrievedSession.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved Session created at does not match: got %s, want %s", retrievedSession.CreatedAt, updatedSession.CreatedAt)
	}
}

func TestDeleteSession(t *testing.T) {
	repo, cleanup := setupSession(t)
	defer cleanup()

	//delete a Session
	deletedSession, err := repo.Delete(testSession.ID)
	if err != nil {
		t.Fatalf("failed to delete Session: %v", err)
	}

	//verify the deleted Session
	if deletedSession.ID != testSession.ID {
		t.Errorf("retrieved Session ID does not match: got %d, want %d", deletedSession.ID, testSession.ID)
	}
	if deletedSession.UserID != testSession.UserID {
		t.Errorf("retrieved Session user ID does not match: got %d, want %d", deletedSession.UserID, testSession.UserID)
	}
	if deletedSession.Token != testSession.Token {
		t.Errorf("retrieved Session Token does not match: got %s, want %s", deletedSession.Token, testSession.Token)
	}
	if !deletedSession.CreatedAt.Equal(testUser.CreatedAt) {
		t.Errorf("retrieved Session created at does not match: got %s, want %s", deletedSession.CreatedAt, testSession.CreatedAt)
	}

	// Verify the user is deleted
	_, err = repo.FindById(testSession.ID)
	if err == nil {
		t.Error("deleted user still exists")
	}
}
