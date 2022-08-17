package hr

import (
	"github.com/qinchende/gofast/store/orm"
)

type SysDepartment struct {
	orm.CommonFields
	ParentID int16  `pms:"parent_id" v:"required,range=[0:]"`
	Name     string `pms:"name" v:"required"`
}

//
//func (*SysDepartment) GfAttrs(super orm.OrmStruct) *orm.ModelAttrs {
//	return &orm.ModelAttrs{
//		TableName: "sys_department",
//	}
//}
