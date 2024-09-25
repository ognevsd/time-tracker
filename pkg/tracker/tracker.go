package tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var ErrEmptyTaskList = errors.New("tracker: Empty task list. Nothing to stop")
var ErrFinishedTask = errors.New("tracker: Last task is already finished")

type Task struct {
	Name     string         `json:"name"`
	Project  string         `json:"project"`
	Started  time.Time      `json:"started"`
	Finished *time.Time     `json:"finished"`
	Duration *time.Duration `json:"duration"`
}

type Tasks []Task

var tasksPath string = path.Join(".", "test.json")

func LoadTasks() (Tasks, error) {
	tasks := Tasks{}
	data, err := os.ReadFile(tasksPath)
	if errors.Is(err, os.ErrNotExist) {
		return tasks, nil
	} else if err != nil {
		return tasks, err
	}
	if err := json.Unmarshal(data, &tasks); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (t *Task) String() string {
	startFmt := t.Started.Format("2006-01-02 15:01:02")

	var finishedFmt string
	if t.Finished != nil {
		finishedFmt = t.Finished.Format("2006-01-02 15:03:03")
	} else {
		finishedFmt = "not finished"
	}

	var durationFmt string
	if t.Duration != nil {
		durationFmt = t.Duration.String()
	} else {
		durationFmt = "not finished"
	}

	return fmt.Sprintf("Task{name: %s, project: %s, started: %s, finished: %s, duration: %s}",
		t.Name, t.Project, startFmt, finishedFmt, durationFmt)
}

func (t *Tasks) String() string {
	var res strings.Builder

	for i, v := range *t {
		res.WriteString(v.String())
		if i != len(*t)-1 {
			res.WriteRune('\n')
		}
	}

	return res.String()
}

func (t *Tasks) Add(name, project string) {
	task := Task{
		Name:     name,
		Project:  project,
		Started:  time.Now(),
		Finished: nil,
		Duration: nil,
	}

	t.StopLastTask()

	*t = append(*t, task)
	t.save()
}

func (t *Tasks) StopLastTask() {
	if len(*t) > 0 {
		lastTask := &(*t)[len(*t)-1]
		completionTime := time.Now()
		taskDuration := completionTime.Sub((*t)[len(*t)-1].Started)
		lastTask.Finished = &completionTime
		lastTask.Duration = &taskDuration
	}
}

func (t *Tasks) save() {
	res, err := json.MarshalIndent(*t, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(tasksPath, res, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
