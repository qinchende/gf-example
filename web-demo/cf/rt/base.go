package rt

import "github.com/qinchende/gofast/fst"

const (
	SucCommon    = 0
	FaiCommon    = 0
	FaiNotFound  = 101
	FaiNotSave   = 102
	FaiSaveError = 103
	FaiBindError = 104
)

const (
	FaiUserUpdateError = 10001 // 用户信息保存失败
)

var (
	//FaiUserAdd    = &fst.Ret{Code: 10001, Msg: "添加失败", Desc: "这就是内部描述"}
	FaiUserUpdate = &fst.Ret{Code: 10002, Msg: "用户信息保存失败", Desc: "这就是内部描述"}
)
