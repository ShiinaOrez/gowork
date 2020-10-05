package task_manager

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Task struct {
	TaskID     int64
	jobMap     map[string]Job
	logicMap   map[string][]string
	jmLocker   sync.Mutex
	lmLocker   sync.Mutex
}

type Job struct {
    ID      string
	Type    string
	Args    Argument
}

func NewJob(t string, arg Argument) (job Job) {
	uuid := uuid.New()
	job.ID = string(uuid[:])
	job.Type = t
	job.Args = arg
	return job
}

func NewTask() (task Task) {
	// new task from metadata
	task = BuildTaskFromJobs()
	return
}

func BuildTaskFromJobs(jobs ...Job) (task Task) {
	task = BuildTaskFromJobSlice(jobs)
	return
}

func BuildTaskFromJobSlice(jobs []Job) (task Task) {
	task.jobMap = make(map[string]Job)
	task.logicMap = make(map[string][]string)
	for idx, job := range jobs {
		task.jobMap[job.ID] = job
		if idx < len(jobs)-1 {
			task.logicMap[job.ID] = []string{jobs[idx+1].ID}
		}
	}
	return
}

