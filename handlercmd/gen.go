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
	Command     string
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
	c.Command = fmt.Sprintf("%s --dir %s --name %s --method %s", c.Command, VarStringDir, VarStringName, VarStringMethod)
	if VarStringUrl != "" {
		c.Command = fmt.Sprintf("%s --url %s", c.Command, VarStringUrl)
	}
	wd, _ := os.Getwd()
	c.wd = wd

	c.handlerName = getName(VarStringDir) + getName(formatUrl(VarStringUrl)) + getName(VarStringName)
	c.routeUrl = fmt.Sprintf("%s%s%s", getUrl(VarStringDir), getUrl(VarStringName), getUrl(VarStringUrl))[1:]

	c.dir = VarStringDir
	c.method = VarStringMethod

	err := genRoute(c)
	if err != nil {
		return err
	}

	wd = filepath.Join(wd, "app")
	c.wd = wd

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
