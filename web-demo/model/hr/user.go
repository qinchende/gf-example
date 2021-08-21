package hr

import "gf-example/web-demo/model"

type User struct {
	model.CommonModel
	Account  string `json:"account" pms:"account" binding:"required"`
	Name     string `json:"name" pms:"name" binding:"required"`
	Nickname string `json:"nickname" pms:"nickname"`
	Age      int16  `json:"age" pms:"age" binding:"required"`
}

func (User) TableName() string {
	return "sys_users"
}
