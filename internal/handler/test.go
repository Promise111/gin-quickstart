package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = "/api/v1/final"
		engine.HandleContext(c)
	}
}

func FinalHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Hello World",
	})
}
