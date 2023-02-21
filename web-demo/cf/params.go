package cf

// 当前配置文件
type appParams struct {
	ProxyUrl string `v:""` // 代理服务器地址
}

var DParams *appParams

func initAppParams() {
	DParams = &AppCnf.CurrAppParams
}
