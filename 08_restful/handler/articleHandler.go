package handler

import (
	_ "08_restful/docs"
	"08_restful/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 提交新的文章内容
// @Id 1
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 添加失败
// @Router /article [post]
func Insert(c *gin.Context) {
	article := model.Article{}
	var id = -1
	if e := c.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
