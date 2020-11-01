package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ddd/pkg/drivingAdapter/api/operation"
	"ddd/pkg/drivingAdapter/api/shared"
	"ddd/pkg/technical/logger"
)

func RegisterHandler(r *Router, dHandler *operation.DebugHandler) {
	r.router.PUT("debug/logLevel", dHandler.UpdateLogLevel)
}

func NewRouter(address string, logLevel logger.Level) *Router {
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
	router.Use(gin.Logger(), gin.Recovery(), shared.ErrorResponseMiddleware)

	return &Router{address: address, router: router}
}

type Router struct {
	address string
	router  *gin.Engine
}

func (r *Router) QuicklyStart() error {
	return r.router.Run(r.address)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
