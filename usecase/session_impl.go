package usecase

import (
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
)

type SessionUsecaseImpl struct {
	SessionRepo repository.SessionRepo
	UserRepo    repository.UserRepo
}

// Delete implements SessionUsecase
func (u *SessionUsecaseImpl) Delete(id int64) (model.Session, error) {
	return u.SessionRepo.Delete(id)
}

// FindAll implements SessionUsecase
func (u *SessionUsecaseImpl) FindAll() ([]model.Session, error) {
	return u.SessionRepo.FindAll()
}

// FindById implements SessionUsecase
func (u *SessionUsecaseImpl) FindById(id int64) (model.Session, error) {
	return u.SessionRepo.FindById(id)
}

// Save implements SessionUsecase
func (u *SessionUsecaseImpl) Save(newSession model.Session) (model.Session, error) {
	user, err := u.UserRepo.FindById(newSession.UserID)
	if err != nil {
		return model.Session{}, err
	}

	newSession.User = user

	return u.SessionRepo.Save(newSession)
}

// Update implements SessionUsecase
func (u *SessionUsecaseImpl) Update(updatedSession model.Session) (model.Session, error) {

	// Mendapatkan entitas Session sebelumnya dari SessionRepo berdasarkan ID
	previousSession, err := u.SessionRepo.FindById(updatedSession.ID)
	if err != nil {
		return model.Session{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousUserID := previousSession.UserID
	previousToken := previousSession.Token
	previousCreatedAt := previousSession.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedSession.UserID == 0 {
		updatedSession.UserID = previousUserID
	}
	if updatedSession.Token == "" {
		updatedSession.Token = previousToken
	}

	if updatedSession.CreatedAt == (time.Time{}) {
		updatedSession.CreatedAt = previousCreatedAt
	}

	user, err := u.UserRepo.FindById(updatedSession.UserID)
	if err != nil {
		return model.Session{}, err
	}

	updatedSession.User = user

	return u.SessionRepo.Update(updatedSession)
}

func NewSessionUsecaseImpl(SessionRepo repository.SessionRepo, UserRepo repository.UserRepo) SessionUsecase {
	return &SessionUsecaseImpl{
		SessionRepo: SessionRepo,
		UserRepo:    UserRepo,
	}
}
