package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session")
		if err != nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Set("student", cookie)
		c.Next()
	}
}

func CheckNotLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session")
		if err != nil {
			c.Next()
			return
		}

		c.Redirect(http.StatusFound, "/library/book/get-all-books")
	}
}
