package newcmd

import (
	"os"
	"os/exec"
	"path/filepath"
)

var (
	VarStringName        string
	VarStringPackageName string
)

type Command struct {
	packageSrc string
	wd         string
}

func (c *Command) Run() error {
	wd, _ := os.Getwd()
	c.wd = wd

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
	InitTidy(c)
	return nil
}

func InitMod(c *Command) error {
	initCmd := []string{"mod", "init"}
	if VarStringPackageName != "" {
		initCmd = append(initCmd, VarStringPackageName)
		c.packageSrc = VarStringPackageName
	} else {
		c.packageSrc = filepath.Base(c.wd)
	}
	cmd := exec.Command("go", initCmd...)
	_, err := cmd.Output()
	return err
}

func InitTidy(c *Command) error {
	initCmd := []string{"mod", "tidy"}
	cmd := exec.Command("go", initCmd...)
	_, err := cmd.Output()
	return err
}
