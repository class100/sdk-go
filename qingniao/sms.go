package qingniao

import (
	`encoding/json`
)

// Sms 短信通知
type Sms struct {
	// To 需要发送的手机号
	To string `json:"to" validate:"required"`
}

// NewSmsNotify 创建新的短信通知
func NewSmsNotify(to string, maxRetry int, data interface{}) Notify {
	return NewNotify(NotifyTypeSms, maxRetry, Sms{
		To: to,
	}, data)
}

func (s Sms) String() string {
	jsonBytes, _ := json.MarshalIndent(s, "", "    ")

	return string(jsonBytes)
}
