package routers

import (
	"vtuanjs/my-project/internal/routers/manager"
	"vtuanjs/my-project/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
