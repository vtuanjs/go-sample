package initialize

import (
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
