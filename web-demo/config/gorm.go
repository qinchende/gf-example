package config

import (
	"github.com/qinchende/gofast/connx/gormc"
	"gorm.io/gorm"
)

var GormZero *gorm.DB

func initGorm() {
	GormZero = gormc.NewGormConn(&SysCnf.SqlGoZeroCnf)
}
