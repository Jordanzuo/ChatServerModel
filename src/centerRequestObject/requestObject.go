package centerRequestObject

// 请求对象
type RequestObject struct {
	// 请求的唯一标识，是需要通过截取请求数据前4位得到并进行手动赋值的
	Id int32

	// 以下属性是由客户端直接传入的，可以直接反序列化直接得到的
	// 请求的方法名称
	MethodName string

	// 请求的参数数组
	Parameters []interface{}
}

func NewRequestObject(methodName string, parameters []interface{}) *RequestObject {
	return &RequestObject{
		MethodName: methodName,
		Parameters: parameters,
	}
}
