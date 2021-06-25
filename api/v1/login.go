package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ray-yd/gin-blog/middleware"
	"github.com/ray-yd/gin-blog/model"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	fmt.Println(c.Request.Body)
	_ = c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.UserName, data.PassWord)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.UserName)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"token":   token,
	})
}
