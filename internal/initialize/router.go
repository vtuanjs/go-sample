package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/vtuanjs/my-project/internal/controller"
	"github.com/vtuanjs/my-project/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())

	v1 := r.Group("/v1")
	{
		adminV1 := v1.Group("/admin")
		{
			adminV1.GET("/users/:id", c.NewUserController().GetUserByID)
		}

		v1.GET("/ping/:name", c.NewPongController().Pong)
		v1.GET("/users/:id", c.NewUserController().GetUserByID)
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	// v2 := r.Group("/v2")
	// {
	// v2.GET("/ping", Pong)
	// v2.POST("/ping", Pong)
	// v2.PUT("/ping", Pong)
	// v2.PATCH("/ping", Pong)
	// v2.OPTIONS("/ping", Pong)
	// }

	return r
}
