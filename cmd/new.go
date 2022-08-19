package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/newcmd"
)

// cronCmd represents the cron command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Initialize the new project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := &newcmd.Command{}
		err := app.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Down .")
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&newcmd.VarStringName, "name", "n", "", "The project name")
	newCmd.Flags().StringVarP(&newcmd.VarStringPackageName, "package", "p", "", "The main package name")
	rootCmd.AddCommand(newCmd)
}
