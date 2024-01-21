package model

type T any

// Result 响应结果
type Result struct {
	Code    int    `json:"code"`           // 业务状态码
	Message string `json:"message"`        // 业务信息
	Data    T      `json:"data,omitempty"` // 响应数据
}

const (
	BusinessOk = iota + 200 // BusinessOk 业务正常
)
const (
	BusinessError = iota + 500 // BusinessError 业务出错
)

// 业务状态码
const (
	WrongParam = iota + 5000 // WrongParam 参数错误
	IllegalID                // IllegalID 非法ID
	BusinessQueryError
	BusinessInsertError
	BusinessDeleteError
)

// RespText 响应文本
func RespText(code int, msg string) *Result {
	res := &Result{}
	res.Code = code
	res.Message = msg
	return res
}

// RespData 响应数据
func RespData(code int, msg string, data T) *Result {
	res := RespText(code, msg)
	res.Data = data
	return res
}

// SuccessText 响应成功文本
func SuccessText() *Result {
	res := RespText(200, "响应成功")
	return res
}

// SuccessData 响应成功数据
func SuccessData(data T) *Result {
	res := SuccessText()
	res.Data = data
	return res
}
