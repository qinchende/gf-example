package model

import (
	"gf-example/server/model/gm"
	"gf-example/server/model/sm"
	"github.com/qinchende/gofast/store/orm"
	"reflect"
)

func InitModelsAttrs() {
	orm.ShareTableAttrs(modelAttrsList)
}

// 批量配置Model属性，统一放在一起，方便对参数进行管理
var modelAttrsList = map[string]*orm.TableAttrs{
	"acc.SysUser":      {CacheAll: true, ExpireS: 30},
	full(&gm.GmInfo{}): {CacheAll: true, ExpireS: 3600 * 0.1},
	full(&sm.SmInfo{}): {CacheAll: true, ExpireS: 3600 * 0.1},
}

// 通过对象变量，反射获取结构体的名称
func full(obj any) string {
	return reflect.TypeOf(obj).Elem().String()
}
