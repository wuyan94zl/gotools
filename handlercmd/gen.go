package handlercmd

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
	dir         string
	name        string
}

func (c *Command) Run() error {
	wd, _ := os.Getwd()
	wd = filepath.Join(wd, "app")
	c.wd = wd
	if VarStringDir == "." {
		VarStringDir = ""
	}
	c.dir = VarStringDir
	c.name = VarStringName

	genTypes(c)

	genHandler(c)

	genLogic(c)

	return nil
}
