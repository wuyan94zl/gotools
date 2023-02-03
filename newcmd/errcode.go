package newcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
)

var genErrCodeTpl = `package response

// 错误码格式为：数字code码 + 错误信息，以英文逗号隔开。
// 业务中使用 response.NewErrorCode(response.SystemError)

const (
	ErrorSystem = "100000,系统错误"
)

`

func genErrCode(c *Command) error {
	wd := filepath.Join(c.wd, "common", "response")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "code.go",
		TemplateFile: genErrCodeTpl,
		Data:         map[string]string{},
	})
}

var genResponseTpl = `package response

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Code    int         {{.code}}
	Message interface{} {{.message}}
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
	Code int         {{.code}}
	Data interface{} {{.data}}
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

func Result(resp any, err error) (int, any) {
	switch err.(type) {
	case nil:
		return 200, NewSuccess(resp)
	case Error:
		return 200, err
	default:
		return 500, err
	}
}

`

func genResponse(c *Command) error {
	wd := filepath.Join(c.wd, "common", "response")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "response.go",
		TemplateFile: genResponseTpl,
		Data: map[string]string{
			"code":    "`json:\"code\"`",
			"message": "`json:\"message\"`",
			"data":    "`json:\"data\"`",
		},
	})
}

func genCommon(c *Command) error {
	err := genErrCode(c)
	if err != nil {
		return err
	}
	return genResponse(c)
}
