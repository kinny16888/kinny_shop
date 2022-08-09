package controllers

import (
	"kinnyShop_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ranking(c *gin.Context) {
	month := c.Param("pkno")
	s := models.Ranking(month)
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  s,
	})
}
func PieChart(c *gin.Context) {
	month := c.Param("pkno")
	s := models.PieChart(month)
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  s,
	})
}
func BarChart(c *gin.Context) {
	s := models.BarChart()
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  s,
	})
}
