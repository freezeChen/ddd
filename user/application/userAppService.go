package application

import (
	"ddd/user/domain"
	"ddd/user/domain/model"
	"ddd/user/domain/repository"
	"ddd/user/interfaces/form"
)

type UserAppService struct {
	us   domain.UserService
	repo *repository.UserRepository
}

func NewUserAppService(us domain.UserService, repo *repository.UserRepository) *UserAppService {
	return &UserAppService{
		us:   us,
		repo: repo,
	}
}

func (s UserAppService) Login(form form.LoginDto) *model.User {

	s.repo.GetUserByName()

}
