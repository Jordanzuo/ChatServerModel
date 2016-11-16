package serverResponseObject

// 服务端响应结果的状态对象，成功是0，非成功以负数来表示
type ResultStatus int

// 返回响应状态枚举值对应的描述信息字符串
func (rs ResultStatus) String() string {
	return status[rs*-1]
}

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
const (
	// 成功
	Con_Success ResultStatus = -1 * iota

	// 数据错误
	Con_DataError

	// API数据错误
	Con_APIDataError

	// 客户端数据错误
	Con_ClientDataError

	// 命令类型未定义
	Con_CommandTypeNotDefined

	// 签名错误
	Con_SignError

	// 尚未登陆
	Con_NoLogin

	// 不在公会中
	Con_NotInUnion

	// 未找到目标
	Con_NotFoundTarget

	// 不能给自己发消息
	Con_CantSendMessageToSelf

	// 玩家不存在
	Con_PlayerNotExist

	// 玩家被封号
	Con_PlayerIsForbidden

	// 玩家被禁言
	Con_PlayerIsInSilent

	// 只支持POST
	Con_OnlySupportPOST

	// API未定义
	Con_APINotDefined

	// 在另一台设备上登录
	Con_LoginOnAnotherDevice

	// 名称错误
	Con_NameError

	// 公会Id错误
	Con_UnionIdError

	// 含有屏蔽词语
	Con_ContainForbiddenWord

	//参数为空
	Con_ParamIsEmpty

	//参数不匹配
	Con_ParamNotMatch

	// 服务器组不存在
	Con_ServerGroupNotExist

	// 不能发送跨服消息
	Con_CantSendCrossServerMessage
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var status = []string{
	"Success",
	"DataError",
	"APIDataError",
	"ClientDataError",
	"CommandTypeNotDefined",
	"SignError",
	"NoLogin",
	"NotInUnion",
	"NotFoundTarget",
	"CantSendMessageToSelf",
	"PlayerNotExist",
	"PlayerIsForbidden",
	"PlayerIsInSilent",
	"OnlySupportPOST",
	"APINotDefined",
	"LoginOnAnotherDevice",
	"NameError",
	"UnionIdError",
	"ContainForbiddenWord",
	"ParamIsEmpty",
	"ParamNotMatch",
	"ServerGroupNotExist",
	"CantSendCrossServerMessage",
}
