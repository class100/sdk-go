package qingniao

import (
	`encoding/json`
)

const (
	// NotifyTypeHttp Http通知
	NotifyTypeHttp NotifyType = "http"
)

type (
	// NotifyType 通知类型
	NotifyType string

	// NotifyReq 回调通知
	NotifyReq struct {
		// Type 通知类型
		Type NotifyType `default:"http" json:"type" validate:"required,oneof=http"`
		// MaxRetry 最大重试次数
		MaxRetry int `default:"3" json:"maxRetry" validate:"omitempty,min=1,max=100"`
		// Notifier 真正的通知者
		Notifier interface{}
		// Data 数据
		Data interface{} `json:"data"`
	}
)

func (n *NotifyReq) Notify() (err error) {
	return n.Notifier.(Notifier).Notify(n.Data)
}

func (n *NotifyReq) UnmarshalJSON(data []byte) (err error) {
	type cloneType NotifyReq

	rawMsg := json.RawMessage{}
	n.Notifier = &rawMsg

	if err = json.Unmarshal(data, (*cloneType)(n)); err != nil {
		return
	}
	switch n.Type {
	case NotifyTypeHttp:
		jwt := NotifyHttp{}
		if err = json.Unmarshal(rawMsg, &jwt); err != nil {
			return
		}
		n.Notifier = jwt
	default:
		err = ErrorNotSupportNotify
	}

	return
}

func (n NotifyReq) String() string {
	jsonBytes, _ := json.MarshalIndent(n, "", "    ")

	return string(jsonBytes)
}
