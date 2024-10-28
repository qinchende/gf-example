package cf

// 当前配置文件
type appParams struct {
	MyHost     string `v:""` // 当前应用地址
	ProxyUrl   string `v:""` // 代理地址
	MmsSendUrl string `v:""` // 短信平台地址
}

var DParams *appParams

func initAppParams() {
	DParams = &Cnf.CurrAppParams
}
