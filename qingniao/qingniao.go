package qingniao

import (
	`encoding/json`
	`fmt`
	`net/http`

	`github.com/go-resty/resty/v2`
	`github.com/mcuadros/go-defaults`
	log `github.com/sirupsen/logrus`
	`github.com/storezhang/gox`

	`github.com/class100/sdk-go`
)

type (
	// Client 青鸟客户端
	Client struct {
		class100.Client

		// Endpoint 地址
		Endpoint string `default:"https://nuwa.class100.com" json:"endpoint"`
	}
)

// New 创建一个新的青鸟客户端
func New(endpoint string, accessKey string, secretKey string) *Client {
	return &Client{
		Endpoint: endpoint,
		Client: class100.Client{
			AccessKey: accessKey,
			SecretKey: secretKey,
		},
	}
}

func (c *Client) Notify(nr NotifyReq, channel class100.Channel, version class100.ApiVersion) (rsp Response, err error) {
	// 设置默认值
	defaults.SetDefaults(c)
	// 设置默认值
	defaults.SetDefaults(nr)

	// 发送请求
	var qingniaoRsp *resty.Response

	if qingniaoRsp, err = class100.NewResty().SetBody(Request{Notify: nr, Request: class100.Request{Channel: channel}}).
		SetResult(&rsp).
		Post(c.parseUrl("notifies", version)); nil != err {
		log.WithFields(log.Fields{
			"nuwa":  c,
			"body":  class100.RestyStringBody(qingniaoRsp),
			"error": err,
		}).Error("提交打包请求出错")

		return
	}

	if http.StatusAccepted != qingniaoRsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(qingniaoRsp.StatusCode()), "提交打包请求出错", qingniaoRsp.String())

		return
	}

	return
}

func (c Client) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}

func (c *Client) parseUrl(path string, version class100.ApiVersion) (url string) {
	if class100.ApiVersionDefault == version {
		url = fmt.Sprintf("%s/api/%s", c.Endpoint, path)
	} else {
		url = fmt.Sprintf("%s/api/%s/%s", c.Endpoint, version, path)
	}

	return
}
