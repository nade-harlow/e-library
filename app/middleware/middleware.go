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
