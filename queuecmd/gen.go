package queuecmd

import (
	"os"
	"path/filepath"
)

var (
	VarStringName string
	VarStringDir  string
)

type Command struct {
	packageName string
	wd          string
}

func (c *Command) Run() error {
	wd, _ := os.Getwd()
	if VarStringDir != "." {
		wd = filepath.Join(wd, VarStringDir)
	}
	wd = filepath.Join(wd, VarStringName)
	_, packageName := filepath.Split(wd)

	c.packageName = packageName
	c.wd = wd
	if err := genQueueGen(c); err != nil {
		return err
	}
	return genQueue(c)
}
