package manager

import (
	"vtuanjs/my-project/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/users")
	adminRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		adminRouterPrivate.POST("/active")
	}
}
