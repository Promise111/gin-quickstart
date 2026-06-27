package handler

import (
	"log"
	"maps"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name     string    `form:"name" json:"name"`
	Address  string    `form:"address" json:"address"`
	Birthday time.Time `form:"birthday" json:"birthday" binding:"required" time_format:"2006-01-02" time_utc:"1" time_location:"Africa/Lagos"`
}

type Person2 struct {
	Name      string    `form:"name,default=Benedict"`
	Age       int       `form:"age,default=10"`
	Friends   []string  `form:"friends,default=Mishael;Daniel"`
	Addresses [2]string `form:"addresses,default=Lagos Owerri" collection_format:"ssv"`
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

type Metadata struct {
	Ids         map[string]string
	CourseNames map[string]string
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,bookabledate,gtfield=CheckIn" time_format:"2006-01-02"`
}

var BookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func GetByName(c *gin.Context) {
	name := c.Param("name")
	gender := c.Query("gender")
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User data fetched successfully",
		"name":    name,
		"gender":  strings.ToUpper(gender),
	})
}

func CreateUser(c *gin.Context) {
	age := c.DefaultPostForm("age", "18")
	gender := c.DefaultQuery("gender", "M")
	name := c.PostForm("name")

	idsQuery := c.QueryMap("ids")
	idMap := make(map[string]string)
	maps.Copy(idMap, idsQuery)

	courseNameMap := make(map[string]string)
	courseNames := c.PostFormMap("names")
	maps.Copy(courseNameMap, courseNames)
	meta := Metadata{
		Ids:         idMap,
		CourseNames: courseNameMap,
	}
	log.Println(meta)

	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "User created successfully",
		"age":     age,
		"gender":  strings.ToUpper(gender),
		"meta":    meta,
		"name":    name,
	})
}

func GetBookableDate(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var booking Booking
		if bindingErr := c.ShouldBindWith(&booking, binding.Query); bindingErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  true,
				"message": bindingErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Booking dates are valid!",
		})
	}
}

func GetStartPage(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var person Person
		if bindErr := c.ShouldBind(&person); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": bindErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   true,
			"message":  "Record fetched",
			"name":     person.Name,
			"address":  person.Address,
			"birthday": person.Birthday,
		})
	}
}

func GetPerson(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var person Person2
		if bindErr := c.ShouldBind(&person); bindErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": bindErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, person)
	}
}
