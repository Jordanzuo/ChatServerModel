package centerResponseObject

// Socket服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	Code ResultStatus

	// 响应结果的状态值所对应的描述信息
	Message string

	// 响应结果的数据
	Data interface{}
}

func NewResponseObject() *ResponseObject {
	return &ResponseObject{
		Code:    Con_Success,
		Message: Con_Success.string(),
		Data:    nil,
	}
}

func (responseObject *ResponseObject) SetResultStatus(rs ResultStatus) *ResponseObject {
	responseObject.Code = rs
	responseObject.Message = rs.string()

	return responseObject
}

func (responseObject *ResponseObject) SetData(data interface{}) *ResponseObject {
	responseObject.Data = data

	return responseObject
}
