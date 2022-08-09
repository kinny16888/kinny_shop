package main

import (
	"kinnyShop_api/controllers"
	"kinnyShop_api/tools"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	//設定CORS的
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	router.Use(cors.New(corsConfig))
	//測試
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"stas": "succeeded",
			"res":  "連線成功",
		})
	})
	//登入註冊
	router.POST("/signUp", controllers.SignUp)
	router.POST("/login", controllers.Login)
	//產品
	router.GET("/product", controllers.ReadAllProduct)                                    //查詢全部
	router.GET("/product/:pkno", controllers.ReadProduct)                                 //查詢單一
	router.GET("/product/detailed/:pkno", controllers.ReadProductDetailed)                //查詢詳細
	router.POST("/product", tools.JWTAuthMiddleware(), controllers.CreateProduct)         //新增
	router.DELETE("/product/:pkno", tools.JWTAuthMiddleware(), controllers.DeleteProduct) //刪除
	router.PUT("/product/:pkno", tools.JWTAuthMiddleware(), controllers.UpdateProduct)    //更新
	//訂單
	router.POST("/order", tools.JWTAuthMiddleware(), controllers.CreateOrder)                          //新增訂單
	router.GET("/order", tools.JWTAuthMiddleware(), controllers.ReadAllOrder)                          //查詢全部訂單
	router.GET("/order/:pkno", tools.JWTAuthMiddleware(), controllers.ReadOrder)                       //查詢單一訂單
	router.GET("/order/detailedOrder/:pkno", tools.JWTAuthMiddleware(), controllers.ReadDetailedOrder) //查詢訂單細項
	router.PUT("/order/:pkno", tools.JWTAuthMiddleware(), controllers.UpdateSendMrk)
	//購物車
	router.POST("/shoppingCart", tools.JWTAuthMiddleware(), controllers.CreateShoppingCart)
	router.GET("/shoppingCart", tools.JWTAuthMiddleware(), controllers.ReadShoppingCart)
	router.DELETE("/shoppingCart/:pkno", tools.JWTAuthMiddleware(), controllers.DeleteShoppingCart)
	router.PUT("/shoppingCart/:pkno", tools.JWTAuthMiddleware(), controllers.UpdateShoppingCart)
	//管理者統計
	router.GET("/statistics/ranking/:pkno", tools.JWTAuthMiddleware(), controllers.Ranking)
	router.GET("/statistics/:pkno", tools.JWTAuthMiddleware(), controllers.PieChart)
	router.GET("/statistics", tools.JWTAuthMiddleware(), controllers.BarChart)
	//檔案處理
	router.GET("/file", controllers.ReadFile)

	router.Run()
}
