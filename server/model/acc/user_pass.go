package acc

import (
	"github.com/qinchende/gofast/store/orm"
)

type SysUserPass struct {
	orm.CommonFields
	Uid      int    `v:"must" dbc:"primary_field"`
	Salt     string `v:"must,len=[4:8]"`
	Sha1Pass string `v:"must,math=base64,len=[40:40]"`
}

const (
	SqlAuthLogin = `select user_id from sys_user_pass where user_id=(select id from sys_user where account=? limit 1) and hash_pass=sha1(concat(md5(id+110),?)) limit 1;`
)
