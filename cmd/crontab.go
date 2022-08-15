package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/crontab"
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
		app := &crontab.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	cronCmd.Flags().StringVarP(&crontab.VarStringName, "name", "n", "", "")
	cronCmd.Flags().StringVarP(&crontab.VarStringDir, "dir", "d", "", "")
	rootCmd.AddCommand(cronCmd)
}
