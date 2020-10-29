package class100

const (
	// ChannelDev 开发
	ChannelDev Channel = "dev"
	// ChannelTest 测试
	ChannelTest Channel = "test"
	// ChannelProd 生产
	ChannelProd Channel = "prod"
	// ChannelLocal 本地环境
	ChannelLocal Channel = "local"
	// ChannelSimulation 模拟请求（不发真实请求到服务器）
	ChannelSimulation Channel = "simulation"
)

// Channel 通道
type Channel string
