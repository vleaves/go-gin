package main

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		var u UserInfo
		c.ShouldBind(&u)
		c.JSON(200, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	})

	r.Run()
}
