package hanlder

import (
	"05_sql/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSave(c *gin.Context) {
	username := c.Param("name")
	c.String(http.StatusOK, "用户"+username+"已保存")
}

func UserSaveByQuery(c *gin.Context) {
	username := c.Query("name")
	// age := c.Query("age")
	age := c.DefaultQuery("age", "20")
	c.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		c.String(http.StatusBadRequest, "输入的数据不合法")
	}
	id := user.Save()
	log.Println("id is ", id)
	c.Redirect(http.StatusMovedPermanently, "/")
}
