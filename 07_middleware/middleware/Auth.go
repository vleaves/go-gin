package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, e := c.Request.Cookie("user_cookie")
		if e == nil {
			c.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			c.Next()
		} else {
			c.Abort()
			c.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}
