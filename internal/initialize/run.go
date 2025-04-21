package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitPostgres()
	InitPostgresC()
	InitServiceInterface()
	InitRedis()
	InitKafka()

	r := InitRouter()
	return r
}
