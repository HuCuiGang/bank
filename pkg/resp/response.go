package resp

type ResponseCode int

const (
	SUCCESS ResponseCode = 0
	ERROR   ResponseCode = -1
)

type ResponseCommon struct {
	Code    ResponseCode `json:"code"` // 状态代码
	Message string       `json:"message"`  // 错误信息
}

type AAAResponse struct {
	ResponseCommon
	Data interface{} `json:"data"`
}
