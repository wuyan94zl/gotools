package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type FileGenConfig struct {
	Dir          string
	Filename     string
	TemplateFile string
	Data         interface{}
}

func GenFileCover(c FileGenConfig) error {
	fp, _, err := CreteFile(c.Dir, c.Filename)
	defer fp.Close()
	if err != nil {
		return err
	}
	text := c.TemplateFile
	var t = template.Must(template.New("name").Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.Data)
	if err != nil {
		return err
	}
	code := FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func GenFile(c FileGenConfig) error {
	fp, created, err := MaybeCreteFile(c.Dir, c.Filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()
	text := c.TemplateFile
	var t = template.Must(template.New("name").Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.Data)
	if err != nil {
		return err
	}
	code := FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func FormatCode(code string) string {
	ret, err := format.Source([]byte(code))
	if err != nil {
		return code
	}
	return string(ret)
}

func MaybeCreteFile(dir string, fileName string) (fp *os.File, created bool, err error) {
	if _, err = os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0777)
	}
	fPath := filepath.Join(dir, fileName)
	_, err = os.Stat(fPath)
	if err == nil {
		return nil, false, nil
	}
	fp, err = os.Create(fPath)
	created = err == nil
	return
}

func CreteFile(dir string, fileName string) (fp *os.File, created bool, err error) {
	if _, err = os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0777)
	}
	fPath := filepath.Join(dir, fileName)
	fp, err = os.Create(fPath)
	created = err == nil
	return
}

func GetDir(method string, name string) string {
	baseDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(baseDir, "app", method, name)
}

func GetPackage(wd string) (string, error) {
	open, _ := os.Open(filepath.Join(wd, "go.mod"))
	defer open.Close()

	buf := bufio.NewReader(open)
	line, _, _ := buf.ReadLine()
	if len(line) < 7 {
		return "", errors.New("执行位置错误")
	}
	return string(line)[7:], nil
}

func WriteInfile(filePath, code string) error {
	fp, _ := os.Create(filePath)
	defer fp.Close()
	code = FormatCode(code)
	_, err := fp.WriteString(code)
	return err
}

func UpperOne(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

func ToLowers(str ...*string) {
	for _, v := range str {
		strings.ToLower(*v)
	}
}

func ToUppers(str ...*string) {
	for _, v := range str {
		strings.ToUpper(*v)
	}
}

func AppendFileCode(filePath, searchCode, addCode, find string) (string, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	str := string(file)
	return AppendStrCode(str, searchCode, addCode, find)
}

func AppendStrCode(str, searchCode, addCode, find string) (string, error) {
	i := strings.Index(str, searchCode)
	if i == -1 {
		point := strings.Index(str, find)
		fileStr := fmt.Sprintf("%s%s\n\t%s", str[0:point-1], addCode, str[point-1:])
		return fileStr, nil
	}
	return str, nil
}
