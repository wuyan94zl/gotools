package crontabcmd

import (
	"errors"
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"os"
	"path/filepath"
	"regexp"
)

var (
	VarStringName string
)

type Command struct {
	Command     string
	packageName string
	wd          string
}

func (c *Command) Run() error {
	if VarStringName == "" {
		return errors.New("crontab name is required")
	}
	err := validateFlags()
	if err != nil {
		return err
	}
	c.Command = fmt.Sprintf("%s --name %s", c.Command, VarStringName)

	wd, _ := os.Getwd()
	wd = filepath.Join(wd, "crontab")

	c.packageName = VarStringName
	c.wd = filepath.Join(wd, VarStringName)
	genJob(c)

	c.packageName = "crontab"
	c.wd = wd
	genBase(c)
	return nil
}

func validateFlags() error {
	utils.ToLowers(&VarStringName)
	ok, err := regexp.MatchString("^([a-z/]+)$", VarStringName)
	if err != nil || !ok {
		return errors.New("the --name parameter is invalid")
	}
	return nil
}
