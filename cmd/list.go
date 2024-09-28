package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ognevsd/time-tracker/pkg/tracker"
	"github.com/spf13/cobra"
)

var yesterday bool
var all bool
var thisWeek bool

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List tasks added to tracker",
	Long:    `Will list all tasks added to the tracker today. Use flags to modify output`,
	Run:     listRun,
}

func addTask(t *table.Writer, task *tracker.Task) {
	(*t).AppendRow(table.Row{task.Project,
		task.Name,
		task.Started,
		task.Finished,
		task.Duration})
}

func listRun(cmd *cobra.Command, args []string) {
	tasks, err := tracker.LoadTasks()
	if err != nil {
		fmt.Println(err)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Project", "Task name", "Started", "Finished", "Duration"})
	for _, task := range tasks {
		switch {
		case yesterday:
			if task.Started.Day() == time.Now().Day()-1 {
				addTask(&t, &task)
			}
		case thisWeek:
			if task.Started.Day() >= (time.Now().Day() - int(time.Now().Weekday()-1)) {
				addTask(&t, &task)
			}
		case all:
			addTask(&t, &task)
		default:
			if task.Started.Day() == time.Now().Day() {
				addTask(&t, &task)
			}
		}
	}
	t.SetAutoIndex(true)
	t.Render()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&yesterday, "yesterday", false, "Show tasks that were started yesterday")
	listCmd.Flags().BoolVar(&all, "all", false, "Show all tasks")
	listCmd.Flags().BoolVar(&thisWeek, "thisweek", false, "Show all tasks, that were started this week")
}
