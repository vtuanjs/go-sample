package initialize

import (
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/internal/service"
	"vtuanjs/my-project/internal/service/impl"
)

func InitServiceInterface() {
	service.InitUserLogin(impl.NewUserLogin(global.PDBC))
	service.InitUserAdmin(impl.NewUserAdmin(global.PDBC))
	service.InitUserInfo(impl.NewUserInfo(global.PDBC))
}
