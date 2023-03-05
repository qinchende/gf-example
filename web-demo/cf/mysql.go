package cf

import (
	"github.com/qinchende/gofast/connx/orm"
	"github.com/qinchende/gofast/store/sqlx"
)

var Zero *sqlx.OrmDB

func InitMysql() {
	Zero = orm.OpenMysql(&AppCnf.MysqlGoZeroCnf)
}
