package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/gormcmd"
)

// gormCmd represents the model command
var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "mysql ddl to generating  gorm model original code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		model := &gormcmd.Command{}
		err := model.Run()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	gormCmd.Flags().StringVarP(&gormcmd.VarStringSource, "source", "s", "root:123456@tcp(localhost:3306)/blogs", "Default root:123456@tcp(localhost:3306)/blogs")
	gormCmd.Flags().StringVarP(&gormcmd.VarStringDir, "dir", "d", "", "")
	gormCmd.Flags().StringVarP(&gormcmd.VarTable, "table", "t", "", "")
	rootCmd.AddCommand(gormCmd)
}
