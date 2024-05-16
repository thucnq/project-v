package worker

import "encoding/json"

// ISchedulerTask ...
type ISchedulerTask interface {
	// GetName returns the task name
	GetName() string
	// GetNameWithSuffix returns the task name with suffix
	GetNameWithSuffix() string
	// Before the func call when task exec
	Before()
	// Handle the func process task logic
	Handle() ProcessStatus
	// After the func call when task exec done
	After()
}

// IConsumerTask ...
type IConsumerTask interface {
	Handle(msg []byte) ProcessStatus
}

// HandlerWOption ...
type HandlerWOption struct {
	MsgHandler
	Replica int64
}

// MsgHandler ...
type MsgHandler interface {
	Handle(msg []byte) ProcessStatus
}

// Event ...
type Event struct {
	ExchangeName   string           `json:"exchangeName"`
	RoutingKey     string           `json:"routingKey"`
	IssueAt        int64            `json:"issueAt"`
	Issuer         string           `json:"issuer"`
	MessageVersion string           `json:"messageVersion"`
	RawMessage     *json.RawMessage `json:"message"`
}

type IWorker interface {
	Start()
}

// Worker ...
type worker struct {
	scheduler *Scheduler
	consumer  *Consumer
	//consumer  *Task
	// producer
	//stopped bool
}

func (w *worker) Start() {
	if w.scheduler != nil {
		go w.scheduler.Start()
	}
	if w.consumer != nil {
		go w.consumer.Start()
	}
}
