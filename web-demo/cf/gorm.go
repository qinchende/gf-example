package cf

import (
	"github.com/qinchende/gofast/connx/gormc"
	"gorm.io/gorm"
)

var GormZero *gorm.DB

func initGorm() {
	GormZero = gormc.NewGormConn(&AppCnf.MysqlGoZeroCnf)
}
