package controllers

import (
	"fmt"
	"kinnyShop_api/entity"
	"kinnyShop_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var o entity.Order
	account, _ := c.Get("account")
	c.BindJSON(&o)
	if len(o.Commodity) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "購物車是空的，無法送出訂單",
		})
	}
	fmt.Println(o.Address)
	if len(o.Address) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "地址不可為空",
		})
		return
	}
	//把帳號跟現在時間當訂單編號
	now := strconv.Itoa(time.Now().Year()) +
		strconv.Itoa(int(time.Now().Month())) +
		strconv.Itoa(time.Now().Day()) +
		strconv.Itoa(time.Now().Hour()) +
		strconv.Itoa(time.Now().Minute()) +
		strconv.Itoa(int(time.Now().Second()))
	o.Pkno = account.(string) + now
	o.Uid = account.(string)
	o.SendMark = false

	if check, returnItem := models.CreateOrder(o); check {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "訂單生成成功",
		})
	} else if !check && returnItem != nil {
		c.JSON(http.StatusOK, gin.H{
			"stas": "-1",
			"res":  returnItem,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "訂單生成失敗",
		})
	}
}
func ReadOrder(c *gin.Context) {
	pkno := c.Param("pkno")
	o := models.ReadOrder(pkno)
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  o,
	})
}

func ReadAllOrder(c *gin.Context) {
	account, _ := c.Get("account")
	permission, _ := c.Get("permission")
	o := models.ReadAllOrder(account.(string), permission.(string))
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  o,
	})
}
func ReadDetailedOrder(c *gin.Context) {
	pkno := c.Param("pkno")
	if pkno == "" {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "請輸入查詢條件",
		})
		return
	}
	p := models.ReadDetailedOrder(pkno)
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  p,
	})
}
func UpdateSendMrk(c *gin.Context) {
	pkno := c.Param("pkno")
	p := models.UpdateSendMrk(pkno)
	if p {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "更新出貨註記成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "更新出貨註記失敗",
		})
	}

}
