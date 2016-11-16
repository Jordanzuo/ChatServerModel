package channelType

// 聊天频道类型定义
type ChannelType int

func (ct ChannelType) String() string {
	if message, ok := data[ct]; ok {
		return message
	}

	return ""
}

const (
	// 系统频道
	System ChannelType = iota

	// 世界频道
	World

	// 公会频道
	Union

	// 私聊频道
	Private

	// Avatar频道，用于游戏服务器通知客户端
	Avatar

	// 跨服频道，用于跨服聊天
	CrossServer
)

var data map[ChannelType]string = map[ChannelType]string{
	0: "系统频道",
	1: "世界频道",
	2: "公会频道",
	3: "私聊频道",
	4: "Avatar频道",
	5: "跨服频道",
}
