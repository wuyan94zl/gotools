package apicmd

import (
	"errors"
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var (
	VarStringName   string
	VarStringDir    string
	VarStringMethod string
	VarStringParams string
)

type Command struct {
	// 命令参数
	dir    string
	name   string
	method string
	params string

	projectPkg    string // 项目包名
	wd            string // 执行目录
	dirCamelCase  string // 路径大驼峰
	nameCamelCase string // 函数大驼峰
	routeBase     string // 路由前缀
	routeUrl      string // 路由地址
	Command       string // 完整命令
	isRequest     string // 是否有request参数
}

func (c *Command) Run() error {
	err := c.validateFlags()
	if err != nil {
		return err
	}
	err = c.initParams()
	if err != nil {
		return err
	}

	err = genTypes(c)
	if err != nil {
		return err
	}

	err = genLogic(c)
	if err != nil {
		return err
	}

	err = genHandler(c)
	if err != nil {
		return err
	}

	err = genRoute(c)
	if err != nil {
		return err
	}
	return nil
}

// 参数校验
func (c *Command) validateFlags() error {
	if VarStringDir == "" {
		return errors.New("api dir is required")
	}
	if VarStringName == "" {
		return errors.New("api name is required")
	}
	if VarStringMethod == "" {
		return errors.New("api method is required")
	}
	utils.ToLowers(&VarStringDir, &VarStringParams)
	ok, err := regexp.MatchString("^([a-z_/]*)$", VarStringDir)
	if err != nil || !ok {
		return errors.New("the --dir parameter is invalid")
	}

	ok, err = regexp.MatchString("^([A-z_]+)$", VarStringName)
	if err != nil || !ok {
		return errors.New("the --name parameter is invalid")
	}

	utils.ToUppers(&VarStringMethod)
	if VarStringMethod != "POST" && VarStringMethod != "GET" && VarStringMethod != "PUT" && VarStringMethod != "DELETE" && VarStringMethod != "RESTFUL" {
		return errors.New("the --method parameter is invalid, only GET or POST or PUT or DELETE")
	}

	if VarStringParams != "" {
		ok, err = regexp.MatchString("^([a-z:/]+)$", VarStringParams)
		if err != nil || !ok {
			return errors.New("the --params parameter is invalid")
		}
	}
	return nil
}

// 数据初始化
func (c *Command) initParams() error {
	c.dir = VarStringDir
	c.name = VarStringName
	c.method = VarStringMethod
	c.params = VarStringParams
	wd, _ := os.Getwd()
	pkg, err := utils.GetPackage(wd)
	if err != nil {
		return err
	}
	c.projectPkg = pkg
	if c.params != "" {
		c.Command = fmt.Sprintf("%s --params %s", c.Command, c.params)
	}
	c.Command = fmt.Sprintf("%s --dir %s --name %s --method %s", c.Command, VarStringDir, VarStringName, VarStringMethod)
	c.wd = wd
	c.dirCamelCase = utils.GetCamelCaseName(c.dir)
	c.nameCamelCase = utils.GetCamelCaseName(c.name)
	if c.method == "GET" || c.method == "DELETE" {
		c.isRequest = ""
	} else {
		c.isRequest = "true"
	}
	c.routeBase = VarStringDir
	if c.name == "create" && c.method == "POST" {
		//c.routeUrl = fmt.Sprintf("%s", getUrl(VarStringDir))[1:]
		c.routeUrl = ""
	} else if c.name == "update" && c.method == "PUT" {
		c.params = ":id"
		//c.routeUrl = fmt.Sprintf("%s/:id", getUrl(VarStringDir))[1:]
		c.routeUrl = ":id"
	} else if c.name == "delete" && c.method == "DELETE" {
		c.params = ":id"
		//c.routeUrl = fmt.Sprintf("%s/:id", getUrl(VarStringDir))[1:]
		c.routeUrl = ":id"
	} else if c.name == "detail" && c.method == "GET" {
		c.params = ":id"
		//c.routeUrl = fmt.Sprintf("%s/:id", getUrl(VarStringDir))[1:]
		c.routeUrl = ":id"
	} else if c.name == "list" && c.method == "GET" {
		//c.routeUrl = fmt.Sprintf("%s", getUrl(VarStringDir))[1:]
		c.routeUrl = ""
	} else {
		c.routeUrl = fmt.Sprintf("%s%s", getUrl(nameToUrl(VarStringName)), getUrl(VarStringParams))[1:]
	}

	//fmt.Println("pkg", c.projectPkg, "dirname", c.dirName, "routerReg", c.routeReg, "handlerName", c.handlerName)
	return nil
}

func getUrl(str string) string {
	if str == "" {
		return ""
	}
	return "/" + str
}

func nameToUrl(str string) string {
	newStr := ""
	for i, v := range str {
		if unicode.IsUpper(v) {
			newStr += "/"
		}
		newStr += strings.ToLower(str[i : i+1])
	}
	return newStr
}
