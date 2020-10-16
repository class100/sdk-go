package qingniao

type (
	// Notifier
	Notifier interface {
		// Notify 发送通知
		Notify(data interface{}) (err error)
	}
)
