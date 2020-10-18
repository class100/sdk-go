package qingniao

import (
	`encoding/json`
)

// Voice 语音通知
type Voice struct {
	// To 需要发送的手机号
	To string `json:"to" validate:"required"`
}

// NewVoiceNotify 修建新的Voice类型的通知
func NewVoiceNotify(to string, maxRetry int, data interface{}) Notify {
	return NewNotify(NotifyTypeVoice, maxRetry, Voice{
		To: to,
	}, data)
}

func (s Voice) String() string {
	jsonBytes, _ := json.MarshalIndent(s, "", "    ")

	return string(jsonBytes)
}
