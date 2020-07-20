package core

const CodeSuccess int = 0
const CodeFailure int = 1
const CodeParamMissing int = 2

const CodeLoginFirst = 1001

// ApiSuccess(msg string)
// ApiSuccess(code int, msg string)
// ApiSuccess(code int, msg string, data interface)
func ApiSuccess(args ...interface{}) ApiError {
	if len(args) == 1 {
		return ApiError{CodeSuccess, args[0].(string), nil}
	}
	if len(args) == 2 {
		return ApiError{args[0].(int), args[1].(string), nil}
	}
	if len(args) >= 3 {
		return ApiError{args[0].(int), args[1].(string), args[2]}
	}
	return ApiError{CodeSuccess, "success", nil}
}

// ApiFailure(msg string)
// ApiFailure(code int, msg string)
// ApiFailure(code int, msg string, data interface)
func ApiFailure(args ...interface{}) ApiError {
	if len(args) == 1 {
		return ApiError{CodeFailure, args[0].(string), nil}
	}
	if len(args) == 2 {
		return ApiError{args[0].(int), args[1].(string), nil}
	}
	if len(args) >= 3 {
		return ApiError{args[0].(int), args[1].(string), args[2]}
	}
	return ApiError{CodeFailure, "failure", nil}
}
