package qingniao

import (
	`encoding/json`
	`fmt`

	`github.com/storezhang/gox`

	`github.com/class100/sdk-go`
)

// Http Http通知
type Http struct {
	// Url 通信地址
	Url string `json:"url" validate:"required"`
	// Headers 请求头
	Headers map[string]string `json:"headers"`
}

// NewSimpleJWTNotify 简单JWT通知
func NewSimpleJWTNotify(url string, scheme string, token string) Notify {
	return NewJWTNotify(url, scheme, token, class100.DefaultRetryTimes, nil)
}

// NewJWTNotify 创建新的JWT类型的通知
func NewJWTNotify(url string, scheme string, token string, maxRetry int, data interface{}) Notify {
	return NewHttpNotify(url, map[string]string{
		gox.HeaderAuthorization: fmt.Sprintf("%s %s", scheme, token),
	}, maxRetry, data)
}

// NewHttpNotify 修建新的Http类型的通知
func NewHttpNotify(url string, headers map[string]string, maxRetry int, data interface{}) Notify {
	return NewNotify(NotifyTypeHttp, maxRetry, Http{
		Url:     url,
		Headers: headers,
	}, data)
}

func (h Http) String() string {
	jsonBytes, _ := json.MarshalIndent(h, "", "    ")

	return string(jsonBytes)
}
