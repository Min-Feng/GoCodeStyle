package api

import (
	"github.com/gin-gonic/gin"

	"ddd/pkg/technical/logger"
)

func NewRouter(logLevel string) Router {
	switch logLevel {
	case logger.TraceLevel, logger.DebugLevel:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	router := gin.New()

	// c.Next() 之前, 初始化順序, 右到左
	// c.Next() 之後, 執行順序, 左到右, 也就是說 gin.Logger() 最後執行
	router.Use(gin.Logger(), gin.Recovery(), ErrorResponseMiddleware)

	return Router{router}
}

type Router struct {
	*gin.Engine
}
