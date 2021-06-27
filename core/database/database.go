package database

import (
	"log"
	"sync"

	"whats/core/config"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	// DB MysqlDB
	DB       *gorm.DB
	initOnce sync.Once
)

// InitDB InitDB
func InitDB() {
	initOnce.Do(func() {
		var dialects gorm.Dialector
		switch config.GetEnv("DB_CONNECTION") {
		case "mysql":
			dialects = mysql.New(mysql.Config{
				DSN: config.GetEnv("DB_DSN"),
			})
		case "sqlite":
			dialects = sqlite.Open(config.GetEnv("DB_DSN"))
		case "postgres":
			dialects = postgres.Open(config.GetEnv("DB_DSN"))
		case "sqlserver":
			dialects = sqlserver.Open(config.GetEnv("DB_DSN"))
		case "clickhouse":
			dialects = clickhouse.Open(config.GetEnv("DB_DSN"))
		default:
			log.Println("Not setting db args or not support DB_CONNECTION.")
			return
		}
		DB, err = gorm.Open(dialects, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			PrepareStmt: true,
		})
		if err != nil {
			log.Panicf("init DB err:%v", err)
		}
		db, err := DB.DB()
		if err != nil {
			log.Panicf("MysqlDB.DB() err:%v", err)
		}
		db.SetMaxIdleConns(100)
		db.SetMaxOpenConns(100)
		migrate(DB)
	})
}
