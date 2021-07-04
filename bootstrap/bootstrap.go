package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"whats/core"
	"whats/core/cache"
	"whats/core/config"
	"whats/core/database"
	"whats/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// init 初始化 需要的数据
func init() {
	if !config.GetEnvToBool("APP_DEBUG") {
		gin.SetMode(gin.ReleaseMode)
	}
	// init database
	database.InitAll()
	cache.Init()
}

// Run 开始
func Run() {
	engine := gin.New()

	//f, _ := os.Create(config.GetDefaultEnv("LOG_FILE", "gin.log"))
	//gin.DefaultWriter = io.MultiWriter(f)
	engine.Use(gin.Logger())

	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	engine.Use(cors.Default())

	engine.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusOK, core.NewE(500, fmt.Sprintf("System exception:%s", recovered)))
		c.Abort()
	}))

	// load  routers
	router.Routers(engine)
	err := engine.Run(":" + config.GetDefaultEnv("HTTP_PORT", "1993"))
	if err != nil {
		log.Panic(err)
	}
}
