package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

const (
	CodeOK                                     = 0    // 成功ok
	CodeInternalServerError                    = 9500 // 服务内部失败
	CodeUnauthorized                           = 9501 // 401 用户未传token
	CodeForbidden                              = 9502 // 403 鉴权失败，如token无效或者过期
	CodeInvalidParams                          = 9503 // 400 参数错误
	CodeResourcesNotFount                      = 9504 // 404 资源未找到
	CodeResourcesHasExist                      = 9505 // 409 资源已存在
	CodeResourcesConflict                      = 9506 // 409 状态冲突
	CodeRecordCallInspectionCompleted          = 9507
	CodeInspectionRecordScoreDetailMatchFailed = 9508
	CodeUnknownError                           = 9509 // 未知异常
	CodeNoRight2Modify                         = 9510 // 用户没权限修改相关资源
)

var statusCode = map[int]int{
	CodeOK:                                     http.StatusOK,
	CodeInternalServerError:                    http.StatusInternalServerError,
	CodeUnauthorized:                           http.StatusUnauthorized,
	CodeForbidden:                              http.StatusForbidden,
	CodeInvalidParams:                          http.StatusBadRequest,
	CodeResourcesNotFount:                      http.StatusNotFound,
	CodeResourcesHasExist:                      http.StatusConflict,
	CodeResourcesConflict:                      http.StatusConflict,
	CodeRecordCallInspectionCompleted:          http.StatusConflict,
	CodeInspectionRecordScoreDetailMatchFailed: http.StatusConflict,
	CodeUnknownError:                           http.StatusInternalServerError,
	CodeNoRight2Modify:                         http.StatusForbidden,
}

var errorMessage = map[int]string{
	CodeOK:                                     "ok",
	CodeInternalServerError:                    "internal server error",
	CodeUnauthorized:                           "header authorization is empty",
	CodeForbidden:                              "token is forbidden",
	CodeInvalidParams:                          "invalid param",
	CodeResourcesNotFount:                      "the resource was not found",
	CodeResourcesHasExist:                      "the resource has already exists",
	CodeResourcesConflict:                      "the resource has conflict expect status value",
	CodeRecordCallInspectionCompleted:          "record call is inspection completed",
	CodeInspectionRecordScoreDetailMatchFailed: "inspection record call success to match score detail not all",
	CodeUnknownError:                           "unknown error",
	CodeNoRight2Modify:                         "have no right to modify resource",
}

const (
	XCode    = "x-code"
	XMsg     = "x-msg"
	XTraceID = "x-trace-id"
)

type CustomError struct {
	Cause     error  // 真正错误原因
	ErrorCode int    // 错误码
	ErrorMsg  string // 错误信息
}

func NewCustomError(err error, code int, msg string) error {
	return CustomError{
		Cause:     errors.Wrap(err, msg),
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (err CustomError) Error() string {
	if err.Cause != nil {
		return err.Cause.Error()
	}

	return err.ErrorMsg
}

func GetStatusCode(code int) int {
	return statusCode[code]
}

func GetErrorMessage(code int) string {
	return errorMessage[code]
}
