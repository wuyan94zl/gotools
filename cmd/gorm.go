package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/gorm"
)

// gormCmd represents the model command
var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "mysql ddl to gorm model code",
	Long:  `mysql ddl to gorm model code`,
	Run: func(cmd *cobra.Command, args []string) {
		model := &gorm.Command{}
		err := model.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Down .")
		}
	},
}

func init() {
	gormCmd.Flags().StringVarP(&gorm.VarStringSrc, "src", "s", "", "The path or path globbing patterns of the ddl")
	gormCmd.Flags().StringVarP(&gorm.VarStringDir, "dir", "d", "", "The path or path globbing patterns of the ddl")
	gormCmd.Flags().BoolVarP(&gorm.VarBoolCache, "cache", "c", false, "is cache")
	rootCmd.AddCommand(gormCmd)
}
