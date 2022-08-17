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
		} else {
			fmt.Println("Down .")
		}
	},
}

func init() {
	gormCmd.Flags().StringVarP(&gormcmd.VarStringSrc, "src", "s", "", "The path or path globbing patterns of the ddl")
	gormCmd.Flags().StringVarP(&gormcmd.VarStringDir, "dir", "d", "", "The generated path")
	gormCmd.Flags().BoolVarP(&gormcmd.VarBoolCache, "cache", "c", false, "The model is set to cache")
	rootCmd.AddCommand(gormCmd)
}
