package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sferawann/test_mnc/model"
)

type SessionRepoImpl struct {
	filePath string
}

// DeleteByToken implements SessionRepo
func (r *SessionRepoImpl) DeleteByToken(token string) (model.Session, error) {
	Sessions, err := r.FindAll()
	if err != nil {
		return model.Session{}, err
	}

	var deletedSession model.Session
	for i, Session := range Sessions {
		if Session.Token == token {
			deletedSession = Session
			Sessions = append(Sessions[:i], Sessions[i+1:]...)
			break
		}
	}

	err = r.writeSessionsToFile(Sessions)
	if err != nil {
		return model.Session{}, err
	}

	return deletedSession, nil
}

// Delete implements SessionRepo
func (r *SessionRepoImpl) Delete(id int64) (model.Session, error) {
	Sessions, err := r.FindAll()
	if err != nil {
		return model.Session{}, err
	}

	var deletedSession model.Session
	for i, Session := range Sessions {
		if Session.ID == id {
			deletedSession = Session
			Sessions = append(Sessions[:i], Sessions[i+1:]...)
			break
		}
	}

	err = r.writeSessionsToFile(Sessions)
	if err != nil {
		return model.Session{}, err
	}

	return deletedSession, nil
}

// FindAll implements SessionRepo
func (r *SessionRepoImpl) FindAll() ([]model.Session, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Session{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var Sessions []model.Session
	err = json.NewDecoder(file).Decode(&Sessions)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return Sessions, nil
}

// FindById implements SessionRepo
func (r *SessionRepoImpl) FindById(id int64) (model.Session, error) {
	Sessions, err := r.FindAll()
	if err != nil {
		return model.Session{}, err
	}

	for _, Session := range Sessions {
		if Session.ID == id {
			return Session, nil
		}
	}

	return model.Session{}, fmt.Errorf("session by id: %d not found", id)
}

// Save implements SessionRepo
func (r *SessionRepoImpl) Save(newSession model.Session) (model.Session, error) {
	Sessions, err := r.FindAll()
	if err != nil {
		return model.Session{}, err
	}

	newSession.ID = generateUniqueIDSession(Sessions)
	newSession.CreatedAt = time.Now()

	Sessions = append(Sessions, newSession)

	err = r.writeSessionsToFile(Sessions)
	if err != nil {
		return model.Session{}, err
	}

	return newSession, nil
}

// Update implements SessionRepo
func (r *SessionRepoImpl) Update(updatedSession model.Session) (model.Session, error) {
	Sessions, err := r.FindAll()
	if err != nil {
		return model.Session{}, err
	}

	var found bool
	for i, Session := range Sessions {
		if Session.ID == updatedSession.ID {
			Sessions[i] = updatedSession
			found = true
			break
		}
	}

	if !found {
		return model.Session{}, fmt.Errorf("session by id: %d not found", updatedSession.ID)
	}

	err = r.writeSessionsToFile(Sessions)
	if err != nil {
		return model.Session{}, err
	}

	return updatedSession, nil
}

func (r *SessionRepoImpl) writeSessionsToFile(Sessions []model.Session) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(Sessions)
	if err != nil {
		return err
	}

	return nil
}

func generateUniqueIDSession(Sessions []model.Session) int64 {
	var maxID int64
	for _, Session := range Sessions {
		if Session.ID > maxID {
			maxID = Session.ID
		}
	}
	return maxID + 1
}

func NewSessionRepoImpl(filePath string) SessionRepo {
	return &SessionRepoImpl{
		filePath: filePath,
	}
}
