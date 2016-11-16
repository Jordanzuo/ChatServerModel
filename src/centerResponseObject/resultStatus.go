package centerResponseObject

// 服务端响应结果的状态对象，成功是0，非成功以负数来表示
type ResultStatus int

// 返回响应状态枚举值对应的描述信息字符串
func (rs ResultStatus) string() string {
	return status[rs*-1]
}

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
const (
	// 成功
	Con_Success ResultStatus = -1 * iota

	// 数据错误
	Con_DataError

	// 方法未定义
	Con_MethodNotDefined

	// 参数为空
	Con_ParamIsEmpty

	// 参数不匹配
	Con_ParamNotMatch

	// 参数类型错误
	Con_ParamTypeError

	// 只支持POST
	Con_OnlySupportPOST

	// API未定义
	Con_APINotDefined

	// API数据错误
	Con_APIDataError

	// API参数错误
	Con_APIParamError

	// IP无效
	Con_InvalidIP

	// 重新加载出错
	Con_ReloadError

	// 玩家不存在
	Con_PlayerNotExist

	// 服务器组不存在
	Con_ServerGroupNotExist
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var status = [...]string{
	// 成功
	"Success",

	// 数据错误
	"DataError",

	// 方法未定义
	"MethodNotDefined",

	// 参数为空
	"ParamIsEmpty",

	// 参数不匹配
	"ParamNotMatch",

	// 参数类型错误
	"ParamTypeError",

	// 只支持POST
	"OnlySupportPOST",

	// API未定义
	"APINotDefined",

	// API数据错误
	"APIDataError",

	// API参数错误
	"APIParamError",

	// IP无效
	"InvalidIP",

	// 重新加载出错
	"ReloadError",

	// 玩家不存在
	"PlayerNotExist",

	// 服务器组不存在
	"ServerGroupNotExist",
}
