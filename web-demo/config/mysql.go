package config

import (
	"github.com/qinchende/gofast/connx/mysql"
)

var MysqlZero *mysql.MSqlX

func initMysql() {
	MysqlZero = mysql.NewMysqlConn(&EnvParams.SqlGoZeroCnf)
}
