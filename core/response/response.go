package response

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
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

func ValidateError(code int, err error, req interface{}) Error {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		s := reflect.TypeOf(req)
		msg := ""
		for _, fieldError := range errs {
			filed, _ := s.FieldByName(fieldError.Field())
			// 获取统一错误消息
			errText := filed.Tag.Get("msg")
			if errText == "" {
				errText = fieldError.Error()
			}
			if msg == "" {
				msg = errText
			} else {
				msg = fmt.Sprintf("%s\n%s", msg, errText)
			}
		}
		return Error{Code: code, Message: msg}
	}
	return Error{Code: code, Message: err.Error()}
}
