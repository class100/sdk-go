module github.com/class100/sdk-go

go 1.15

require (
	github.com/class100/core v0.0.12
	github.com/go-playground/validator/v10 v10.4.0
	github.com/go-resty/resty/v2 v2.3.0
	github.com/mcuadros/go-defaults v1.2.0
	github.com/rs/xid v1.2.1
	github.com/sirupsen/logrus v1.6.0
	github.com/storezhang/gox v1.2.25
	github.com/storezhang/replace v1.0.7
	github.com/storezhang/transfer v1.0.3
)

// replace github.com/storezhang/gox => ../../storezhang/gox
// replace github.com/storezhang/replace => ../../storezhang/replace
