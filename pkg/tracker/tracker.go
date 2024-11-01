package tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func getFilePath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := user.HomeDir
	tasksPath := path.Join(homeDir, ".time-tracker.json")
	return tasksPath
}

var ErrEmptyTaskList = errors.New("tracker: Empty task list. Nothing to stop")
var ErrFinishedTask = errors.New("tracker: Last task is already finished")
var ErrNegativeDuration = errors.New("tracker: Duration shouldn't contain negative values")

type Task struct {
	Name     string         `json:"name"`
	Project  string         `json:"project"`
	Started  time.Time      `json:"started"`
	Finished *time.Time     `json:"finished"`
	Duration *time.Duration `json:"duration"`
}

type Tasks []Task

func LoadTasks() (Tasks, error) {
	tasks := Tasks{}
	data, err := os.ReadFile(getFilePath())
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
	startFmt := t.Started.Format(timeFormat)

	var finishedFmt string
	if t.Finished != nil {
		finishedFmt = t.Finished.Format(timeFormat)
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

func (t *Tasks) AddWithDuration(name, project, duration string) error {
	if strings.Contains(duration, "-") {
		return ErrNegativeDuration
	}

	parsedDuration, err := time.ParseDuration(duration)
	if err != nil {
		fmt.Println(err)
	}

	task := Task{
		Name:     name,
		Project:  project,
		Started:  time.Now(),
		Finished: nil,
		Duration: &parsedDuration,
	}

	*t = append(*t, task)
	t.save()

	return nil
}

func (t *Tasks) StopLastTask() error {
	if len(*t) == 0 {
		return ErrEmptyTaskList
	}

	lastTask := &(*t)[len(*t)-1]

	if lastTask.Finished != nil {
		return ErrFinishedTask
	}

	completionTime := time.Now()
	taskDuration := completionTime.Sub((*t)[len(*t)-1].Started)
	lastTask.Finished = &completionTime
	lastTask.Duration = &taskDuration

	t.save()

	return nil
}

func (t *Tasks) save() {
	res, err := json.MarshalIndent(*t, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(getFilePath(), res, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
