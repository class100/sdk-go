package qingniao

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storezhang/gox"

	class100 "github.com/class100/sdk-go"
)

// NotifyHttp Http通知
type NotifyHttp struct {
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
		Data: NotifyHttp{
			Url:     url,
			Headers: headers,
		},
	}
}

func (nh *NotifyHttp) Notify(data interface{}) (err error) {
	var (
		req *resty.Request
		rsp *resty.Response
	)

	req = class100.NewResty()
	if 0 < len(nh.Headers) {
		req.SetHeaders(nh.Headers)
	}

	if rsp, err = class100.NewResty().SetBody(data).Post(nh.Url); nil != err {
		log.WithFields(log.Fields{
			"url":     nh.Url,
			"headers": nh.Headers,
			"body":    class100.RestyStringBody(rsp),
			"error":   err,
		}).Error("通知服务器出错")

		return
	}

	if http.StatusOK != rsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(rsp.StatusCode()), "通知服务器出错", rsp.String())

		return
	}

	return
}

func (nh NotifyHttp) String() string {
	jsonBytes, _ := json.MarshalIndent(nh, "", "    ")

	return string(jsonBytes)
}
