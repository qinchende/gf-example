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
	"hr.SysUser":              {CacheAll: true, ExpireS: 3600 * 1},
	"hr.Title":                {CacheAll: true, ExpireS: 3600 * 1},
	"hr.SysDepartment":        {CacheAll: true, ExpireS: 3600 * 1, TableName: "sys_department"},
	full(&hr.SysDepartment{}): {CacheAll: true, ExpireS: 3600 * 1, TableName: "sys_department"},
}

func full(obj any) string {
	return reflect.TypeOf(obj).Elem().String()
}
