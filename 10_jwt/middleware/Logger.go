package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		url := c.Request.URL
		method := c.Request.Method
		fmt.Printf("%s::%s \t %s \t %s \n", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		c.Next()
		fmt.Println(c.Writer.Status())
	}
}
