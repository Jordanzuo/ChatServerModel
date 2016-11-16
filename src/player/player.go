package player

import (
	"time"

	"github.com/Jordanzuo/goutil/debugUtil"
)

// 定义玩家对象
type Player struct {
	// 玩家Id
	Id string

	// 玩家名称
	Name string

	// 合作商Id
	PartnerId int

	// 服务器Id
	ServerId int

	// 玩家公会Id
	UnionId string

	// 额外透传信息
	ExtraMsg string

	// 服务器名称
	ServerName string

	// 服务器组Id
	ServerGroupId int

	//注册时间
	RegisterTime time.Time `json:"-"`

	//登录时间
	LoginTime time.Time `json:"-"`

	//是否封号
	IsForbidden bool `json:"-"`

	//禁言结束时间
	SilentEndTime time.Time `json:"-"`

	// 客户端Id(由于在数据传输过程当中，需要用到ClientId来进行判断，所以需要在JSON序列化时保留)
	ClientId int32
}

// 设置服务器信息
// serverGroupId：服务器组Id
// serverName：服务器名称
func (playerObj *Player) SetServerInfo(serverGroupId int, serverName string) {
	playerObj.ServerGroupId = serverGroupId
	playerObj.ServerName = serverName
}

// 判断玩家是否处于禁言状态
// 返回值：
// 是否处于禁言状态
// 禁言剩余分钟数
func (playerObj *Player) IsInSilent() (bool, int) {
	debugUtil.Printf("SilentEndTime:%d,Now:%d\n", playerObj.SilentEndTime.Unix(), time.Now().Unix())
	leftSeconds := playerObj.SilentEndTime.Unix() - time.Now().Unix()
	debugUtil.Printf("leftSeconds:%d\n", leftSeconds)
	if leftSeconds <= 0 {
		return false, 0
	} else {
		if leftSeconds%60 == 0 {
			return true, int(leftSeconds / 60)
		} else {
			return true, int(leftSeconds/60) + 1
		}
	}
}

// 初始化一个玩家对象
func InitPlayer(id, name string, partnerId, serverId int, unionId string, extraMsg string) *Player {
	return &Player{
		Id:            id,
		Name:          name,
		PartnerId:     partnerId,
		ServerId:      serverId,
		UnionId:       unionId,
		ExtraMsg:      extraMsg,
		RegisterTime:  time.Now(),
		LoginTime:     time.Now(),
		IsForbidden:   false,
		SilentEndTime: time.Now(),
	}
}

// 使用现有数据构造一个新的玩家对象
func NewPlayer(id, name string, partnerId, serverId int, unionId, extraMsg string, registerTime, loginTime time.Time, isForbidden bool, silentEndTime time.Time) *Player {
	return &Player{
		Id:            id,
		Name:          name,
		PartnerId:     partnerId,
		ServerId:      serverId,
		UnionId:       unionId,
		ExtraMsg:      extraMsg,
		RegisterTime:  registerTime,
		LoginTime:     loginTime,
		IsForbidden:   isForbidden,
		SilentEndTime: silentEndTime,
	}
}
