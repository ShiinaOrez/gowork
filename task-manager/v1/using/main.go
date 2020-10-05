package main

import (
	"fmt"
	task_manager "github.com/ShiinaOrez/gowork/task-manager/v1"
	"time"
)

type Job struct {
	ID   string
	Time time.Duration
}

func (job Job) GetID() string {
	return job.ID
}

func (job Job) Do() error {
	time.Sleep(job.Time)
	return nil
}

func main() {
	go task_manager.RUN(time.Duration(4 * time.Second))

	time.Sleep(time.Second)

	task := task_manager.CreateTask(5, 0)
	fmt.Println("[Log]: Start to add jobs.")
	job1 := Job{
		ID:   "JOB-1",
		Time: time.Second,
	}
	job2 := Job{
		ID:   "JOB-2",
		Time: 3 * time.Second,
	}
	job3 := Job{
		ID:   "JOB-3",
		Time: 2 * time.Second,
	}
	job4 := Job{
		ID:   "JOB-4",
		Time: time.Second + 500 * time.Millisecond,
	}
	job5 := Job{
		ID:   "JOB-5",
		Time: 2 * time.Second + 500 * time.Millisecond,
	}
	task.AddJob(job1)
	task.AddJob(job2)
	task.AddJob(job3)
	task.AddJob(job4)
	task.AddJob(job5)
	time.Sleep(20 * time.Second)
}