package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ray-yd/gin-blog/model"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddCategory 新增分類
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ErrorCategoryUsed {
		code = errmsg.ErrorCategoryUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"data":    data,
	})

}

// 查詢分類下的文章

// GetCategoryList 查詢分類列表
func GetCategoryList(c *gin.Context) {
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
	data := model.GetCategoryList(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"data":    data,
	})
}

// EditCategory 編輯分類
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ErrorCategoryUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}

// DeleteCategory 刪除分類
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
