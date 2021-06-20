package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"sync"
	"whats/core/config"
)

var (
	// MysqlDB MysqlDB
	MysqlDB  *gorm.DB
	initOnce sync.Once
)

// InitMysql InitMysql
func InitMysql() {
	initOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
			config.GetDefaultEnv("DB_USERNAME", "root"),
			config.GetDefaultEnv("DB_PASSWORD", "123456"),
			config.GetDefaultEnv("DB_HOST", "127.0.0.1"),
			config.GetDefaultEnv("DB_PORT", "3306"),
			config.GetDefaultEnv("DB_DATABASE", "whats"),
			strings.Replace(config.GetDefaultEnv("DB_TIMEZONE", "Local"), "/", "%2F", -1),
		)
		MysqlDB, err = gorm.Open(mysql.New(mysql.Config{
			//DriverName:                "",
			DSN: dsn,
			//Conn:                      nil,
			//SkipInitializeWithVersion: false,
			//DefaultStringSize:         0,
			//DisableDatetimePrecision:  false,
			//DontSupportRenameIndex:    false,
			//DontSupportRenameColumn:   false,
		}), &gorm.Config{
			//SkipDefaultTransaction:                   false,
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "",
				SingularTable: true,
			},
			//Logger:                                   nil,
			//NowFunc:                                  nil,
			//DryRun:                                   false,
			PrepareStmt: true,
			//DisableAutomaticPing:                     false,
			//DisableForeignKeyConstraintWhenMigrating: false,
			//ClauseBuilders:                           nil,
			//ConnPool:                                 nil,
			//Dialector:                                nil,
			//Plugins:                                  nil,
		})
		if err != nil {
			log.Panicf("init MysqlDB err:%v,dsn:%v", err, dsn)
		}

		db, err := MysqlDB.DB()
		if err != nil {
			log.Panicf("MysqlDB.DB() err:%v,dsn:%v", err, dsn)
		}
		db.SetMaxIdleConns(100)
		db.SetMaxOpenConns(100)

		migrate(MysqlDB)
	})
}
