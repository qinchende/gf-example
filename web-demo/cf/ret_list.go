package cf

import (
	"github.com/qinchende/gofast/cst"
)

const (
	SucCommon    = 0
	FaiCommon    = 0
	FaiNotFound  = 101
	FaiNotSave   = 102
	FaiSaveError = 103
	FaiBindError = 104
)

var (
	FaiUserAdd    = &cst.Ret{Code: 10001, Msg: "", Desc: "这就是内部描述说明"}
	FaiUserUpdate = &cst.Ret{Code: 10002, Msg: "用户信息保存失败", Desc: "这就是内部描述"}
)
