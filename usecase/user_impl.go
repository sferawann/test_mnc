package usecase

import (
	"time"

	"github.com/sferawann/test_mnc/model"
	"github.com/sferawann/test_mnc/repository"
	"github.com/sferawann/test_mnc/utils"
)

type UserUsecaseImpl struct {
	UserRepo repository.UserRepo
}

// Delete implements UserUsecase
func (u *UserUsecaseImpl) Delete(id int64) (model.User, error) {
	return u.UserRepo.Delete(id)
}

// FindAll implements UserUsecase
func (u *UserUsecaseImpl) FindAll() ([]model.User, error) {
	return u.UserRepo.FindAll()
}

// FindById implements UserUsecase
func (u *UserUsecaseImpl) FindById(id int64) (model.User, error) {
	return u.UserRepo.FindById(id)
}

// FindByUsername implements UserUsecase
func (u *UserUsecaseImpl) FindByUsername(username string) (model.User, error) {
	return u.UserRepo.FindByUsername(username)
}

// Save implements UserUsecase
func (u *UserUsecaseImpl) Save(newUser model.User) (model.User, error) {
	// Validasi username
	if err := utils.ValidateUsernameMinLength(newUser.Username); err != nil {
		return model.User{}, err
	}

	// Validasi password
	if err := utils.ValidatePasswordMinLength(newUser.Password); err != nil {
		return model.User{}, err
	}

	// Validasi email
	if err := utils.ValidateEmailFormat(newUser.Email); err != nil {
		return model.User{}, err
	}
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return model.User{}, err
	}
	newUser.Password = hashedPassword

	return u.UserRepo.Save(newUser)
}

// Update implements UserUsecase
func (u *UserUsecaseImpl) Update(updatedUser model.User) (model.User, error) {
	hashedPassword, err := utils.HashPassword(updatedUser.Password)
	if err != nil {
		return model.User{}, err
	}

	// Mendapatkan entitas User sebelumnya dari UserRepo berdasarkan ID
	previousUser, err := u.UserRepo.FindById(updatedUser.ID)
	if err != nil {
		return model.User{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousUsername := previousUser.Username
	previousPassword := previousUser.Password
	previousEmail := previousUser.Email
	previousCreatedAt := previousUser.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedUser.Username == "" {
		updatedUser.Username = previousUsername
	}
	if updatedUser.Password == "" {
		updatedUser.Password = previousPassword
	}
	if updatedUser.Email == "" {
		updatedUser.Email = previousEmail
	}
	if updatedUser.CreatedAt == (time.Time{}) {
		updatedUser.CreatedAt = previousCreatedAt
	}

	// Hash password baru jika ada perubahan
	if updatedUser.Password != previousPassword {
		updatedUser.Password = hashedPassword
	}

	return u.UserRepo.Update(updatedUser)
}

func NewUserUsecaseImpl(UserRepo repository.UserRepo) UserUsecase {
	return &UserUsecaseImpl{
		UserRepo: UserRepo,
	}
}
