package response

import (
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code:%d,message:%v", e.Code, e.Message)
}

func NewErrorCode(errCode string) Error {
	err := strings.Split(errCode, ",")
	code, _ := strconv.Atoi(err[0])
	return NewError(-code, err[1])
}

func NewError(code int, message string) Error {
	return Error{Code: code, Message: message}
}

type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewSuccess(data interface{}) Success {
	return Success{Code: 200, Data: data}
}
