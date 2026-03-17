package hanlder

import (
	"06_file/model"
	"06_file/utils"
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

func UserLogin(c *gin.Context) {
	var user model.UserModel
	if e := c.ShouldBind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id":    u.Id,
		})
	}
}

func UserProfile(c *gin.Context) {
	id := c.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	c.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误 ", err.Error())
	}
	file, e := c.FormFile("avatar-file")
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}

	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	log.Println(path)
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	finalPath := filepath.Join(path, fileName)
	e = c.SaveUploadedFile(file, finalPath)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	avatarUrl := "http://127.0.0.1:8080/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(user.Id)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))
}
