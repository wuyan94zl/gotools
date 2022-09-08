package handlercmd

import (
	"errors"
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
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
	err := validateHandlerFlags()
	if err != nil {
		return err
	}
	c.Command = fmt.Sprintf("%s --dir %s --name %s --method %s", c.Command, VarStringDir, VarStringName, VarStringMethod)
	if VarStringUrl != "" {
		c.Command = fmt.Sprintf("%s --url %s", c.Command, VarStringUrl)
	}
	if VarStringParams != "" {
		c.Command = fmt.Sprintf("%s --params %s", c.Command, VarStringParams)
	}
	wd, _ := os.Getwd()
	c.wd = wd
	c.handlerName = getName(VarStringDir) + getName(formatUrl(VarStringUrl)) + getName(VarStringName)
	c.routeUrl = fmt.Sprintf("%s%s%s%s", getUrl(VarStringDir), getUrl(VarStringUrl), getUrl(VarStringName), getUrl(VarStringParams))[1:]

	c.dir = VarStringDir
	c.method = VarStringMethod

	err = genRoute(c)
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

func validateHandlerFlags() error {
	utils.ToLowers(&VarStringDir, &VarStringName, &VarStringUrl, &VarStringParams)
	ok, err := regexp.MatchString("^([a-z]+)$", VarStringDir)
	if err != nil || !ok {
		return errors.New("the --dir parameter is invalid")
	}
	ok, err = regexp.MatchString("^([a-z]+)$", VarStringName)
	if err != nil || !ok {
		return errors.New("the --name parameter is invalid")
	}
	if VarStringMethod != "POST" && VarStringMethod != "GET" {
		return errors.New("the --method parameter is invalid, only GET or POST")
	}
	ok, err = regexp.MatchString("^([a-z/]+)$", VarStringUrl)
	if err != nil || !ok {
		return errors.New("the --url parameter is invalid")
	}
	ok, err = regexp.MatchString("^([a-z:/]+)$", VarStringParams)
	if err != nil || !ok {
		return errors.New("the --params parameter is invalid")
	}
	return nil
}
