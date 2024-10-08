package cmd

import (
	"fmt"

	"github.com/ognevsd/time-tracker/pkg/tracker"
	"github.com/spf13/cobra"
)

var project string

var startCmd = &cobra.Command{
	Use:   "start [task name] [flags]",
	Short: "Start time tracking for new task",
	Long: `When you start working on the new task, it can be added to the tracker
	using the _start_ command. Project can be specified using --project flag`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: missing task name that should be added to tracking")
			cmd.Help()
		}
		tasks, err := tracker.LoadTasks()
		if err != nil {
			fmt.Println(err)
		}
		taskName := args[0]
		tasks.Add(taskName, project)
		fmt.Printf("Added: %s, with project: %s\n", taskName, project)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"generic",
		"Add a project to your task, it will help with assigning time to correct project",
	)
}
