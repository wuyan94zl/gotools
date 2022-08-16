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
		if len(args) == 0 {
			fmt.Println("cron name not null")
			return
		}
		app := &crontabcmd.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	cronCmd.Flags().StringVarP(&crontabcmd.VarStringName, "name", "n", "", "")
	cronCmd.Flags().StringVarP(&crontabcmd.VarStringDir, "dir", "d", "", "")
	rootCmd.AddCommand(cronCmd)
}
