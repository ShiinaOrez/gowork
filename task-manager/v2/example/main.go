package main

import (
	"fmt"
	task_manager "github.com/ShiinaOrez/gowork/task-manager/v2"
	"time"
)

func JobHandler1(arg task_manager.Argument) error {
	id, err := arg.GetInt("id")
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func StartServer() {
	server := task_manager.NewServer("test_kafka", "test.mq")
	hm := make(task_manager.HandlerMap)
	hm["job-1"] = JobHandler1
	server.Start(hm)
}

func main() {
	go StartServer()
	time.Sleep(10 * time.Second)
	c := task_manager.NewClient("test_kafka", "test.mq")
	job1 := task_manager.NewJob("job-1", map[string]interface{}{
		"id": 1123,
	})
	job2 := task_manager.NewJob("job-1", map[string]interface{}{
		"id": 2234,
	})
	task := task_manager.BuildTaskFromJobs(job1, job2)
	c.Do(task)
}