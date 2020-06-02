package interfaces

import "github.com/google/wire"

type Controller struct {
}

func InitController() Controller {

	wire.Build()
}
