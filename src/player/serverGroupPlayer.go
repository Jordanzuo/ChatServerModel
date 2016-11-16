package player

import (
	"sync"
)

// 服务器组玩家对象
type ServerGroupPlayer struct {
	// 服务器组Id
	serverGroupId int

	// 玩家集合
	playerMap map[string]*Player

	// 锁对象
	mutex sync.RWMutex
}

func (serverGroupPlayerObj *ServerGroupPlayer) GetPlayerList() (playerList []*Player) {
	serverGroupPlayerObj.mutex.RLock()
	defer serverGroupPlayerObj.mutex.RUnlock()

	for _, playerObj := range serverGroupPlayerObj.playerMap {
		playerList = append(playerList, playerObj)
	}

	return
}

func (serverGroupPlayerObj *ServerGroupPlayer) GetPlayerListInUnion(unionId string) (playerList []*Player) {
	serverGroupPlayerObj.mutex.RLock()
	defer serverGroupPlayerObj.mutex.RUnlock()

	for _, playerObj := range serverGroupPlayerObj.playerMap {
		if playerObj.UnionId == unionId {
			playerList = append(playerList, playerObj)
		}
	}

	return
}

func (serverGroupPlayerObj *ServerGroupPlayer) AddPlayer(playerObj *Player) {
	serverGroupPlayerObj.mutex.Lock()
	defer serverGroupPlayerObj.mutex.Unlock()

	serverGroupPlayerObj.playerMap[playerObj.Id] = playerObj
}

func (serverGroupPlayerObj *ServerGroupPlayer) DeletePlayer(playerObj *Player) {
	serverGroupPlayerObj.mutex.Lock()
	defer serverGroupPlayerObj.mutex.Unlock()

	delete(serverGroupPlayerObj.playerMap, playerObj.Id)
}

func NewServerGroupPlayer(_serverGroupId int) *ServerGroupPlayer {
	return &ServerGroupPlayer{
		serverGroupId: _serverGroupId,
		playerMap:     make(map[string]*Player, 1024),
	}
}
