package hr

import (
	"github.com/qinchende/gofast/store/orm"
)

type Title struct {
	orm.CommonFields
	Level int16
	Name  string `pms:"name" v:"required"`
	Desc  string `pms:"desc" v:""`
}

func (t *Title) GfAttrs(orm.OrmStruct) *orm.ModelAttrs {
	return &orm.ModelAttrs{
		TableName: "sys_title",
	}
}
