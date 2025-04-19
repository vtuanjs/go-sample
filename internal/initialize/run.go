package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgres()
	InitPostgresC()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
