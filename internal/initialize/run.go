package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgres()
	InitPostgresC()
	InitRedis()
	InitKafka()

	r := InitRouter()
	r.Run(":8002")
}
