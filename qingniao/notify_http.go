package qingniao

import (
	`encoding/json`
	`fmt`

	`github.com/storezhang/gox`
)

// Http Http通知
type Http struct {
	// Url 通信地址
	Url string `json:"url" validate:"required"`
	// Headers 请求头
	Headers map[string]string `json:"headers"`
}

// NewJWTNotify 创建新的JWT类型的通知
func NewJWTNotify(url string, scheme string, token string) NotifyReq {
	return NewHttpNotify(url, map[string]string{
		gox.HeaderAuthorization: fmt.Sprintf("%s %s", scheme, token),
	})
}

// NewHttpNotify 修建新的Http类型的通知
func NewHttpNotify(url string, headers map[string]string) NotifyReq {
	return NotifyReq{
		Type: NotifyTypeHttp,
		Data: Http{
			Url:     url,
			Headers: headers,
		},
	}
}

func (nh Http) String() string {
	jsonBytes, _ := json.MarshalIndent(nh, "", "    ")

	return string(jsonBytes)
}
