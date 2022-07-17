package model

import (
	"time"
	"turan/example-goWeb/admin/db"
)

type User struct {
	Id int `grom:"id"`
	Email string `gorm:"email"`
	Iphone string `gorm:"iphone"`
	Password string `gorm:"password"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
 }

func (u *User) Name() string {
	return "user"
}

//表插入
func(u *User) RegisterEmail() error {
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	err := db.GetDB().Create(u).Error
	return err
}

func (u *User) EmailLogin()  error {
	err := db.GetDB().Where("email =?",u.Email).First(u).Error
	return  err
}
