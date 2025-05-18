package service

import (
	"go-template/internal/repository"
)

type IUserService interface {
	GetUser() string
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetUser() string {
	user := us.userRepo.FetchUser()
	return user
}
