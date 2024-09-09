package cf

import (
	"github.com/qinchende/gofast/connx/orm"
	"github.com/qinchende/gofast/store/sqlx"
)

var DDemo *sqlx.OrmDB

func InitMysql() {
	DDemo = orm.OpenMysql(&AppCnf.MysqlDemoCnf)
}
