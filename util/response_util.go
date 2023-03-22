package util

import "github.com/gin-gonic/gin"

func SendResponse(c *gin.Context, code int, message string, success bool, data any) {
	c.JSON(code, gin.H{
		"succes":  success,
		"message": message,
		"data":    data,
	})
}
