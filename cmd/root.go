package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotools",
	Short: "Code generator tools",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func commandLog(command string) error {
	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "command.log")
	_, err := os.Stat(filePath)
	if err != nil {
		file, _, _ := utils.CreteFile(wd, "command.log")
		file.Close()
	}
	file, err := ioutil.ReadFile(filePath)
	fileStr := string(file)
	i := strings.Index(fileStr, command)
	if i == -1 {
		fileStr = fmt.Sprintf("%s%s\n", fileStr, command)
	}
	fp, _ := os.Create(filePath)
	defer fp.Close()
	_, err = fp.WriteString(fileStr)
	return err
}
