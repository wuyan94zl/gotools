package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/handlercmd"
)

// cronCmd represents the cron command
var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "Generate code based on template files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := &handlercmd.Command{}
		err := app.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Down .")
		}
	},
}

func init() {
	handlerCmd.Flags().StringVarP(&handlercmd.VarStringName, "name", "n", "", "The handler name")
	handlerCmd.Flags().StringVarP(&handlercmd.VarStringDir, "dir", "d", "", "The handler path")
	handlerCmd.Flags().StringVarP(&handlercmd.VarStringMethod, "method", "m", "", "The handler method")
	handlerCmd.Flags().StringVarP(&handlercmd.VarStringUrl, "url", "u", "", "The handler uri")
	rootCmd.AddCommand(handlerCmd)
}
