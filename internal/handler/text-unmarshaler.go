package handler

import (
	"encoding"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Birthday string

var requestText struct {
	Birthday     Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
	Birthdays    []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
	BirthdaysDef []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
}

func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

var _ encoding.TextUnmarshaler = (*Birthday)(nil)

func TextUmarshal(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if bindErr := c.ShouldBindQuery(&requestText); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":false,
				"message": bindErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Ok",
			"data":    requestText,
		})
	}
}
