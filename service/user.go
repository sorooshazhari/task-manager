package service

import (
	"fmt"
	"task-manager/entity"
)

type ServiceUserRepo interface {
	CreateNewUser(entity.User) (entity.User, error)
	ValidateUser(email string, pass string) bool
}

type UserService struct {
	userRepo ServiceUserRepo
}

func NewUserService(uRepo ServiceUserRepo) UserService {
	return UserService{userRepo: uRepo}
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	User entity.User
}

type ValidateUserRequest struct {
	Email    string
	Password string
}

type ValidateUserResponse struct {
	IsValidate bool
}

func (u UserService) Register(uReq CreateUserRequest) (CreateUserResponse, error) {
	user, cErr := u.userRepo.CreateNewUser(entity.User{
		Email:    uReq.Email,
		Name:     uReq.Name,
		Password: uReq.Password,
	})
	if cErr != nil {
		return CreateUserResponse{}, fmt.Errorf("we can't create user with %v info", uReq)
	}
	return CreateUserResponse{User: user}, nil
}
func (u UserService) Login(vReq ValidateUserRequest) ValidateUserResponse {
	isOk := u.userRepo.ValidateUser(vReq.Email, vReq.Password)
	if isOk {
		return ValidateUserResponse{IsValidate: true}
	}
	return ValidateUserResponse{IsValidate: false}
}
