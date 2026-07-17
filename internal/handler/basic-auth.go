package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// simulate some private keys
var secrets = gin.H{
	"foo":           gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin":        gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":          gin.H{"email": "lena@guapa.com", "phone": "523443"},
	"promiseihunna": gin.H{"email": "promiseihunna@gmail.com", "phone": "08090"},
}

func HandleSecrets(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		// user, exists := c.Get(gin.AuthUserKey)
		log.Printf("Auth user %v", user)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	}
}
