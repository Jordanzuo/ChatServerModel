package serverRequestObject

import (
	"github.com/Jordanzuo/ChatServerModel/src/channelType"
)

type Command struct {
	// 玩家信息对象
	// [[
	// 玩家Id
	Id string

	// 玩家名称
	Name string

	// 合作商Id
	PartnerId int

	// 服务器Id
	ServerId int

	// 公会Id
	UnionId string

	// 玩家额外信息
	ExtraMsg string

	// 签名
	Sign string

	// ]]

	// 聊天信息对象
	// [[
	// 聊天频道类型
	ChannelType channelType.ChannelType

	// 聊天消息
	Message string

	// 私聊目标玩家Id
	ToPlayerId string
	// ]]
}
