package cmd

import (
	"fmt"
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
		app := &queuecmd.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	queueCmd.Flags().StringVarP(&crontabcmd.VarStringName, "name", "n", "", "")
	queueCmd.Flags().StringVarP(&crontabcmd.VarStringDir, "dir", "d", "", "")
	rootCmd.AddCommand(queueCmd)
}
