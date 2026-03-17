package router

import (
	"github.com/gin-gonic/gin"

	handler "04_form/handler"
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
	index := eng.Group("/")
	{
		index.Any("", handler.Index)
	}
	// 添加 user
	userRouter := eng.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}

	return eng
}
