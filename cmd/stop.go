package cmd

import (
	"fmt"

	"github.com/ognevsd/time-tracker/pkg/tracker"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop last added task",
	Long:  `When work on the task is finished, it can be stopped using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tracker.LoadTasks()
		if err != nil {
			fmt.Println(err)
		}
		err = tasks.StopLastTask()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
