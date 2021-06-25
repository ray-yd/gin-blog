package model

import (
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(200); not null" json:"title"`
	Cid      uint     `gorm:"type:int; not null" json:"cid"`
	Name     string   `gorm:"type:varchar(20); not null" json:"name"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditArticle 編輯文章
func EditArticle(id int, data *Article) int {
	var editMap = make(map[string]interface{})
	editMap["title"] = data.Title
	editMap["cid"] = data.Cid
	editMap["desc"] = data.Desc
	editMap["content"] = data.Content
	editMap["img"] = data.Img
	err := db.Model(&Article{}).Where(`id = ?`, id).Updates(editMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArticle 刪除文章
func DeleteArticle(id int) int {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticleList 查詢分類列表
func GetArticleList(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// GetCategoryArticle 查詢分類下的所有文章
func GetCategoryArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ErrorCategoryNotExist
	}
	return cateArtList, errmsg.SUCCESS
}

// GetArticleInfo 查詢單個文章訊息
func GetArticleInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ErrorArticleNotExist
	}
	return art, errmsg.SUCCESS
}
