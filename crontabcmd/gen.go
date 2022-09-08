package crontabcmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
