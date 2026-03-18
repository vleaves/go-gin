package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "07_middleware/handler"
	"07_middleware/middleware"
	"07_middleware/utils"
)

func InitRouter() *gin.Engine {
	eng := gin.New()
	// eng := gin.Default()
	eng.Use(middleware.Logger(), gin.Recovery())
	if mode := gin.Mode(); mode == gin.TestMode {
		eng.LoadHTMLGlob("./../templates/*")
	} else {
		eng.LoadHTMLGlob("templates/*")
	}
	eng.StaticFile("/favicon.ico", "./favicon.ico")
	eng.Static("/statics", "./statics")
	eng.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	index := eng.Group("/")
	{
		index.Any("", handler.Index)
	}
	// 添加 user
	userRouter := eng.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	return eng
}
