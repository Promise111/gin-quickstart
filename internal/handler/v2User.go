package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func BindMultipleStruct(engine *gin.Engine) gin.HandlerFunc {
	var formA struct {
		Foo string `form:"foo" json:"foo" binding:"required" xml:"foo"`
	}
	var formB struct {
		Bar string `form:"bar" json:"bar" binding:"required" xml:"bar"`
	}
	return func (c *gin.Context) {
		if errA := c.ShouldBindWith(&formA, binding.JSON); errA == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":true,
				"message": "matched formA", 
				"foo": formA.Foo,
			})
			return
		}
		if errB := c.ShouldBindWith(&formB, binding.JSON); errB == nil {
			c.JSON(http.StatusOK, gin.H{
				"status":true,
				"message": "matched formB", 
				"foo": formB.Bar,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "request body did not match any known format",
		})
	}
}
