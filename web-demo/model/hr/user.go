package hr

import (
	"gf-example/web-demo/model"
	"github.com/qinchende/gofast/store/orm"
	"time"
)

type SysUser struct {
	model.CommonFields
	Account  string `pms:"account" v:"required,len=[3:3]"` // 不能为空，长度3字符
	Name     string `pms:"name" v:"required"`              // 不能为空
	Nickname string `pms:"nickname" v:"def=qinchende"`     // 无验证
	Age      int8   `pms:"age" v:"range=[0:130]"`          // 年龄: >=0 && <=130
	Email    string `pms:"email" v:"match=email"`          // 可以为空，否则需要匹配email类型
}

func (*SysUser) GfAttrs() *orm.ModelAttrs {
	return &orm.ModelAttrs{
		CacheAll: true,
		ExpireS:  3600 * 12, // 每条记录缓存12个小时
	}
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// 适合 Gorm的model定义
type SysUserDemo struct {
	ID        uint32    `dbc:"primary_field"`
	Status    int8      `v:"min=-3"`
	CreatedAt time.Time `dbc:"created_field"`
	UpdatedAt time.Time `dbc:"updated_field"`
	Account   string    `pms:"account" v:"required,len=[3:3]"` // 不能为空，长度3字符
	Name      string    `pms:"name" v:"required"`              // 不能为空
	Nickname  string    `pms:"nickname"`                       // 无验证
	Age       int8      `pms:"age" v:"required,range=[0:130]"` // 年龄: >=0 && <=130
	Email     string    `pms:"email" v:"match=email"`          // 可以为空，否则需要匹配email类型
}

func (*SysUserDemo) TableName() string {
	return "`sys_user`"
}
