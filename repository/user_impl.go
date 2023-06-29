package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sferawann/test_mnc/model"
)

type UserRepoImpl struct {
	filePath string
}

// Delete implements UserRepo
func (r *UserRepoImpl) Delete(id int64) (model.User, error) {
	users, err := r.FindAll()
	if err != nil {
		return model.User{}, err
	}

	var deletedUser model.User
	for i, user := range users {
		if user.ID == id {
			deletedUser = user
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	err = r.writeUsersToFile(users)
	if err != nil {
		return model.User{}, err
	}

	return deletedUser, nil
}

// FindAll implements UserRepo
func (r *UserRepoImpl) FindAll() ([]model.User, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var users []model.User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return users, nil
}

// FindById implements UserRepo
func (r *UserRepoImpl) FindById(id int64) (model.User, error) {
	users, err := r.FindAll()
	if err != nil {
		return model.User{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("user by id: %d not found", id)
}

// FindByUsername implements UserRepo
func (r *UserRepoImpl) FindByUsername(username string) (model.User, error) {
	users, err := r.FindAll()
	if err != nil {
		return model.User{}, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("user by username: %s not found", username)
}

// Save implements UserRepo
func (r *UserRepoImpl) Save(newUser model.User) (model.User, error) {
	users, err := r.FindAll()
	if err != nil {
		return model.User{}, err
	}

	newUser.ID = generateUniqueIDUser(users)
	newUser.CreatedAt = time.Now()

	users = append(users, newUser)

	err = r.writeUsersToFile(users)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

// Update implements UserRepo
func (r *UserRepoImpl) Update(updatedUser model.User) (model.User, error) {
	users, err := r.FindAll()
	if err != nil {
		return model.User{}, err
	}

	var found bool
	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i] = updatedUser
			found = true
			break
		}
	}

	if !found {
		return model.User{}, fmt.Errorf("user by id: %d not found", updatedUser.ID)
	}

	err = r.writeUsersToFile(users)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (r *UserRepoImpl) writeUsersToFile(users []model.User) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(users)
	if err != nil {
		return err
	}

	return nil
}

func generateUniqueIDUser(users []model.User) int64 {
	var maxID int64
	for _, user := range users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	return maxID + 1
}

func NewUserRepoImpl(filePath string) UserRepo {
	return &UserRepoImpl{
		filePath: filePath,
	}
}
