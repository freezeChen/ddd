package domain

import (
	"ddd/user/domain/repository"
)

type (
	UserService interface {
	}

	userService struct {
		repo *repository.UserRepository
	}
)

func New(repo *repository.UserRepository) UserService {

	return &userService{
		repo: repo,
	}
}
