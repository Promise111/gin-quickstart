package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a" binding:"required"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructB
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetB(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b StructB
		if bindErr := c.ShouldBind(&b); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": bindErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"a": b.NestedStruct,
			"b": b.FieldB,
		})
	}
}

func GetC(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rc StructC
		if bindErr := c.ShouldBind(&rc); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"messgae": bindErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"a": rc.NestedStructPointer,
			"b": rc.FieldC,
		})
	}
}

func GetD(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var d StructD
		if bindErr := c.ShouldBind(&d); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"messgae": bindErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"a": d.NestedAnonyStruct,
			"b": d.FieldD,
		})
	}
}
