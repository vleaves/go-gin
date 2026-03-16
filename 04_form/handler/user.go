package hanlder

import (
	"06_tmpl/model"
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
		println("err ->", err.Error())
		return
	}
	println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)

	// email := c.PostForm("email")
	// password := c.DefaultPostForm("password", "Wa123456")
	// passwordAgain := c.DefaultPostForm("password-again", "Wa123456")
	// println("email", email, "password", password, "password again", passwordAgain)
}
