package apiLog

import (
	"time"
)

// API日志对象
type ApiLog struct {
	// Api名称
	apiName string

	// 参数内容
	content string

	// 请求时间
	crtime time.Time
}

func (log *ApiLog) GetApiName() string {
	return log.apiName
}

func (log *ApiLog) GetContent() string {
	return log.content
}

func (log *ApiLog) GetCrtime() time.Time {
	return log.crtime
}

func NewApiLog(_apiName, _content string) *ApiLog {
	return &ApiLog{
		apiName: _apiName,
		content: _content,
		crtime:  time.Now(),
	}
}
