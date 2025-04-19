package initialize

import (
	"fmt"
	"time"
	"vtuanjs/my-project/global"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitPostgres() {
	m := global.Config.Postgres
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	var s = fmt.Sprintf(dsn, m.Host, m.Username, m.Password, m.DBName, m.Port)
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Init Postgres failed")
	global.Logger.Info("Init Postgres success")
	global.PDB = db
	SetPool()
	// migrationTables()
}

func SetPool() {
	m := global.Config.Postgres
	sqlDB, err := global.PDB.DB()

	if err != nil {
		global.Logger.Error("Set Pool failed", zap.Error(err))
		panic(err)
	}

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

// func migrationTables() {
// 	global.PDB.AutoMigrate(
// 	// &models.User{},
// 	// &models.Role{},
// 	)
// }
