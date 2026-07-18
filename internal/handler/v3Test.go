package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func V3Ping(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "pong",
		})
	}
}
