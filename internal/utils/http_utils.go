package utils

import (
	"context"
	"fmt"
	"strconv"

	"stuff/internal/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ErrorMessage struct {
	XCode int    `json:"x-code"` // x-code
	XMsg  string `json:"x-msg"`  // x-msg
}

type GenericResponse struct {
	StatusCode int
	ErrMsg     *ErrorMessage
	Body       interface{}
}

func (resp *GenericResponse) String() string {
	str := fmt.Sprintf("StatusCode:%d, ", resp.StatusCode)
	if resp.ErrMsg != nil {
		str += fmt.Sprintf("XCode:%+v, XMsg:%s, ", resp.ErrMsg.XCode, resp.ErrMsg.XMsg)
	}
	str += fmt.Sprintf("Body:%+v", resp.Body)
	return str
}

func HTTPJSONResponse(ctx context.Context, c *gin.Context, resp *GenericResponse) {
	log.Debugf("%s, resp: %+v", ShouldGetTraceID(ctx), resp)

	c.Header(errors.XCode, strconv.Itoa(resp.ErrMsg.XCode))
	c.Header(errors.XMsg, resp.ErrMsg.XMsg)
	c.Header(errors.XTraceID, ShouldGetTraceID(ctx))

	if resp.Body != nil {
		c.JSON(resp.StatusCode, resp.Body)
	} else {
		c.Status(resp.StatusCode)
	}
}

func SetSuccessfulResponse(resp *GenericResponse, code int, body interface{}) {
	setResponse(resp, code, "", body)
}

func SetErrorResponse(resp *GenericResponse, errCode int, errMsg string) {
	setResponse(resp, errCode, errMsg, nil)
}

func SetErrorResponseWithError(resp *GenericResponse, err error) {
	if customErr, ok := err.(errors.CustomError); ok {
		setResponse(resp, customErr.ErrorCode, customErr.ErrorMsg, nil)
		return
	}
	setResponse(resp, errors.CodeInternalServerError, errors.GetErrorMessage(errors.CodeInternalServerError), nil)
}

func DefaultResponse() GenericResponse {
	return GenericResponse{
		ErrMsg: &ErrorMessage{
			XCode: errors.CodeOK,
			XMsg:  errors.GetErrorMessage(errors.CodeOK),
		},
	}
}

func setResponse(resp *GenericResponse, errCode int, errMsg string, body interface{}) {
	resp.StatusCode = errors.GetStatusCode(errCode)
	if body != nil {
		resp.Body = body
	} else {
		resp.ErrMsg.XCode = errCode
		if len(errMsg) == 0 {
			errMsg = errors.GetErrorMessage(errCode)
		}
		resp.ErrMsg.XMsg = errMsg
	}
}
