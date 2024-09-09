package sm

import (
	"github.com/qinchende/gofast/store/orm"
	"time"
)

// user的关联表
type SmInfo struct {
	orm.CommonFields
	Uid      int        `dbc:"primary_field" v:"must"`
	IsOpen   int8       `v:"def=0,enum=0|1"`
	OpenTime *time.Time ``
}
