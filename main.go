package main

import (
	"ddd/user/di"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	controller := di.InitUserController()

	group := engine.Group("user")
	controller.Router(group)

	engine.Run()
}
