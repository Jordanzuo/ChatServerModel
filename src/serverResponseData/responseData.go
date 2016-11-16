package serverResponseData

import (
	"github.com/Jordanzuo/ChatServerModel/src/channelType"
	"github.com/Jordanzuo/ChatServerModel/src/player"
)

type ResponseData struct {
	ChannelType channelType.ChannelType // 聊天频道
	Message     string                  // 聊天消息
	From        *player.Player          `json:"From,omitempty"` // 发送人
	To          *player.Player          `json:"To,omitempty"`   // 接收人
}

func NewResponseData(_channelType channelType.ChannelType, message string, from, to *player.Player) *ResponseData {
	return &ResponseData{
		ChannelType: _channelType,
		Message:     message,
		From:        from,
		To:          to,
	}
}
