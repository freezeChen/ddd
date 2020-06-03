//+build wireinject

package di

import (
	"ddd/user/application"
	"ddd/user/domain"
	"ddd/user/domain/repository"
	"ddd/user/interfaces/controller"

	"github.com/google/wire"
)


//go:generate wire
func InitUserController() *controller.UserController {
	panic(wire.Build(repository.Provider, domain.New,application.NewUserAppService, controller.NewUserController))
}
