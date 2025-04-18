package initialize

import (
	"fmt"

	"github.com/vtuanjs/my-project/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Load config Mysql ", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()

	r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
