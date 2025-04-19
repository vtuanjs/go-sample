//go:build wireinject
// +build wireinject

package wire

import (
	"vtuanjs/my-project/internal/controller"
	"vtuanjs/my-project/internal/repo"
	"vtuanjs/my-project/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
