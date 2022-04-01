package hr

import (
	"gf-example/web-demo/model"
	"time"
)

type SysUser struct {
	model.CommonFields
	Account  string `pms:"account" valid:"required,len=3"`     // 不能为空，长度3字符
	Name     string `pms:"name" valid:"required"`              // 不能为空
	Nickname string `pms:"nickname"`                           // 无验证
	Age      int8   `pms:"age" valid:"required,gte=0,lte=130"` // 年龄: >=0 && <=130
	Email    string `pms:"email" valid:"omitempty,email"`      // 可以为空，否则需要匹配email类型
}

//func (*SysUser) TableName() string {
//	return "`sys_user`"
//}

// ++++++++++++++++++++++++++++++++++++++++++
type SysUserDemo struct {
	ID        uint32    `dbc:"primary_field"`
	Status    int8      `valid:"min=-3"`
	CreatedAt time.Time `dbc:"created_field"`
	UpdatedAt time.Time `dbc:"updated_field"`
	Account   string    `pms:"account" valid:"required,len=3"`     // 不能为空，长度3字符
	Name      string    `pms:"name" valid:"required"`              // 不能为空
	Nickname  string    `pms:"nickname"`                           // 无验证
	Age       int8      `pms:"age" valid:"required,gte=0,lte=130"` // 年龄: >=0 && <=130
	Email     string    `pms:"email" valid:"omitempty,email"`      // 可以为空，否则需要匹配email类型
}

func (*SysUserDemo) TableName() string {
	return "`sys_user`"
}
