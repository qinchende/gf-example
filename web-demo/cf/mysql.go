package cf

import (
	"github.com/qinchende/gofast/connx/mysql"
	"github.com/qinchende/gofast/store/sqlx"
)

var Zero *sqlx.MysqlORM

func initMysql() {
	Zero = mysql.OpenMysql(&AppCnf.SqlGoZeroCnf)
}
