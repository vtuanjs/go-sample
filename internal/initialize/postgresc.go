package initialize

import (
	"context"
	"fmt"
	"vtuanjs/my-project/global"
	"vtuanjs/my-project/internal/database"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitPostgresC() {
	m := global.Config.Postgres
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	var s = fmt.Sprintf(dsn, m.Host, m.Username, m.Password, m.DBName, m.Port)
	// db, err := sql.Open("pgx/v5", s)
	conn, err := pgx.Connect(context.Background(), s)
	checkErrorPanicC(err, "Init PostgresC failed")
	global.Logger.Info("Init PostgresC success")
	global.PDBC = database.New(conn)

	// SetPool()
	// migrationTables()
}

// func SetPoolC() {
// 	m := global.Config.Postgres
// 	sqlDB := global.PDBC

// 	// if err != nil {
// 	// 	global.Logger.Error("Set PoolC failed", zap.Error(err))
// 	// 	panic(err)
// 	// }

// 	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
// 	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
// 	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
// }

// func migrationTables() {
// 	global.PDB.AutoMigrate(
// 	// &models.User{},
// 	// &models.Role{},
// 	)
// }
