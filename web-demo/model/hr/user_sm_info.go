package hr

import (
	"github.com/qinchende/gofast/store/orm"
	"time"
)

// user的关联表
type SysUserSmInfo struct {
	orm.CommonFields
	UserID   int        `dbc:"primary_field" v:"required"`
	IsOpen   int8       `v:"def=0,enum=0|1"`
	OpenTime *time.Time ``
}
