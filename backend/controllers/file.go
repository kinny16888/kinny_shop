package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ReadFile(c *gin.Context) {
	// fileName := c.Param("pkno")
	fileName := c.Query("fileName")
	folder := c.Query("folder")
	fmt.Println("img/" + folder + "/" + fileName)
	c.File("img/" + folder + "/" + fileName)
}
