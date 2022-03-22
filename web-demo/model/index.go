package model

import (
	"time"
)

// dbc: 数据库相关的配置参数
// dbf: 数据库字段的名称
// pms: 绑定数值时候的字段名称
// valid: 验证命令配置

// GoFast框架的ORM定义，所有Model必须公用的方法
type CommonFields struct {
	ID        uint      `dbc:"primary_field"`
	Status    int8      `valid:"min=-3"`
	CreatedAt time.Time `dbf:"created_at" dbc:"created_field"`
	UpdatedAt time.Time `dbf:"updated_at" dbc:"updated_field"`
}

// 保存之前修改相关字段
// TODO: 万一更新失败，这里的值已经修改，需要回滚吗？？？
func (cf *CommonFields) BeforeSave() {
	if cf.ID == 0 && cf.CreatedAt.IsZero() {
		cf.CreatedAt = time.Now()
	}
	cf.UpdatedAt = time.Now()
	cf.Status = 99
}
