package user

import (
	"vtuanjs/my-project/internal/middlewares"
	"vtuanjs/my-project/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	// public router
	userRouterPublic := Router.Group("/users")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/users")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/me")
	}
}
