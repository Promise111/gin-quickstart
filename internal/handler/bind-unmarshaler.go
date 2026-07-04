package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// type Birthday string

var requestBind struct{
	Birthday Birthday `form:"birthday"`
	Birthdays []Birthday `form:"birthdays" collection_format:"csv"`
	BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02" collection_format:"csv"`
}

func (b *Birthday) UnmarshalParam(param string) error {
	*b = Birthday(strings.Replace(param, "-", "/", -1))
	return nil
}

var _ binding.BindUnmarshaler = (*Birthday)(nil)

func UnmarshalParam(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var _ = c.BindQuery(&requestBind)
		c.JSON(http.StatusOK, gin.H{
			"status":true,
			"message": requestBind,
		} )
	}
}
