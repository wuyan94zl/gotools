package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/crontabcmd"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "crontab",
	Short: "create crontab script",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := &crontabcmd.Command{}
		err := app.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Down .")
		}
	},
}

func init() {
	cronCmd.Flags().StringVarP(&crontabcmd.VarStringName, "name", "n", "", "")
	rootCmd.AddCommand(cronCmd)
}
