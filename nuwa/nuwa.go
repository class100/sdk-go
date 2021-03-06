package nuwa

import (
	`encoding/json`
	`fmt`
	`net/http`

	`github.com/class100/core`
	`github.com/go-resty/resty/v2`
	`github.com/mcuadros/go-defaults`
	`github.com/rs/xid`
	log `github.com/sirupsen/logrus`
	`github.com/storezhang/gox`
	`github.com/storezhang/validatorx`

	`github.com/class100/sdk-go`
)

const (
	// 打包状态
	// PackageStatusSuccess 打包成功
	PackageStatusSuccess PackageStatus = "success"
	// PackageStatusFailed 打包失败
	PackageStatusFailed PackageStatus = "failed"
)

type (
	// Client 女娲打包系统
	Client struct {
		class100.Client

		// Endpoint 地址
		Endpoint string `default:"https://nuwa.class100.com" json:"endpoint"`
	}

	// PackageStatus 打包结果
	PackageStatus string
)

// New 创建一个新的女娲客户端
func New(endpoint string, accessKey string, secretKey string) *Client {
	return &Client{
		Endpoint: endpoint,
		Client: class100.Client{
			AccessKey: accessKey,
			SecretKey: secretKey,
		},
	}
}

func (c Client) parseUrl(path string, version class100.ApiVersion) (url string) {
	if class100.ApiVersionDefault == version {
		url = fmt.Sprintf("%s/api/%s", c.Endpoint, path)
	} else {
		url = fmt.Sprintf("%s/api/%s/%s", c.Endpoint, version, path)
	}

	return
}

func (c *Client) Package(pkg *Package, environment core.Environment, version class100.ApiVersion) (rsp Response, err error) {
	if core.EnvironmentSimulation == environment {
		rsp = Response{
			Id:  xid.New().String(),
			Key: xid.New().String(),
		}

		return
	}

	// 设置默认值
	defaults.SetDefaults(c)
	// 设置默认值
	defaults.SetDefaults(pkg)
	// 确保参数合法
	if err = validatorx.Validate(pkg); nil != err {
		return
	}

	// 发送请求
	var nuwaRsp *resty.Response

	if nuwaRsp, err = class100.NewResty().SetBody(Request{Package: pkg, Request: class100.Request{Environment: environment}}).
		SetResult(&rsp).
		Post(c.parseUrl("packages", version)); nil != err {
		log.WithFields(log.Fields{
			"nuwa":  c,
			"body":  class100.RestyStringBody(nuwaRsp),
			"error": err,
		}).Error("提交打包请求出错")

		return
	}

	if http.StatusAccepted != nuwaRsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(nuwaRsp.StatusCode()), "提交打包请求出错", nuwaRsp.String())

		return
	}

	return
}

func (c Client) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}
