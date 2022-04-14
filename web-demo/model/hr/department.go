package hr

import (
	"gf-example/web-demo/model"
	"github.com/qinchende/gofast/store/orm"
)

type Department struct {
	model.CommonFields
	ParentID int16  `pms:"parent_id" valid:"required,min=0"`
	Name     string `pms:"name" valid:"required"`
}

func (*Department) GfAttrs() *orm.ModelAttrs {
	return &orm.ModelAttrs{
		TableName: "sys_department",
		CacheAll:  true,
	}
}

//func (*Department) TableName() string {
//	return "sys_department"
//}
