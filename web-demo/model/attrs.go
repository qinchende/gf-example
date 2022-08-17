package model

import (
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/store/orm"
	"reflect"
)

func init() {
	orm.ShareModelAttrs(modelAttrsList)
}

var modelAttrsList = map[string]*orm.ModelAttrs{
	"hr.SysUser":             {CacheAll: true, ExpireS: 3600 * 12}, // 每条记录缓存12个小时
	"hr.Title":               {CacheAll: true, ExpireS: 3600 * 12},
	"hr.SysDepartment":       {CacheAll: true, ExpireS: 3600 * 12, TableName: "sys_department"},
	full(hr.SysDepartment{}): {CacheAll: true, ExpireS: 3600 * 12, TableName: "sys_department"},
}

func full(obj any) string {
	return reflect.TypeOf(obj).Elem().String()
}