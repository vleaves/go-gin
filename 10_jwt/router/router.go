package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	handler "10_jwt/handler"
	"10_jwt/middleware"
)

func InitRouter() *gin.Engine {
	eng := gin.New()
	// eng := gin.Default()
	eng.Use(middleware.Logger(), gin.Recovery())
	eng.StaticFile("/favicon.ico", "./favicon.ico")
	eng.GET("/", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})
	// 添加 user
	userRouter := eng.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.CreateJwt)
	}

	return eng
}
