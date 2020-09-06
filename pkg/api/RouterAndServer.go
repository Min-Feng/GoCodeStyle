package api

import (
	"github.com/gin-gonic/gin"

	"ddd/pkg/helper/helperlog"
)

func NewServer(address string, router *gin.Engine) *Server {
	return &Server{address: address, router: router}
}

type Server struct {
	address string
	router  *gin.Engine
}

func (s *Server) RegisterHandler(dHandler *DebugHandler) {
	s.router.PUT("debug/logLevel", dHandler.UpdateLogLevel)
}

func (s *Server) Start() error {
	return s.router.Run(s.address)
}

func NewRouter(logLevel helperlog.Level) *gin.Engine {
	switch logLevel {
	case helperlog.TraceLevel, helperlog.DebugLevel:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	router := gin.New()

	// c.Next() 之前, 初始化順序, 右到左
	// c.Next() 之後, 執行順序, 左到右, 也就是說 gin.Logger() 最後執行
	router.Use(gin.Logger(), gin.Recovery(), ErrorResponseMiddleware)

	return router
}
