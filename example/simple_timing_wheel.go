package main

import (
	"fmt"
	"log"
	"time"

	tw "github.com/nicolerobin/timingwheel"
)

type SimpleTask struct {
	Name string
}

func NewSimpleTask(taskName string) SimpleTask {
	return SimpleTask{
		Name: taskName,
	}
}

func (st SimpleTask) Execute(t tw.Task) error {
	log.Printf("task name:%s\n", st.Name)
	return nil
}

func main() {
	fmt.Println("vim-go")
	tw1 := tw.NewSimpleTimingWheel(30)
	t1 := NewSimpleTask("simple task 1")
	tw1.Add(5, t1)

	t2 := NewSimpleTask("simple task 2")
	tw1.Add(15, t2)

	tw1.Start()

	time.Sleep(60 * time.Second)
}
