package biz_error

type errorType struct {
	Code    int
	Message string
}

type BizError struct {
	errorType
	ErrorMsg string
}

var (
	mysqlError    = errorType{1000, "读写mysql错误"}
	redisError    = errorType{1001, "读写redis错误"}
	resourceError = errorType{1002, "请求资源不存在或无权限"}
	paramError    = errorType{1003, "参数错误"}
)

func NewMysqlError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{mysqlError, errMsg}
}

func NewRedisError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{redisError, errMsg}
}

func NewResourceError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{resourceError, errMsg}
}

func NewParamError(err error) *BizError {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return &BizError{paramError, errMsg}
}
