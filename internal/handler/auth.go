package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=8,max=16"`
}

func LoginHandler(enging *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		log.Printf("%+v", json)

		if json.Email != "Promise" && json.Password != "passPass" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "You are logged in",
		})
	}
}
