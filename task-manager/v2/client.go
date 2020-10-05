package task_manager

import (
	test_mq "github.com/ShiinaOrez/gowork/task-manager/v2/test-mq"
	"sync"
	"time"
)

type Client struct {
    mqConf  test_mq.MqConf
    optsMap map[int64][]Option
    locker  sync.Mutex
}

func NewClient(topicName, cluster string) (client Client) {
	client.mqConf.TopicName = topicName
	client.mqConf.Cluster = cluster
	client.optsMap = make(map[int64][]Option)
	return
}

type Option interface {}

type (
	retryOption   int
	timeoutOption time.Duration
)

var (
	defaultMaxRetry = 0
	defaultTimeout  = 10 * time.Minute
)

type option struct {
	retry     int
	timeout   time.Duration
}

func composeOptions(opts ...Option) option {
	newOpt := option{
		retry:   defaultMaxRetry,
		timeout: defaultTimeout,
	}
	for _, opt := range opts {
		switch opt := opt.(type) {
		case retryOption:
			newOpt.retry = int(opt)
		case timeoutOption:
			newOpt.timeout = time.Duration(opt)
		default:
		}
	}
	return newOpt
}

func (c *Client) Do(task Task) {
	// send msg
}