package newcmd

import (
	"os"
	"os/exec"
)

var (
	VarStringPackageName string
)

type Command struct {
	packageSrc string
	wd         string
}

func (c *Command) Run() error {
	wd, _ := os.Getwd()
	c.wd = wd
	c.packageSrc = VarStringPackageName

	err := InitMod(c)
	if err != nil {
		return err
	}

	err = genConfig(c)
	if err != nil {
		return err
	}
	err = genContainer(c)
	if err != nil {
		return err
	}
	err = genRouter(c)
	if err != nil {
		return err
	}
	err = genMain(c)
	if err != nil {
		return err
	}
	return InitTidy()
}

func InitMod(c *Command) error {
	initCmd := []string{"mod", "init", c.packageSrc}
	cmd := exec.Command("go", initCmd...)
	return cmd.Start()
}

func InitTidy() error {
	initCmd := []string{"mod", "tidy"}
	cmd := exec.Command("go", initCmd...)
	return cmd.Start()
}
