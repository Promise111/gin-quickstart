package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func V4Login(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user", "Promise")
		session.Save()
		slog.Info("Request log", "request", *c.Request, "name", "Promise Ihunna")
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "logged in",
		})
	}
}

func V4Profile(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "user data fetched",
			"data":    user,
		})
	}
}

func V4Logout(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// session.Delete("user") // delete session key value
		session.Clear() // clear all session key value
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Logged out",
		})
	}
}
