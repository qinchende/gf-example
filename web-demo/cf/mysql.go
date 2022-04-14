package cf

import (
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/store/sqlx"
)

var Zero *sqlx.MysqlORM

func initMysql() {
	Zero = gform.OpenMysql(&AppCnf.MysqlGoZeroCnf)
}
