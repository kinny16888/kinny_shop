package controllers

import (
	"encoding/base64"
	"fmt"
	"kinnyShop_api/entity"
	"kinnyShop_api/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ReadProduct(c *gin.Context) {
	name := c.Param("pkno")
	p := models.ReadProduct(name)
	if p == nil {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "查無資料",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  p,
		})
	}

}

func ReadAllProduct(c *gin.Context) {
	p := models.ReadAllProduct()
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  p,
	})
}

func ReadProductDetailed(c *gin.Context) {
	pkno, _ := strconv.Atoi(c.Param("pkno"))
	p := models.ReadProductDetailed(pkno)
	c.JSON(http.StatusOK, gin.H{
		"stas": "succeeded",
		"res":  p,
	})

}
func CreateProduct(c *gin.Context) {
	var p entity.Product
	c.BindJSON(&p)
	fmt.Println(len(p.Name))
	if len(p.Name) < 1 || p.Price < 1 || p.Stock < 1 ||
		len(p.Introduce) < 1 || len(p.FilePath) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "所有欄位皆不可為空",
		})
		return
	}
	check, p := models.CreateProduct(p)
	if !check {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "新增失敗",
		})
		return
	}
	fmt.Println("這邊喔")
	fmt.Println(p)
	for i, v := range p.FilePath {
		if p.FilePath[i][11] == 'j' {
			p.FilePath[i] = v[23:]
		}
		var title string = "home" //因為第一張照片是預覽圖片
		if i > 0 {
			title = ""
		}
		os.Mkdir("./img/"+strconv.Itoa(p.Pkno), os.ModePerm)
		var filePath = "./img/" + strconv.Itoa(p.Pkno) + "/" + title + time.Now().Format("20220529150405") + strconv.Itoa(i) + ".jpg"
		img, _ := base64.StdEncoding.DecodeString(p.FilePath[i])
		file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer file.Close()
		file.Write(img)
		p.FilePath[i] = title + time.Now().Format("20220529150405") + strconv.Itoa(i) + ".jpg"
	}
	if models.InsertFile(p) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "新增成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "新增失敗",
		})
	}
}
func DeleteProduct(c *gin.Context) {
	id := c.Param("pkno")
	pkno, _ := strconv.Atoi(id)
	if models.DeleteProduct(pkno) {
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
func UpdateProduct(c *gin.Context) {
	var p entity.Product
	id := c.Param("pkno")
	c.BindJSON(&p)
	pkno, _ := strconv.Atoi(id)
	if models.UpdateProduct(pkno, p) {
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
