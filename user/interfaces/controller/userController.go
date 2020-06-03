package controller

import (
	"ddd/user/application"
	"ddd/user/domain"
	"ddd/user/domain/repository"
	"ddd/user/interfaces/form"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uas  *application.UserAppService
	us   domain.UserService
	repo *repository.UserRepository
}

func NewUserController(uas *application.UserAppService, us domain.UserService, repo *repository.UserRepository) *UserController {
	c := &UserController{
		uas:  uas,
		us:   us,
		repo: repo,
	}
	return c
}

func (c UserController) Router(router *gin.RouterGroup) {
	router.POST("login", c.login)
}

func (c UserController) login(ctx *gin.Context) {
	var dto form.LoginDto
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.AbortWithError(200, err)
		return
	}




}
