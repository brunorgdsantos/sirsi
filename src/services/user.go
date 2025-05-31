package services

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/utils"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(email, password string) error {
	user := dtos.User{
		Email:    email,
		Password: password,
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err := service.repo.Create(contextServer, user)
	if err != nil {
		return err
	}
	return nil
}
