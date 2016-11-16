package serverRequestObject

import (
	"github.com/Jordanzuo/ChatServerModel/src/commandType"
)

// 客户端请求对象
type RequestObject struct {
	// 命令类型
	CommandType commandType.CommandType

	// 命令内容
	Command *Command
}
