package config

import (
	"github.com/qinchende/gofast/connx/mysql"
	"github.com/qinchende/gofast/store/sqlx"
)

var MysqlZero *sqlx.MysqlORM

func initMysql() {
	MysqlZero = mysql.NewMysqlConn(&AppCnf.SqlGoZeroCnf)
}
