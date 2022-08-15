package cmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/crontab"
	"github.com/wuyan94zl/gotools/queue"

	"github.com/spf13/cobra"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "create queue script",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("model name not null")
			return
		}
		app := &queue.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	queueCmd.Flags().StringVarP(&crontab.VarStringName, "name", "n", "", "")
	queueCmd.Flags().StringVarP(&crontab.VarStringDir, "dir", "d", "", "")
	rootCmd.AddCommand(queueCmd)
}
