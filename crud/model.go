package crud

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
)

var modelTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/response"
	"github.com/wuyan94zl/go-api/pkg/utils"
	"github.com/wuyan94zl/mysql"
	"strconv"
	"time"
)

type {{.StructName}} struct {
	{{.StructFields}}
}

func (st *{{.StructName}}) Lists(c *gin.Context) *mysql.PageList {
	data := GetService().PaginateData(c)
	mysql.GetInstance(){{.AuthWhere}}.Order("id desc").Paginate(data)
	return data
}

func (st *{{.StructName}}) Create(c *gin.Context) {
	{{.ValidateData}}
	err := mysql.GetInstance().Create(st)
	if err != nil {
		response.Error(500, err.Error())
	}
}

func (st *{{.StructName}}) Update(c *gin.Context) {
	{{.ValidateData}}
	err := mysql.GetInstance().Save(st)
	if err != nil {
		response.Error(500, err.Error())
	}
}

func (st *{{.StructName}}) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := mysql.GetInstance(){{.AuthWhere}}.First(st, id)
	if err != nil {
		response.Error(500, err.Error())
	}
}

func (st *{{.StructName}}) Delete(c *gin.Context) {
	st.Info(c)
	err := mysql.GetInstance().Delete(st)
	if err != nil {
		response.Error(500, err.Error())
	}
}

`

func setModel(structData *jsonStruct) error {
	StructFields, ValidateData, AuthWhere := structData.getStructFields()
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          getDir(structData.PackageName),
		Filename:     "model.go",
		TemplateFile: modelTpl,
		Data: map[string]string{
			"package":      structData.PackageName,
			"StructName":   structData.StructName,
			"StructFields": StructFields,
			"ValidateData": ValidateData,
			"AuthWhere":    AuthWhere,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
