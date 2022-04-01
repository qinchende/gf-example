package hr

import "gf-example/web-demo/model"

type Department struct {
	model.CommonFields
	ParentID int16  `pms:"parent_id" valid:"required,min=0"`
	Name     string `pms:"name" valid:"required"`
}

func (*Department) TableName() string {
	return "sys_department"
}
