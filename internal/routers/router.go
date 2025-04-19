package routers

// import (
// 	"fmt"

// 	c "vtuanjs/my-project/internal/controller"
// 	"vtuanjs/my-project/internal/middlewares"

// 	"github.com/gin-gonic/gin"
// )

// func MiddlewareA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before middleware A")
// 		c.Next()
// 		fmt.Println("After middleware A")
// 	}
// }

// func MiddlewareB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before middleware B")
// 		c.Next()
// 		fmt.Println("After middleware B")
// 	}
// }

// func MiddlewareC() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before middleware C")
// 		c.Next()
// 		fmt.Println("After middleware C")
// 	}
// }

// func MiddlewareD(c *gin.Context) {
// 	fmt.Println("Before middleware D")
// 	c.Next()
// 	fmt.Println("After middleware D")
// }

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.Use(middlewares.AuthMiddleware(), MiddlewareA(), MiddlewareB(), MiddlewareC(), MiddlewareD)

// 	v1 := r.Group("/v1")
// 	{
// 		adminV1 := v1.Group("/admin")
// 		{
// 			adminV1.GET("/users/:id", c.NewUserController().GetUserByID)
// 		}

// 		v1.GET("/ping/:name", c.NewPongController().Pong)
// 		v1.GET("/users/:id", c.NewUserController().GetUserByID)
// 		// v1.PUT("/ping", Pong)
// 		// v1.PATCH("/ping", Pong)
// 		// v1.OPTIONS("/ping", Pong)
// 	}

// 	// v2 := r.Group("/v2")
// 	// {
// 	// v2.GET("/ping", Pong)
// 	// v2.POST("/ping", Pong)
// 	// v2.PUT("/ping", Pong)
// 	// v2.PATCH("/ping", Pong)
// 	// v2.OPTIONS("/ping", Pong)
// 	// }

// 	return r
// }
