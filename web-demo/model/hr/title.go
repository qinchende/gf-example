package hr

import "gf-example/web-demo/model"

type Title struct {
	model.CommonFields
	Level int16
	Name  string `pms:"name" valid:"required"`
	Desc  string `pms:"desc" valid:"omitempty"`
}

func (*Title) TableName() string {
	return "sys_title"
}
