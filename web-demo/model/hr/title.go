package hr

import (
	"gf-example/web-demo/model"
	"github.com/qinchende/gofast/store/orm"
)

type Title struct {
	model.CommonFields
	Level int16
	Name  string `pms:"name" valid:"required"`
	Desc  string `pms:"desc" valid:"omitempty"`
}

func (*Title) GfAttrs() *orm.ModelAttrs {
	return &orm.ModelAttrs{
		TableName: "sys_title",
		CacheAll:  true,
	}
}

//func (*Title) TableName() string {
//	return "sys_title"
//}
