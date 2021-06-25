package model

import (
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

// 文章分類接口

// CheckCategory 查詢分類是否存在
func CheckCategory(name string) int {
	var data Category
	db.Select(`id`).Where(`name = ?`, name).First(&data)
	if data.ID > 0 {
		return errmsg.ErrorCategoryUsed
	}
	return errmsg.SUCCESS
}

// CreateCategory 新增分類
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCategory 編輯分類
func EditCategory(id int, data *Category) int {
	var editMap = make(map[string]interface{})
	editMap["name"] = data.Name
	err := db.Model(&Category{}).Where(`id = ?`, id).Updates(editMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCategory 刪除分類
func DeleteCategory(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategoryList 查詢分類列表
func GetCategoryList(pageSize int, pageNum int) []Category {
	var categoryList []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categoryList).Error
	if err != nil {
		return nil
	}
	return categoryList
}
