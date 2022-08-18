package handlercmd

import (
	"errors"
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"os"
	"path/filepath"
	"strings"
)

var (
	VarStringName   string
	VarStringDir    string
	VarStringMethod string
	VarStringUrl    string
)

type Command struct {
	packageName string
	wd          string
	dir         string
	handlerName string
	routeUrl    string
	method      string
	middleware  string
}

func (c *Command) Run() error {
	if VarStringName == "" {
		return errors.New("handler name is required")
	}
	if VarStringMethod == "" {
		return errors.New("handler method is required")
	}
	wd, _ := os.Getwd()
	c.wd = wd

	c.handlerName = getName(VarStringDir) + getName(VarStringUrl) + getName(VarStringName)
	c.routeUrl = fmt.Sprintf("%s%s%s", getUrl(VarStringDir), getUrl(VarStringUrl), getUrl(VarStringName))[1:]

	c.dir = VarStringDir
	c.method = VarStringMethod
	fmt.Println(c.handlerName, c.routeUrl, VarStringName, VarStringDir, VarStringMethod, VarStringUrl)
	genRoute(c)

	wd = filepath.Join(wd, "app")
	c.wd = wd

	genTypes(c)

	genHandler(c)

	genLogic(c)

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
