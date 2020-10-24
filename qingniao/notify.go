package qingniao

import (
	`encoding/json`

	`github.com/class100/sdk-go`
)

const (
	// NotifyTypeHttp Http通知
	NotifyTypeHttp NotifyType = "http"
	// NotifyTypeVoice 语音通知
	NotifyTypeVoice NotifyType = "voice"
	// NotifyTypeSMS 短信通知
	NotifyTypeSms NotifyType = "sms"
	// NotifyTypeEmail 邮件通知
	NotifyTypeEmail NotifyType = "email"
)

type (
	// NotifyType 通知类型
	NotifyType string

	// Notify 回调通知
	Notify struct {
		// Type 通知类型
		Type NotifyType `default:"http" json:"type" validate:"required,oneof=http voice sms email"`
		// MaxRetry 最大重试次数
		MaxRetry int `default:"6" json:"maxRetry" validate:"omitempty,min=1,max=100"`
		// Notifier 真正的通知者
		Notifier interface{} `json:"notifier" validate:"required"`
		// Data 数据
		Data interface{} `json:"data"`
	}
)

// NewSimpleNotify 创建简单的通知
func NewSimpleNotify(notifyType NotifyType, notifier interface{}, data interface{}) *Notify {
	return NewNotify(notifyType, class100.DefaultRetryTimes, notifier, data)
}

// NewNotify 创建一个新的通知
func NewNotify(notifyType NotifyType, maxRetry int, notifier interface{}, data interface{}) *Notify {
	return &Notify{
		Type:     notifyType,
		MaxRetry: maxRetry,
		Notifier: notifier,
		Data:     data,
	}
}

func (n *Notify) UnmarshalJSON(data []byte) (err error) {
	type cloneType Notify

	rawMsg := json.RawMessage{}
	n.Notifier = &rawMsg

	if err = json.Unmarshal(data, (*cloneType)(n)); err != nil {
		return
	}
	switch n.Type {
	case NotifyTypeHttp:
		http := Http{}
		if err = json.Unmarshal(rawMsg, &http); err != nil {
			return
		}
		n.Notifier = http
	case NotifyTypeVoice:
		voice := Voice{}
		if err = json.Unmarshal(rawMsg, &voice); err != nil {
			return
		}
		n.Notifier = voice
	case NotifyTypeSms:
		sms := Sms{}
		if err = json.Unmarshal(rawMsg, &sms); err != nil {
			return
		}
		n.Notifier = sms
	case NotifyTypeEmail:
		email := Email{}
		if err = json.Unmarshal(rawMsg, &email); err != nil {
			return
		}
		n.Notifier = email
	default:
		err = ErrorNotSupportNotify
	}

	return
}

func (n Notify) String() string {
	jsonBytes, _ := json.MarshalIndent(n, "", "    ")

	return string(jsonBytes)
}
