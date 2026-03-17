package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "06_file/handler"
	"06_file/utils"
)

func InitRouter() *gin.Engine {
	eng := gin.Default()
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
		userRouter.GET("/profile/", handler.UserProfile)
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	return eng
}
