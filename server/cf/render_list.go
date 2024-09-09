package cf

import (
	"github.com/qinchende/gofast/core/cst"
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
	SucLogout          = &cst.Ret{Code: 10000, Msg: "退出成功"}
	FaiLogoutPanic     = &cst.Ret{Code: 10000, Msg: "退出异常"}
	SucLogin           = &cst.Ret{Code: 10001, Msg: "登录成功"}
	FaiLogin           = &cst.Ret{Code: 10002, Msg: "登录失败"}
	FaiLoginPanic      = &cst.Ret{Code: 10003, Msg: "登录异常"}
	FaiLoginAccPassErr = &cst.Ret{Code: 10004, Msg: "账号密码错误"}

	FaiUserAdd        = &cst.Ret{Code: 50001, Msg: "", Desc: "内部日志描述说明"}
	FaiUserUpdate     = &cst.Ret{Code: 50002, Msg: "用户信息保存失败", Desc: "这就是内部描述"}
	FaiNeedAdminPower = &cst.Ret{Code: 90001, Msg: "需要管理员权限"}
)
