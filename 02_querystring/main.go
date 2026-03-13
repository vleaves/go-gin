package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//name := c.Query("query")
		name := c.DefaultQuery("query", "没有取到值")
		c.JSON(200, gin.H{
			"message": name,
		})
	})

	r.Run()
}
