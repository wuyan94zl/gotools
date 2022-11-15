package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/crontabcmd"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "crontab",
	Short: "generating crontab original code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := &crontabcmd.Command{
			Command: cmd.CommandPath(),
		}
		err := app.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			//commandLog(app.Command)
			fmt.Println(app.Command, "Down .")
		}
	},
}

func init() {
	cronCmd.Flags().StringVarP(&crontabcmd.VarStringName, "name", "n", "", "The crontab package name")
	rootCmd.AddCommand(cronCmd)
}
