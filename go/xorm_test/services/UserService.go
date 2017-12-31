package services

import (
	"errors"
	"fmt"
	"time"

	"../datamodels"
	"../repositories"
)

type UserService interface {
	Create(password string, new_user datamodels.User) (datamodels.User, error)
	GetById(id int64) datamodels.User
	GetByUsernameAndPassword(username string, password string) (datamodels.User, bool)
}
type userService struct {
	repo repositories.UserRepository
	//UserModel *datamodels.User
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
func (this *userService) Create(password string, new_user datamodels.User) (datamodels.User, error) {
	fmt.Println("is here?------------------------------>")
	if new_user.Id > 0 || new_user.Firstname == "" || new_user.Lastname == "" || new_user.Username == "" || new_user.Email == "" || password == "" {
		return new_user, errors.New("unable to create this user")

	}
	hashed, err := new_user.GeneratePassword(password)

	if err != nil {
		return new_user, err
	}

	new_user.Password = hashed
	new_user.CreatedAt = time.Now()
	fmt.Println("HERE CREATE SERVICE------------------------------>")
	return this.repo.Store(new_user)
}
