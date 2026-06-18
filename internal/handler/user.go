package handler

import (
	"log"
	"maps"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Metadata struct {
	Ids         map[string]string
	CourseNames map[string]string
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
