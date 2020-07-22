package core

// ApiError 接口返回通用结构体
type ApiError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 错误码定义
const CodeSuccess = 0
const CodeFailure = 1
const CodeParamMissing = 2

const CodeLoginFirst = 1001

// 错误码对应的错误消息定义
var CodeErrMap = map[int]string{
	CodeSuccess:    "success",
	CodeFailure:    "failure",
	CodeLoginFirst: "login first",
}

// ApiSuccess(msg string)
// ApiSuccess(code int, msg string)
// ApiSuccess(code int, msg string, data interface)
func ApiSuccess(args ...interface{}) ApiError {
	var (
		code int
		msg  string
		data interface{}
	)
	if len(args) == 1 {
		msg = args[0].(string)
	} else if len(args) == 2 {
		code, msg = args[0].(int), args[1].(string)
	} else if len(args) >= 3 {
		code, msg, data = args[0].(int), args[1].(string), args[2]
	}
	if msg == "" {
		if m, ok := CodeErrMap[code]; ok {
			msg = m
		}
	}
	return ApiError{code, msg, data}
}

// ApiFailure(msg string)
// ApiFailure(code int, msg string)
// ApiFailure(code int, msg string, data interface)
func ApiFailure(args ...interface{}) ApiError {
	var (
		code = CodeFailure
		msg  string
		data interface{}
	)
	if len(args) == 1 {
		msg = args[0].(string)
	} else if len(args) == 2 {
		code, msg = args[0].(int), args[1].(string)
	} else if len(args) >= 3 {
		code, msg, data = args[0].(int), args[1].(string), args[2]
	}
	if msg == "" {
		if m, ok := CodeErrMap[code]; ok {
			msg = m
		}
	}
	return ApiError{code, msg, data}
}
