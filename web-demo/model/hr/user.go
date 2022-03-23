package hr

import (
	"gf-example/web-demo/model"
)

type SysUser struct {
	model.CommonFields
	Account  string `pms:"account" valid:"required,len=3"`     // 不能为空，长度3字符
	Name     string `pms:"name" valid:"required"`              // 不能为空
	Nickname string `pms:"nickname"`                           // 无验证
	Age      int8   `pms:"age" valid:"required,gte=0,lte=130"` // 年龄: >=0 && <=130
	Email    string `pms:"email" valid:"omitempty,email"`      // 可以为空，否则需要匹配email类型
}

//func (*User) TableName() string {
//	return "`sys_user`"
//}
