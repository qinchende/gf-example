package gm

import (
	"github.com/qinchende/gofast/store/orm"
	"time"
)

// user的关联表
type GmInfo struct {
	orm.CommonFields
	Uid      int        `dbc:"primary_field" v:"must"`
	IsOpen   int8       `v:"def=0,range=[0:1]"`
	OpenTime *time.Time `` // 可以为空的date,datetime数据库字段，这里需要用引用类型的*time.Time
}
