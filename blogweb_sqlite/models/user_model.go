package models

import (
	"../database"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int	// 0 正常状态， 1删除
	Createtime int64
}

//插入
func InsertUser(user User) error {
	db := database.GetDB()
	db = db.Create(&user)
	return db.Error
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	var user User
	db := database.GetDB()
	db = db.Where("username = ?",username).First(&user)
	id := 0
	if db.Error == nil {
		id = user.Id
	}
	return id
}

//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string) int {
	var user User
	db := database.GetDB()
	db = db.Where("username = ? and password = ?",username,password).First(&user)
	id := 0
	if db.Error == nil {
		id = user.Id
	}
	return id
}

func CreateTableWithUser()  {
	db := database.GetDB()
	//db.CreateTable(&User{})
	db.AutoMigrate(&User{})
}