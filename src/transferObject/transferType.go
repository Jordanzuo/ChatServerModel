package transferObject

type TransferType string

const (
	// 登陆
	Login TransferType = "Login"

	// 转发信息
	Forward TransferType = "Forward"

	// 更新客户端和玩家数量
	UpdateClientAndPlayerCount TransferType = "UpdateClientAndPlayerCount"
)
