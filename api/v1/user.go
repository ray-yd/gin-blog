package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ray-yd/gin-blog/model"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"net/http"
	"strconv"
)

var code int

// AddUser 新增帳號
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})
}

// 查詢帳號

// GetUsers 查詢帳號列表
func GetUsers(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		fmt.Println(err)
	}
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		fmt.Println(err)
	}
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUserList(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"data":    data,
	})
}

// EditUser 編輯帳號
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}

// DeleteUser 刪除帳號
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
