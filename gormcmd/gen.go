package gormcmd

import (
	"errors"
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"github.com/wuyan94zl/sql2gorm/parser"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var VarStringSrc string
var VarStringDir string
var VarStringDeleted string
var VarBoolCache bool

type Command struct {
	Command       string
	packageName   string
	structName    string
	structData    string
	deletedFiled  string
	hasSoftDelete string
	wd            string
}

func (c *Command) GetDir() string {
	return ""
}

func (c *Command) Run() error {
	if VarStringSrc == "" {
		return errors.New("gorm --src is required")
	}
	if VarStringDir == "" {
		return errors.New("gorm --dir is required")
	}
	err := validateGormFlags()
	if err != nil {
		return err
	}
	c.Command = fmt.Sprintf("%s --src %s --dir %s --cache %v --deleted %s", c.Command, VarStringSrc, VarStringDir, VarBoolCache, VarStringDeleted)

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
	c.packageName = filepath.Base(VarStringDir)
	c.structData = structData.StructCode[0].Code
	c.structName = structData.StructCode[0].Table
	c.deletedFiled = VarStringDeleted
	c.wd = wd
	i := strings.Index(string(file), c.deletedFiled)
	if i == -1 {
		c.hasSoftDelete = "0"
	} else {
		c.hasSoftDelete = "1"
	}
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

func validateGormFlags() error {
	utils.ToLowers(&VarStringDir, &VarStringDeleted)
	ok, err := regexp.MatchString("^([a-z/]+)$", VarStringDir)
	if err != nil || !ok {
		return errors.New("the --dir parameter is invalid")
	}
	ok, err = regexp.MatchString("^([A-z/]+)$", VarStringSrc)
	if err != nil || !ok {
		return errors.New("the --src parameter is invalid")
	}
	ok, err = regexp.MatchString("^([a-z_]+)$", VarStringDeleted)
	if err != nil || !ok {
		return errors.New("the --deleted parameter is invalid")
	}
	return nil
}
