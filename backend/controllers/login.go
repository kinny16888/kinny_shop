package controllers

import (
	"kinnyShop_api/entity"
	"kinnyShop_api/models"
	"kinnyShop_api/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	//接傳來的值
	var u entity.User
	u.Account = c.Param("input")
	c.BindJSON(&u)
	if len(u.Account) < 8 ||
		len(u.Password) < 8 ||
		len(u.Account) > 14 ||
		len(u.Password) > 14 {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "帳號密碼需為8~14",
		})
		return
	}
	if len(u.Phone) < 9 ||
		len(u.Name) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "姓名不可為空，手機號碼需10碼",
		})
		return
	}
	//判斷資料庫是否有該筆帳號
	if models.FindAccount(u) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "此帳號已被使用",
		})
	} else if models.FindPhone(u) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "此手機號碼已被使用",
		})
	} else {
		//新增帳號
		tmp, err := models.AddAccount(u)
		if tmp {
			c.JSON(http.StatusOK, gin.H{
				"stas": "succeeded",
				"res":  "註冊成功",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"stas": "failed",
				"res":  err,
			})
		}

	}
}

func Login(c *gin.Context) {
	//接傳來的值
	var u entity.User
	c.BindJSON(&u)

	if len(u.Account) < 8 || len(u.Password) < 8 {
		//密碼為空提示
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "帳號密碼長度需為8~14",
		})
		return
	}
	if ynTmp, permission := models.CheckAccount(u); ynTmp {
		//拿到JWT  回傳
		jwtToken, _ := tools.GenToken(u.Account)
		c.JSON(http.StatusOK, gin.H{
			"stas":       "succeeded",
			"res":        jwtToken,
			"permission": permission,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stas": "failed",
			"res":  "帳號或密碼錯誤",
		})
	}

}
