package controllers

import (
	"kinnyShop_api/entity"
	"kinnyShop_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateShoppingCart(c *gin.Context) {
	var s entity.ShoppingCart
	account, _ := c.Get("account")
	s.Uid = account.(string)
	c.BindJSON(&s)
	if s.Quantity < 1 {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "數量不可為0",
		})
		return
	}
	if models.CreateShoppingCart(s) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "新增購物車成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "新增購物車失敗",
		})
	}

}
func ReadShoppingCart(c *gin.Context) {
	account, _ := c.Get("account")
	s := models.ReadShoppingCart(account.(string))
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  s,
	})

}
func UpdateShoppingCart(c *gin.Context) {
	var s entity.ShoppingCart
	account, _ := c.Get("account")
	c.BindJSON(&s)
	tmp := c.Param("pkno")
	if models.UpdateShoppingCart(tmp, account.(string), s) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "更新成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "更新失敗",
		})
	}
}
func DeleteShoppingCart(c *gin.Context) {
	var s entity.ShoppingCart
	account, _ := c.Get("account")
	s.Pid, _ = strconv.Atoi(c.Param("pkno"))
	if models.DeleteShoppingCart(account.(string), s) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "刪除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "刪除失敗",
		})
	}
}
