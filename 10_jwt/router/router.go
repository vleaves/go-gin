package router

import (
	"github.com/gin-gonic/gin"

	handler "10_jwt/handler"
	"10_jwt/middleware"
)

func InitRouter() *gin.Engine {
	eng := gin.New()
	// eng := gin.Default()
	eng.Use(middleware.Logger(), gin.Recovery())
	eng.StaticFile("/favicon.ico", "./favicon.ico")
	index := eng.Group("/")
	{
		index.Any("", handler.Index)
	}
	// 添加 user
	userRouter := eng.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.CreateJwt)
	}

	return eng
}
