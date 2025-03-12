package model

type Works struct {
	Name    string
	Tags    *[]Tags
	Author  string
	Link    string
	Reading int    // 阅读量
	Stars   int    // 收藏量
	Like    int    // 点赞量
	Text    string // 简介
	Url     string
}
