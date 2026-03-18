package test

import (
	"08_restful/model"
	"08_restful/router"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var eng *gin.Engine

func init() {
	eng = router.InitRouter()
}

func TestArticleInsert(t *testing.T) {
	article := model.Article{}
	article.Type = "go"
	article.Content = "Hello Go"
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article", bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-typ", "application/json")
	eng.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, "{id:-1}", w.Body.String())
}
