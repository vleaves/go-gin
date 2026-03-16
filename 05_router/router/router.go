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
	router.GET("/", retHelloGinAndMethod)
	router.POST("/", retHelloGinAndMethod)
	router.PUT("/", retHelloGinAndMethod)
	router.DELETE("/", retHelloGinAndMethod)
	router.PATCH("/", retHelloGinAndMethod)
	router.HEAD("/", retHelloGinAndMethod)
	router.OPTIONS("/", retHelloGinAndMethod)

	router.GET("/user/:name", handler.UserSave)
	router.GET("/user", handler.UserSaveByQuery)
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin get method")
	// })
	// router.POST("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin post method")
	// })
	// router.PUT("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin put method")
	// })
	// router.DELETE("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin delete method")
	// })
	// router.PATCH("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin patch method")
	// })
	// router.HEAD("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin head method")
	// })
	// router.OPTIONS("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello gin options method")
	// })
	return router
}
