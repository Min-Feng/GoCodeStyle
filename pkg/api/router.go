package api

import (
	"github.com/gin-gonic/gin"

	"ddd/pkg/helper/helperlog"
)

func NewRouter(logLevel helperlog.Level) *gin.Engine {
	switch logLevel {
	case helperlog.TraceLevel, helperlog.DebugLevel:
		gin.SetMode(gin.DebugMode)
	case helperlog.InfoLevel:
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	router := gin.New()

	// 初始化順序, 右到左
	// 執行順序, 左到右, 也就是說 gin.Logger() 最後執行
	router.Use(gin.Logger(), gin.Recovery(), ErrorResponseMiddleware)

	return router
}
