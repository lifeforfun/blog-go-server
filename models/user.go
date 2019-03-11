package models

import "fmt"

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Pwd      string `json:"-"`
	UserType int    `json:"userType"`
	Email    string `json:"email"`
}

const (
	UserTypeAdmin  = 1
	UserTypeNormal = 2
)

func AddUser(data map[string]interface{}) (*User, bool) {
	userInfo := User{
		Username: data["username"].(string),
		Pwd:      data["pwd"].(string),
		UserType: data["user_type"].(int),
		Email:    data["email"].(string),
	}
	db.Create(&userInfo)
	if userInfo.Id == 0 {
		return &userInfo, false
	}
	return &userInfo, true
}

func CheckAuth(username, password string) (*User, bool) {
	var user User
	db.Where(User{Username: username, Pwd: password}).First(&user)
	if user.Id > 0 {
		return &user, true
	}
	return &user, false
}

func ExistUserByID(id int) bool {
	var user User
	db.Select("id").
		Where("id = ?", id).
		First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

func ExistUserByUsername(username string) bool {
	var user User
	db.Select("username").
		Where("username = ?", username).
		First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

func GetUsers(offset int, pageSize int, maps interface{}, q string) (users []User) {
	if q != "" {
		db.Where(maps).Order("id DESC").Offset(offset).Limit(pageSize).Find(&users)
	} else {
		db.Where(maps).
			Where("username LIKE ?", fmt.Sprintf("%%%s%%", q)).
			Order("id DESC").Offset(offset).Limit(pageSize).Find(&users)
	}
	return
}

func GetUserTotal(maps interface{}, q string) (count int) {
	if q != "" {
		db.Model(&User{}).
			Where(maps).
			Where("username LIKE ?", fmt.Sprintf("%%%s%%", q)).
			Count(&count)
	} else {
		db.Model(&User{}).Where(maps).Count(&count)
	}
	return
}

func DeleteUser(id int) bool {
	db.Where("id = ?", id).Delete(User{})
	return true
}
