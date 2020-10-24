package qingniao

import (
	`encoding/json`
)

// Email 邮件通知
type Email struct {
	// Subject 标题
	Subject string `json:"subject" validate:"required"`
	// CCs 抄送人员列表
	CCs []string `json:"ccs" validate:"dive,email"`
	// BCCs 密送人员列表
	BCCs []string `json:"bccs" validate:"dive,email"`
	// To 需要发送的手机号
	To string `json:"to" validate:"required,email"`
}

// NewEmailNotify 创建新的邮件通知
func NewEmailNotify(to string, maxRetry int, data interface{}) Notify {
	return NewNotify(NotifyTypeEmail, maxRetry, Email{
		To: to,
	}, data)
}

func (e Email) String() string {
	jsonBytes, _ := json.MarshalIndent(e, "", "    ")

	return string(jsonBytes)
}
