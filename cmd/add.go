/*
Copyright © 2024 ognevds <https://github.com/ognevsd>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/ognevsd/time-tracker/pkg/tracker"
	"github.com/spf13/cobra"
)

var projectAdd string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name] [duration]",
	Short: "Use this command in case you want to add task with specific time",
	Long: `In case you just want to add task and time that you've spent on the
	task, this command can be used. Example input: 1h30m. Please, don't use
	fractional ad negative durations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(
				"Error: please, provide two arguments: task name and duration ",
				"in format 1h10m")
			cmd.Help()
		}

		tasks, err := tracker.LoadTasks()

		if err != nil {
			fmt.Println(err)
		}

		taskName := args[0]
		duration := args[1]

		if err := tasks.AddWithDuration(taskName, projectAdd, duration); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(
		&projectAdd,
		"project",
		"p",
		"generic",
		"Add a project to your task, it will help with assigning time to correct project",
	)
}
