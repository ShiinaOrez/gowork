package task_manager

import (
	"fmt"
	"sync"
	"time"
)

const (
	DoJob      = "DoJob"
	FinishJob  = "FinishJob"
	FinishTask = "FinishTask"
)

var (
	DefaultUpdateTimeout time.Duration
    MetadataDB           []Metadatatask
    MessageQueue         chan Message

    TaskManagerDB        []*Task

	FOREVER              = 24 * time.Hour
)

type Metadatatask struct {
	ID       int64
	Progress int64
	Total    int64
	Done     int64
}

type Task struct {
	TaskID        int64
	Done          bool
	JobDoneTotal  int64
	JobTotal      int64
	UpdateTimeout time.Duration
	WG            sync.WaitGroup
	LastUpdatedTS time.Time
}

type Job interface {
	GetID() string
	Do()    error
}

type Message struct {
	Task   *Task
	DoJob      *Job
	Type       string
	Error      error
}

func metadataCreateTask(tot int64) int64 {
	id := int64(len(MetadataDB))
	MetadataDB = append(MetadataDB, Metadatatask{
		ID:       id,
		Progress: 0,
		Total:    tot,
		Done:     0,
	})
	return id
}

func metadataUpdateTask(taskID, doneTotal int64) {
	MetadataDB[taskID].Progress = int64(float64(doneTotal)/float64(MetadataDB[taskID].Total) * 100)
	fmt.Printf("[TaskID]: %d new progress is %d.\n", taskID, MetadataDB[taskID].Progress)
	MetadataDB[taskID].Done = doneTotal
	return
}

func CreateTask(total int64, updateTimeout time.Duration) *Task {
	if updateTimeout == 0 {
		updateTimeout = DefaultUpdateTimeout
	}
	taskID := metadataCreateTask(total)
	task := Task{
		TaskID:        taskID,
		JobDoneTotal:  0,
		JobTotal:      total,
		Done:          false,
		UpdateTimeout: updateTimeout,
		WG:            sync.WaitGroup{},
		LastUpdatedTS: time.Now(),
	}
	task.WG.Add(int(total))
	TaskManagerDB = append(TaskManagerDB, &task)
	go func() {
		task.WG.Wait()
		task.Done = true
		fmt.Printf("[TaskID]: %d done.\n", taskID)
	}()
	return &task
}

func (t *Task) AddJob(job Job) {
	MessageQueue<- Message{
		Task:     t,
		DoJob:    &job,
		Type:     DoJob,
		Error:    nil,
	}
}

func (t *Task) DoneJob() {
	t.JobDoneTotal += 1
	t.LastUpdatedTS = time.Now()
	metadataUpdateTask(t.TaskID, t.JobDoneTotal)
	if !t.Done {
		t.WG.Done()
	}
}

func handleMessage(message Message) {
	t := message.Task
	switch message.Type {
	case FinishTask: {
		for i:=t.JobDoneTotal; i<t.JobTotal; i++ {
			t.WG.Done()
		}
	}
	case FinishJob: {
		t.DoneJob()
	}
	case DoJob: {
		go func() {
			err := (*message.DoJob).Do()
			if err != nil {
				fmt.Printf("[TaskID]: %d, [JobID]: %s failed, err is %#+v.\n", t.TaskID, (*message.DoJob).GetID(), err)
			}
			MessageQueue<- Message{
				Task: t,
				Error:    err,
				Type:     FinishJob,
			}
		}()
	}
	}
}

func allTaskTimeoutMonitor() {
	for {
		for i, task := range TaskManagerDB {
			time.Sleep(500 * time.Millisecond)
			if time.Now().Sub(task.LastUpdatedTS) > task.UpdateTimeout && !task.Done{
				// 超时强行终止
				fmt.Printf("[TaskID]: %d timeout.\n", task.TaskID)
				TaskManagerDB[i].Done = true
				MessageQueue<- Message{
					Task: TaskManagerDB[i],
					Type:     FinishTask,
					Error:    nil,
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func RUN(defaultUpdateTimeout time.Duration) {
	fmt.Println("TaskManager Start... ")

	DefaultUpdateTimeout = defaultUpdateTimeout
	MessageQueue = make(chan Message, 100)
	defer close(MessageQueue)

	go allTaskTimeoutMonitor()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			select {
			case Message := <- MessageQueue: {
				handleMessage(Message)
			}
			}
		}
	}()
	time.Sleep(FOREVER)
}