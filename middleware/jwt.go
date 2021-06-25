package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ray-yd/gin-blog/utils"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// SetToken 產生Token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginBlog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// CheckToken 驗證Token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	}
	return nil, errmsg.ERROR
}

// JwtToken JWT中介層
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		if tokenHerder == "" {
			code = errmsg.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage,
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong
			c.Abort()
			return
		}

		key, checkCode := CheckToken(checkToken[1])
		if checkCode == errmsg.ERROR {
			code = errmsg.ErrorTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage,
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage,
			})
			c.Abort()
			return
		}

		c.Set("username", key.UserName)
		c.Next()
	}
}
