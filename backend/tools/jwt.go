package tools

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MyClaims Customer jwt.StandardClaims
type MyClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

var jwtSecret = []byte("secret")

func GenToken(account string) (string, error) {
	c := MyClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString(jwtSecret)
}

// ParseToken Parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware Middleware of JWT
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		//fmt.Println(authHeader)
		//如果jwt是空的
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"stas": "failed",
				"res":  "token為空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		fmt.Println(parts)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"stas": "failed",
				"res":  "token格式錯誤",
			})
			c.Abort()
			return
		}
		// parts[0] is Bearer, parts is token.
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"stas": "failed",
				"res":  "無效的token",
			})
			c.Abort()
			return
		}
		// Store Account info into Context
		row := Sqlconnect().QueryRow("SELECT permission FROM users WHERE account = ?", mc.Account)
		var p string
		row.Scan(&p)
		c.Set("account", mc.Account)
		c.Set("permission", p)
		// After that, we can get Account info from c.Get("account")
		c.Next()
	}
}
