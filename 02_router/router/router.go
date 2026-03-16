package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	handler "05_router/handler"
)

func retHelloGinAndMethod(c *gin.Context) {
	c.String(http.StatusOK, "hello gin "+strings.ToLower(c.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	index := router.Group("/")
	{
		index.Any("", retHelloGinAndMethod)
	}
	user := router.Group("/user")
	{
		user.GET("/:name", handler.UserSave)
		user.GET("", handler.UserSaveByQuery)
	}
	return router
}
