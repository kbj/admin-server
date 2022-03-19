package enum

const (
	StatusSuccess             int = 0     // 请求成功
	StatusForbidden           int = 40300 // 无权限
	StatusError               int = -1    // 请求有错误
	StatusNotFound            int = 40400 // 找不到资源
	StatusInternalServerError int = 50000 // 服务器报错
	StatusNeedLogin           int = 40301 // 需要重新跳转登录
)
