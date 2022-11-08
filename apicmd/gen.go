package apicmd

import (
	"errors"
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	VarStringName   string
	VarStringDir    string
	VarStringMethod string
	VarStringUrl    string
	VarStringParams string
)

type Command struct {
	// 命令参数
	dir    string
	name   string
	method string
	params string
	url    string

	projectPkg  string // 项目包名
	wd          string // 执行目录
	routeUrl    string // 路由地址
	routeReg    string // 路由handler注册
	handlerName string // 路由handler函数名
	dirName     string // 路径名消去/，字母小写
	Command     string // 完整命令
	packageName string // 包名
	PackageName string // 函数名
	middleware  string // 中间件
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

	err = genRoute(c)
	if err != nil {
		return err
	}
	c.wd = filepath.Join(c.wd, "app")

	err = genTypes(c)
	if err != nil {
		return err
	}

	err = genHandler(c)
	if err != nil {
		return err
	}

	err = genLogic(c)
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
	utils.ToLowers(&VarStringDir, &VarStringName, &VarStringUrl, &VarStringParams)
	ok, err := regexp.MatchString("^([a-z/]+)$", VarStringDir)
	if err != nil || !ok {
		return errors.New("the --dir parameter is invalid")
	}
	ok, err = regexp.MatchString("^([a-z]+)$", VarStringName)
	if err != nil || !ok {
		return errors.New("the --name parameter is invalid")
	}
	if VarStringMethod != "POST" && VarStringMethod != "GET" && VarStringMethod != "PUT" && VarStringMethod != "DELETE" && VarStringMethod != "RESTFUL" {
		return errors.New("the --method parameter is invalid, only GET or POST or PUT or DELETE")
	}
	if VarStringUrl != "" {
		ok, err = regexp.MatchString("^([a-z/]+)$", VarStringUrl)
		if err != nil || !ok {
			return errors.New("the --url parameter is invalid")
		}
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
	c.url = VarStringUrl

	pkg, err := utils.GetPackage()
	if err != nil {
		return err
	}
	c.projectPkg = pkg
	if c.url != "" {
		c.Command = fmt.Sprintf("%s --url %s", c.Command, c.url)
	}
	if c.params != "" {
		c.Command = fmt.Sprintf("%s --params %s", c.Command, c.params)
	}
	c.Command = fmt.Sprintf("%s --dir %s --name %s --method %s", c.Command, VarStringDir, VarStringName, VarStringMethod)
	wd, _ := os.Getwd()
	c.wd = wd
	c.dirName = strings.ToLower(getName(c.dir))
	c.routeReg = getName(c.dir) + getName(formatUrl(c.url))
	c.handlerName = getName(c.dir) + getName(formatUrl(c.url)) + getName(c.name)
	c.routeUrl = fmt.Sprintf("%s%s%s%s", getUrl(VarStringDir), getUrl(VarStringUrl), getUrl(VarStringName), getUrl(VarStringParams))[1:]

	//fmt.Println("pkg", c.projectPkg, "dirname", c.dirName, "routerReg", c.routeReg, "handlerName", c.handlerName)
	return nil
}

func getName(str string) string {
	s := strings.Split(str, "/")
	rlt := ""
	for _, v := range s {
		rlt += utils.UpperOne(v)
	}
	return rlt
}

func getUrl(str string) string {
	if str == "" {
		return ""
	}
	return "/" + str
}

func formatUrl(str string) string {
	i := strings.Index(str, ":")
	if i != -1 {
		return str[0:i]
	}
	return str
}
