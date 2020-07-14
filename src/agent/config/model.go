package config

import "github.com/lyoshur/agentutils"

// 配置文件信息
type Config struct {
	agentutils.AgentConfig
	CORS       Cors       `xml:"cors"`
	SignHeader SignHeader `xml:"sign-header"`
}

// 跨域配置
type Cors struct {
	// 是否开启跨域
	Open bool `xml:"open,attr"`
	// 跨域请求头配置
	Headers []Header `xml:"header"`
}

// 跨域请求头配置
type Header struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

// 请求头签名配置
type SignHeader struct {
	// 是否请求头签名
	Open bool `xml:"open,attr"`
	// 签名
	Sign string `xml:"sign"`
}
