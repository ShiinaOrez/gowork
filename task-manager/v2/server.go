package task_manager

import test_mq "github.com/ShiinaOrez/gowork/task-manager/v2/test-mq"

type Server struct {
	logger     Logger
	broker     test_mq.MQ
	handlerMap HandlerMap
}

type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func NewServer(topicName, cluster string) (server Server) {
	server.broker = test_mq.GetMQFromConf(test_mq.MqConf{
		TopicName: topicName,
		Cluster:   cluster,
	})
	server.handlerMap = make(map[string]Handler)
	return
}

func (server *Server) Start(hm HandlerMap) {
	// go consume msg
}

type Handler func(arg Argument) error

type HandlerMap map[string]Handler