package playerDisconnectType

// 玩家断开连接类型
type PlayerDisconnectType int

const (
	// 来自于封号
	Con_FromForbid PlayerDisconnectType = 1 + iota

	// 来自于禁言
	Con_FromSilent
)
