package router

import (
	"08_restful/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	articleRouter := router.Group("")
	{
		// // 通过获取单篇文章
		// articleRouter.GET("/article/:id", article.GetOne)
		// // 获取所有文章
		// articleRouter.GET("/articles", article.GetAll)
		// // 添加一篇文章
		articleRouter.POST("/article", handler.Insert)
		// articleRouter.DELETE("/article/:id", article.DeleteOne)
	}

	return router
}
