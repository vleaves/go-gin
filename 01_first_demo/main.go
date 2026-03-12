package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("./templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(200, "posts_index.tmpl", gin.H{
			"message": "我是posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users_index.tmpl", gin.H{
			"message": "<a href='https://www.baidu.com'>百度主页</a>",
		})
	})
	r.Run()
}
