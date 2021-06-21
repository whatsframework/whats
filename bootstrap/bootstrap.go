package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"whats/core/cache"
	"whats/core/database"

	"whats/core"
	"whats/core/config"
	"whats/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	ginMode = "release"
)

// init 初始化 需要的数据
func init() {
	if config.GetEnvToBool("APP_DEBUG") {
		ginMode = "debug"
	}
	// init database
	database.InitMysql()
	cache.Init()
}

// Run 开始
func Run() {

	gin.SetMode(ginMode)
	ginEngine := gin.New()

	//f, _ := os.Create(config.GetDefaultEnv("LOG_FILE", "gin.log"))
	//gin.DefaultWriter = io.MultiWriter(f)
	ginEngine.Use(gin.Logger())

	ginEngine.Use(gzip.Gzip(gzip.DefaultCompression))

	ginEngine.Use(cors.Default())

	ginEngine.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusOK, core.NewE(500, fmt.Sprintf("System exception:%s", recovered)))
		c.Abort()
	}))

	// load  routers
	engine := router.Routers(ginEngine)
	err := engine.Run(":" + config.GetDefaultEnv("HTTP_PORT", "1993"))
	if err != nil {
		log.Panic(err)
	}
}
