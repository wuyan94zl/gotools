package gormcmd

import (
	"github.com/wuyan94zl/sql2gorm/parser"
	"io/ioutil"
	"os"
	"path/filepath"
)

var VarStringSrc string
var VarStringDir string
var VarBoolCache bool

type Command struct {
	packageName string
	structName  string
	structData  string
	wd          string
}

func (c *Command) GetDir() string {
	return ""
}

func (c *Command) Run() error {
	wd, _ := os.Getwd()
	if VarStringDir != "." {
		wd = filepath.Join(VarStringDir)
	}
	abs, err := filepath.Abs(VarStringSrc)
	if err != nil {
		return err
	}
	file, err := ioutil.ReadFile(abs)
	if err != nil {
		return err
	}
	structData, err := parser.ParseSql(string(file))
	if err != nil {
		return err
	}
	c.packageName = structData.Package
	c.structData = structData.StructCode[0].Code
	c.structName = structData.StructCode[0].Table
	c.wd = wd
	if VarBoolCache {
		if err := setGormModel(c); err != nil {
			return err
		}
		if err := setGormCustomModel(c); err != nil {
			return err
		}
	} else {
		if err := setGormNoCacheModel(c); err != nil {
			return err
		}
		if err := setGormNoCacheCustomModel(c); err != nil {
			return err
		}
	}

	return nil
}