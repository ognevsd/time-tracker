package main

import (
	"fmt"
	"time"

	"github.com/ognevsd/time-tracker/pkg/tracker"
)

func main() {
	tasks := tracker.Tasks{}
	tasks.Add("Presntation", "PCL")
	time.Sleep(time.Second)
	tasks.Add("WeSave Pres", "Sales Support")
	tasks.Add("Some other task", "General work")
	fmt.Println(tasks.String())
}
