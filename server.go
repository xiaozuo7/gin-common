package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"com.github.gin-common/common/gin_recovery"

	"go.uber.org/zap"

	"com.github.gin-common/common/loggers/gin_logger"
	"com.github.gin-common/common/routers"
	"com.github.gin-common/common/validator_trans"

	"com.github.gin-common/util"

	"com.github.gin-common/app/router"

	"com.github.gin-common/migrate"

	_ "github.com/joho/godotenv/autoload"
)

var routerConfigs = []routers.GinRouterInterface{
	router.UserRouter{},
	router.AuthRouter{},
	router.BookRouter{},
}

func setGinMode() {
	currentEnv := util.GetDefaultEnv("RUN_ENV", "dev")
	if currentEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else if currentEnv == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "migrate" {
		migrate.Migrate()
		return
	}
	setGinMode()
	_ = validator_trans.InitTrans(util.GetDefaultEnv("LOCALE", "zh"))
	r := gin.New()
	ginLogLevel := util.GetDefaultEnv("GIN_LOG_LEVEL", "error")
	// 加载日志中间件
	r.Use(gin_logger.LoggerWithWriter(gin.DefaultWriter, util.GetLogLevel(ginLogLevel), []zap.Option{}))
	// 加载recover中间件
	r.Use(gin_recovery.Recovery())


	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		gin_logger.Log.Info("", zap.String("httpMethod", httpMethod), zap.String("absolutePath", absolutePath))
	}
	routers.CombineRouters(r, routerConfigs...)

	if err := r.Run(":8080"); err != nil {
		gin_logger.Log.Fatal(err.Error())
	}
}
