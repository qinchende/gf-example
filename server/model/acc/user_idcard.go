package acc

import (
	"github.com/qinchende/gofast/store/orm"
)

type SysUserIDCard struct {
	orm.CommonFields
	Uid        int    `v:"must" dbc:"primary_field"`            // 一个人只能有一个身份证号码
	CardNo     string `v:"must,match=id_card"`                  // 身份证号码
	Name       string `v:"must"`                                // 身份证姓名
	Clan       string `v:"must,len=[1:32]"`                     // 名族
	Birthday   string `v:"must,match=date,time_fmt=2006-01-02"` // 出生日期
	Gender     string `v:"must,def=n,enum=m|f|n"`               // 性别
	ValidStart string `v:"must,len=[8:10]"`                     // 有效期开始
	ValidEnd   string `v:"must,len=[2:10]"`                     // 有效期结束
	IsReal     bool   ``                                        // 是否是真实有效证件
}
