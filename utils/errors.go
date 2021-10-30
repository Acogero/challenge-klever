package utils

import (
	"github.com/gin-gonic/gin"
)

func Errors(httpStatus int, message string, err error, c *gin.Context) {
	c.JSON(httpStatus, gin.H{
		"message": message,
		"error":   err,
	})
}
