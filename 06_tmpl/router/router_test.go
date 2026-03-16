package router

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = InitRouter()
}

func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
}

func TestUserSave(t *testing.T) {
	username := "lisi"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已保存", w.Body.String())
}

func TestUserSaveQuery(t *testing.T) {
	username := "lisi"
	age := 18

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}
