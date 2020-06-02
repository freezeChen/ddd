package repository

import (
	"ddd/user/domain/model"
)

type UserRepository struct {
}

func (repo *UserRepository) GetUserById(uid int64) (*model.User, error) {

	return &model.User{
		Id:      uid,
		Name:    "hong",
		Balance: 100,
	}, nil
}
