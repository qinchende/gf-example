package acc

import (
	"github.com/qinchende/gofast/store/orm"
)

type SysUser struct {
	orm.CommonFields
	Name     string `v:"len=[1:32]"`              // 用户姓名
	Mobile   string `v:"must,match=mobile"`       // 绑定手机号
	Nickname string `v:"len=[6:32]"`              // 昵称可以为空
	Email    string `v:"match=email,len=[5:128]"` // 可以为空，否则需要匹配email类型
	Tok      string
	Age      int8
	Login    bool
}

//func (u *SysUser) GfAttrs(orm.OrmStruct) *orm.TableAttrs {
//	mAttr := u.CommonFields.GfAttrs(u)
//	mAttr.TableName = "sys_user"
//	return mAttr
//}
