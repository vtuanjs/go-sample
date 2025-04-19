package global

import (
	"vtuanjs/my-project/pkg/logger"
	"vtuanjs/my-project/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	RDB    *redis.Client
	PDB    *gorm.DB
)
