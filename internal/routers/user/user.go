package user

import (
	"vtuanjs/my-project/internal/controller/account"
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
		userRouterPublic.POST("/login", account.Login.Login)
	}

	// private router
	userRouterPrivate := Router.Group("/users")
	userRouterPrivate.Use(middlewares.NewRateLimiter().UserAndPrivateAPIRateLimiter())
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/me")
		userRouterPrivate.POST("/test-private", account.Login.Login)
	}
}
