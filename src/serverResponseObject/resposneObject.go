package serverResponseObject

import (
	"github.com/Jordanzuo/ChatServerModel/src/commandType"
)

// Chat服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	Code ResultStatus

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}

	// 响应结果对应的请求命令类型
	CommandType commandType.CommandType
}

func NewResponseObject(ct commandType.CommandType) *ResponseObject {
	return &ResponseObject{
		Code:        Con_Success,
		Message:     Con_Success.String(),
		Data:        nil,
		CommandType: ct,
	}
}

func (responseObject *ResponseObject) SetResultStatus(rs ResultStatus) *ResponseObject {
	responseObject.Code = rs
	responseObject.Message = rs.String()

	return responseObject
}

func (responseObject *ResponseObject) SetCommandType(ct commandType.CommandType) *ResponseObject {
	responseObject.CommandType = ct

	return responseObject
}

func (responseObject *ResponseObject) SetData(data interface{}) *ResponseObject {
	responseObject.Data = data

	return responseObject
}
