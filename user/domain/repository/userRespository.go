package repository

import (
	"sync"

	"ddd/user/domain/model"

	"github.com/google/wire"
)

var Provider = wire.NewSet(NewUserRepository)

type UserRepository struct {
}

var (
	once sync.Once
	ur   *UserRepository
)

func NewUserRepository() *UserRepository {
	once.Do(func() {
		ur = &UserRepository{}
	})

	return ur
}

func (repo *UserRepository) GetUserById(uid int64) (*model.User, error) {
	return &model.User{
		Id:      uid,
		Name:    "hong",
		Balance: 100,
	}, nil
}

func (repo UserRepository) GetUserByName(userName string) (*model.User, error) {
	return &model.User{
		Id:      1,
		Name:    userName,
		Balance: 100,
	}, nil

}
