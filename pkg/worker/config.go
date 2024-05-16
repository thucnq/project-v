package worker

import (
	"context"
	"sync"

	"project-v/pkg/broker/asynq"
	"project-v/pkg/l"
	"project-v/pkg/worker/process-status-code"

	"github.com/robfig/cron/v3"
)

var (
	AllowedMaxRetries int = 100
	DefaultRetries    int = 10
)

const ScheduleFailed = 1

const (
	OneTimeMode = "once"
	EndlessMode = "endless"
)

// ProcessStatus ...
type ProcessStatus struct {
	Code    processstatuscode.Status
	Message []byte
}

var (
	ProcessOK            = ProcessStatus{Code: processstatuscode.Success}
	ProcessFailRetry     = ProcessStatus{Code: processstatuscode.Retry}
	ProcessFailDrop      = ProcessStatus{Code: processstatuscode.Drop}
	ProcessFailReproduce = ProcessStatus{Code: processstatuscode.FailReproduce}
)

// ScheduleJobOption ...
type ScheduleJobOption struct {
	id           cron.EntryID // the entry id of job
	spec         string       // the schedule string
	pl           ISchedulerTask
	retries      int    // the number of retries
	scheduleType string // the schedule type can be: once, endless
	finished     bool
	status       int
	name         string
}

// Scheduler ...
type Scheduler struct {
	ScheduleJob   []*ScheduleJobOption
	ctx           context.Context
	lock          sync.Mutex
	c             *cron.Cron
	asynqConsumer asynq.IConsumer
	ll            l.Logger
	cfg           ScheduleConfig
}

// Consumer ...
type Consumer struct {
	ctx         context.Context
	lock        sync.Mutex
	ll          l.Logger
	ConsumerJob []*ConsumerGroup
}

type ScheduleConfig struct {
	Asynq asynq.Config
}
