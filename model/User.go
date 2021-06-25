package model

import (
	"encoding/base64"
	"fmt"
	"github.com/ray-yd/gin-blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:username; type:varchar(20); not null" json:"username" validate:"required,min=3,max=12" label:"帳號"`
	PassWord string `gorm:"column:password; type:varchar(20); not null" json:"password" validate:"required,min=6,max=20" label:"密碼"`
	Role     uint   `gorm:"column:role; type:int DEFAULT:2"  json:"role" validate:"required,gte=2" label:"權限代碼"`
}

// CheckUser 查詢帳號是否存在
func CheckUser(user string) int {
	var data User
	db.Select(`id`).Where(`username = ?`, user).First(&data)
	if data.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CreateUser 新增帳號
func CreateUser(data *User) int {
	data.PassWord = ScryptPassWord(data.PassWord)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUserList 查詢帳號列表
func GetUserList(pageSize int, pageNum int) ([]User, int) {
	var userList []User
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList).Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return userList, int(total)
}

// EditUser 編輯帳號
func EditUser(id int, data *User) int {
	var editMap = make(map[string]interface{})
	editMap["user_name"] = data.UserName
	editMap["role"] = data.Role
	err := db.Model(&User{}).Where(`id = ?`, id).Updates(editMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 刪除帳號
func DeleteUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPassWord 密碼加密
func ScryptPassWord(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 42, 61, 34, 64, 12, 94, 18}

	HashPassWord, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatalln(err)
	}
	fPassWord := base64.StdEncoding.EncodeToString(HashPassWord)
	return fPassWord
}

// BeforeSave 內建將加密後字串寫入資料庫方法(Hook)
//func (u *User) BeforeSave() (err error) {
//	u.PassWord = ScryptPassWord(u.PassWord)
//	return err
//}

// CheckLogin 登入驗證
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	fmt.Println(user.ID, user.Role, user.PassWord)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if ScryptPassWord(password) != user.PassWord {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrorUserNoRight
	}
	return errmsg.SUCCESS
}
