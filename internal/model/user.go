package model

import "time"

type User struct {
	ID         int
	Uuid       string
	passwdHash string
	avatar     string // 头像
	Name       string
	Sex        string
	Email      string
	Age        int
	Level      int
	Phone      string
	WeChat     string
	QQ         int
	IsVip      bool
	CreateAt   time.Time
	UpdateAt   *time.Time
	DeleteAt   int
}
