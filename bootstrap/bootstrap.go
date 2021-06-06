package bootstrap

import (
	"github.com/gin-gonic/gin"
	"log"
	"whats/core/cache"

	"whats/core/config"
	"whats/core/database"
	"whats/router"
)

var (
	ginMode = "release"
)

// init 初始化 需要的数据
func init() {
	if config.GetEnvToBool("APP_DEBUG") {
		ginMode = "debug"
	}
}

// Run 开始
func Run() {
	//cron.InitCron()
	// 初始化静态资源到对象存储
	// init database
	database.InitMysql()
	cache.Init()

	gin.SetMode(ginMode)
	ginEngine := gin.Default()
	// use default settings
	//ginEngine.Use(core.Recover(), middleware.Cors())

	// load  routers
	engine := router.Routers(ginEngine)
	err := engine.Run(":" + config.GetDefaultEnv("HTTP_PORT", "1993"))
	if err != nil {
		log.Panic(err)
	}
}
