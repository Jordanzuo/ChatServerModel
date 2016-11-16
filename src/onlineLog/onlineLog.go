package onlineLog

import (
	"time"
)

type OnlineLog struct {
	// 在线时间
	onlineTime time.Time

	// 序号
	sid int

	// 聊天服务器地址
	serverAddress string

	// 客户端数量
	clientCount int

	// 玩家数量
	playerCount int

	// 所有服务器的总数量
	totalCount int
}

func (log *OnlineLog) GetOnlineTime() time.Time {
	return log.onlineTime
}

func (log *OnlineLog) GetSid() int {
	return log.sid
}

func (log *OnlineLog) GetServerAddress() string {
	return log.serverAddress
}

func (log *OnlineLog) GetClientCount() int {
	return log.clientCount
}

func (log *OnlineLog) GetPlayerCount() int {
	return log.playerCount
}

func (log *OnlineLog) GetTotalCount() int {
	return log.totalCount
}

func (log *OnlineLog) SetTotalCount(value int) {
	log.totalCount = value
}

func NewOnlineLog(_sid int, _serverAddress string, _clientCount, _playerCount int) *OnlineLog {
	return &OnlineLog{
		onlineTime:    time.Now(),
		sid:           _sid,
		serverAddress: _serverAddress,
		clientCount:   _clientCount,
		playerCount:   _playerCount,
	}
}
