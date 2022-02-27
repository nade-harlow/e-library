package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/e-library/app/helper"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("student", helper.Session())
		c.Next()
	}
}
