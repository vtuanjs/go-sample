package initialize

import (
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/internal/middlewares"
	"vtuanjs/my-project/internal/routers"

	"github.com/gin-gonic/gin"
)

func MiddlewareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	r.Use(middlewares.NewRateLimiter().GlobalRateLimiter())
	r.Use(MiddlewareA())
	// r.Use()
	// r.Use()

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("checkStatus")
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		managerRouter.InitUserRouter(mainGroup)
		managerRouter.InitAdminRouter(mainGroup)
	}

	return r
}
