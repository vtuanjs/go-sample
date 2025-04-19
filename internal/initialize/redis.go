package initialize

import (
	"context"
	"fmt"
	"vtuanjs/my-project/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Init Redis failed", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Init Redis success")
	global.RDB = rdb
}
