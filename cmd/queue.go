package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/gotools/queuecmd"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "generating queue original code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		queue := &queuecmd.Command{
			Command: cmd.CommandPath(),
		}
		err := queue.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			//commandLog(queue.Command)
			fmt.Println(queue.Command, "Down .")
		}
	},
}

func init() {
	queueCmd.Flags().StringVarP(&queuecmd.VarStringName, "name", "n", "", "The queue package name")
	rootCmd.AddCommand(queueCmd)
}
