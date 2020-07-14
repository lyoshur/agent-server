package task

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/lyoshur/gutils/cache"
	"net/http"
	"strconv"
	"time"
)

// 获取 请求头签名前置处理
func GetHeaderSign(token string) HeaderSign {
	sign := HeaderSign{}
	sign.token = token
	return sign
}

// 请求头签名前置处理
type HeaderSign struct {
	// 签名密钥
	token string
}

// 实现Task.Do接口
func (headerSign *HeaderSign) Do(w http.ResponseWriter, r *http.Request) bool {
	accessToken := r.Header.Get("AccessToken")
	timeStamp := r.Header.Get("TimeStamp")

	// 判断时间戳参数是否在规定范围内
	now := time.Now().Unix() * 1000
	ts, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		fail := 400
		w.WriteHeader(fail)
		return false
	}
	difference := now - ts
	if difference < 0 {
		difference = 0 - difference
	}
	if difference > 1*60*1000 {
		fail := 400
		w.WriteHeader(fail)
		return false
	}

	// 判断AccessToken是否使用过
	ch := cache.GetHolder()
	i := ch.Get(accessToken).Data
	if i != nil {
		fail := 400
		w.WriteHeader(fail)
		return false
	}
	ch.Set(accessToken, "ok")
	ch.SetSurviveTime(accessToken, 1*time.Minute)

	// 判断AccessToken参数是否正确
	if accessToken != getAccessToken(headerSign.token, timeStamp) {
		fail := 400
		w.WriteHeader(fail)
		return false
	}
	return true
}

// 获取AccessToken参数
func getAccessToken(token string, timeStamp string) string {
	return calculateMD5("token=" + token + "&timeStamp=" + timeStamp)
}

// 计算MD5
func calculateMD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
