package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var BasicAuth = gin.BasicAuth(gin.Accounts{
	"foo":           "bar",
	"austin":        "1234",
	"lena":          "hello2",
	"manu":          "4321",
	"promiseihunna": "5401",
})

// my hands on middleware trial

func DumDumMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		log.Printf("dumdummiddleware %v", t)

		log.Printf("dumdum-middleware auth user %v", c.MustGet(gin.AuthUserKey).(string))

		c.Next()

		log.Printf("dumdummiddleware %v", time.Since(t))
	}
}
