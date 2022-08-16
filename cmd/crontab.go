package cmd

import (
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
		app.Run()
	},
}

func init() {
	cronCmd.Flags().StringVarP(&crontabcmd.VarStringName, "name", "n", "", "")
	cronCmd.Flags().StringVarP(&crontabcmd.VarStringDir, "dir", "d", "", "")
	rootCmd.AddCommand(cronCmd)
}
