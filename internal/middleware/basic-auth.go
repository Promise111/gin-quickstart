package middleware

import "github.com/gin-gonic/gin"

var BasicAuth = gin.BasicAuth(gin.Accounts{
	"foo":           "bar",
	"austin":        "1234",
	"lena":          "hello2",
	"manu":          "4321",
	"promiseihunna": "5401",
})
