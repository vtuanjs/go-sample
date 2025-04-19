package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgres()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
