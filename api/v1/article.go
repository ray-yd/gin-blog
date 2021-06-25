package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ray-yd/gin-blog/model"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"data":    data,
	})

}

// GetArticleList 查詢文章列表
func GetArticleList(c *gin.Context) {
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
	data, code, total := model.GetArticleList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"total":   total,
		"data":    data,
	})
}

// GetArticleInfo 查詢單個文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMessage(code),
	})

}

// GetCategoryArticle 查詢分類下文章
func GetCategoryArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	//if pageSize == 0 {
	//	pageSize = -1
	//}
	//if pageNum == 0 {
	//	pageNum = -1
	//}
	data, code, total := model.GetCategoryArticle(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
		"total":   total,
		"data":    data,
	})
}

// EditArticle 編輯文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}

// DeleteArticle 刪除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"message": errmsg.GetErrorMessage(code),
	})
}
