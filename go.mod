module github.com/class100/sdk-go

go 1.15

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.4+incompatible // indirect
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/go-playground/validator/v10 v10.4.0
	github.com/go-resty/resty/v2 v2.3.0
	github.com/jlaffaye/ftp v0.0.0-20200812143550-39e3779af0db // indirect
	github.com/mcuadros/go-defaults v1.2.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/storezhang/gox v1.2.5
	github.com/storezhang/replace v1.0.1
	github.com/storezhang/transfer v1.0.0
	github.com/tidwall/sjson v1.1.2 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
)

// replace github.com/storezhang/gox => ../../storezhang/gox
replace github.com/storezhang/replace => ../../storezhang/replace
