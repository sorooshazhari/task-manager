package repository

import (
	"fmt"
	"task-manager/entity"
)

type UserPermStore interface {
	Save(entity.User) error
	Load() ([]entity.User, error)
}

type UserStorage struct {
	permStore UserPermStore
	users     []entity.User
}

func NewUserStorage(pStore UserPermStore) UserStorage {
	return UserStorage{
		users:     make([]entity.User, 0),
		permStore: pStore,
	}
}

func (uStore *UserStorage) LoadUsers() error {
	users, lErr := uStore.permStore.Load()
	if lErr != nil {
		return lErr
	}
	uStore.users = users
	return nil
}

func (uStore *UserStorage) CreateNewUser(user entity.User) (entity.User, error) {
	for _, u := range uStore.users {
		if u.Email == user.Email {
			return entity.User{}, fmt.Errorf("you have registered before with this email")
		}
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return entity.User{}, fmt.Errorf("you have missed some fields")
	}
	newUser := entity.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		ID:       len(uStore.users) + 1,
	}
	uStore.users = append(uStore.users, newUser)
	uStore.permStore.Save(newUser)
	return newUser, nil
}

func (uStore *UserStorage) ValidateUser(email, password string) int {
	for _, u := range uStore.users {
		if u.Email == email {
			if u.Password == password {
				return u.ID
			} else {
				fmt.Println("password is not correct")
				return 0
			}
		}
	}
	fmt.Println("email is not registered")
	return 0
}
