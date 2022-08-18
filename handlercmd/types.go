package handlercmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"path/filepath"
	"strings"
)

var typesTpl = `package {{.package}}

type {{.name}}Request struct {
}

type {{.name}}Response struct {
}

`

func genTypes(c *Command) error {
	//packageStr, err := utils.GetPackage()
	//if err != nil {
	//	return err
	//}
	wd := filepath.Join(c.wd, "types")
	//childDir := ""
	//if c.dir != "" {
	//	childDir = "/" + c.dir
	//}
	//typePackage := fmt.Sprintf("%s/%s%s", packageStr, "app/types", childDir)

	fmt.Println(c.name, "name")
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          wd,
		Filename:     c.name + ".go",
		TemplateFile: typesTpl,
		Data: map[string]string{
			"package": filepath.Base(wd),
			"name":    strings.ToUpper(c.name[0:1]) + c.name[1:],
		},
	})
}
