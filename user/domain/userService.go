package domain

import (
	"ddd/user/domain/model"
	"ddd/user/domain/repository"
)

type UserService interface {
	GetUserById(uid int64) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (u userService) GetUserById(uid int64) (*model.User, error) {
	return u.repo.GetUserById(uid)
}
