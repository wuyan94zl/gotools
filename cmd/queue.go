package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/queuecmd"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "create queue script",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		queue := &queuecmd.Command{}
		queue.Run()
	},
}

func init() {
	queueCmd.Flags().StringVarP(&queuecmd.VarStringName, "name", "n", "", "定义队列名称")
	rootCmd.AddCommand(queueCmd)
}
