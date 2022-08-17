package crontabcmd

import (
	"os"
	"path/filepath"
)

var (
	VarStringName string
)

type Command struct {
	packageName string
	wd          string
}

func (c *Command) Run() error {
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
