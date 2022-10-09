package cf

import (
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/store/sqlx"
)

var Zero *sqlx.OrmDB

func InitMysql() {
	Zero = gform.OpenMysql(&AppCnf.MysqlGoZeroCnf)
}
