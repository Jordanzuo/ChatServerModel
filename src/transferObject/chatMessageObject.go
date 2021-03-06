package transferObject

import (
	"github.com/Jordanzuo/ChatServerModel/src/channelType"
	"github.com/Jordanzuo/ChatServerModel/src/player"
	"time"
)

// 在ChatServer和ChatServerCenter之间传输的聊天消息对象
type ChatMessageObject struct {
	// 聊天频道类型
	ChannelType channelType.ChannelType

	// 消息对应的服务器组Id(0表示所有服务器组)
	ServerGroupIds string

	// 聊天消息
	Message string

	// 源玩家对象(如果是系统频道，则Player为nil)
	Player *player.Player

	// 私聊的目标玩家
	ToPlayerId string

	// 推送消息的目标玩家Id集合
	ToPlayerIds []string

	// 推送消息的目标公会Id
	ToUnionId string

	// 封号的玩家Id
	ForbidPlayerId string

	// 禁言的玩家Id
	SilentPlayerId string

	// 禁言的结束时间
	SilentEndTime time.Time

	// 发送时间
	Crtime time.Time

	// 到达服务器的时间
	ArriveServerTime time.Time

	// 服务器处理结束的时间
	HandleEndTime time.Time
}

func (chatMessageObj *ChatMessageObject) SetToPlayerId(toPlayerId string) {
	chatMessageObj.ToPlayerId = toPlayerId
}

func (chatMessageObj *ChatMessageObject) SetToPlayerIds(toPlayerIds []string) {
	chatMessageObj.ToPlayerIds = toPlayerIds
}

func (chatMessageObj *ChatMessageObject) SetToUnionId(toUnionId string) {
	chatMessageObj.ToUnionId = toUnionId
}

func (chatMessageObj *ChatMessageObject) SetForbidPlayerId(forbidPlayerId string) {
	chatMessageObj.ForbidPlayerId = forbidPlayerId
}

func (chatMessageObj *ChatMessageObject) SetSilentInfo(silentPlayerId string, silentEndTime time.Time) {
	chatMessageObj.SilentPlayerId = silentPlayerId
	chatMessageObj.SilentEndTime = silentEndTime
}

func NewChatMessageObject(_channelType channelType.ChannelType, serverGroupIds, message string, playerObj *player.Player) *ChatMessageObject {
	return &ChatMessageObject{
		ChannelType:      _channelType,
		ServerGroupIds:   serverGroupIds,
		Message:          message,
		Player:           playerObj,
		Crtime:           time.Now(),
		ArriveServerTime: time.Now(),
		HandleEndTime:    time.Now(),
	}
}
