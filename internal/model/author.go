package model

import "time"

type Author struct {
	ID         int
	Uuid       string
	Name       string
	passwdHash string
	Sex        string
	Avatar     string // 用户头像
	Email      string
	Age        int
	Level      int
	Works      []*Works // 作品
	Earnings   float64  // 收益
	Phone      string
	WeChat     string
	QQ         int
	CreatAt    time.Time
	UpdateAt   *time.Time
}
