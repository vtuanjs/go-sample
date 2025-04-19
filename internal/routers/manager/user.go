package manager

import (
	"vtuanjs/my-project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private router
	UserRouterPrivate := Router.Group("/admin/users")
	UserRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		UserRouterPrivate.GET("")
	}
}
