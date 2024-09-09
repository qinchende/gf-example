package acc

import (
	"github.com/qinchende/gofast/store/orm"
)

type SysUserPostAddress struct {
	orm.CommonFields
	Uid        int    `v:"must"` // 这不是主键，一个人可能有很多收货地址，其中一个为默认地址
	RecName    string `v:"must,len=[1:32]"`
	RecAddress string `v:"must,len=[1:256]"`
	RecPhone   string `v:"must,match=mobile,len=[11:15]"` // +8613088888888
	RecPhone2  string `v:"match=mobile"`
	IsDef      bool   ``
}
