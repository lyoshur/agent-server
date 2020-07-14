package task

import (
	"agent/config"
	"net/http"
)

// 获取 跨域前置
func GetCrossDomain(headers []config.Header) CrossDomain {
	crossDomain := CrossDomain{}
	crossDomain.headers = headers
	return crossDomain
}

// 跨域前置
type CrossDomain struct {
	headers []config.Header
}

// 实现Task.Do接口
func (crossDomain *CrossDomain) Do(w http.ResponseWriter, r *http.Request) bool {
	// 设置请求头
	for index := range crossDomain.headers {
		header := crossDomain.headers[index]
		w.Header().Add(header.Key, header.Value)
	}
	// Options请求检查
	if r.Method == http.MethodOptions {
		success := 204
		w.WriteHeader(success)
		return false
	}
	return true
}
