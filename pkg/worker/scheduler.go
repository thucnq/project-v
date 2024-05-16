package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"project-v/pkg/broker/asynq"
	"project-v/pkg/cronlog"
	"project-v/pkg/l"
	"project-v/pkg/worker/process-status-code"

	aq "github.com/hibiken/asynq"
	"github.com/robfig/cron/v3"
)

// NewScheduler ...
func NewScheduler(
	ctx context.Context, sj []*ScheduleJobOption, ll l.Logger,
	cfg ScheduleConfig,
) *Scheduler {
	cl := cronlog.NewCronLogger(ll)
	asynqConsumer, err := asynq.NewAsynqConsumer(cfg.Asynq)
	if err != nil {
		panic(err)
	}

	s := &Scheduler{
		cfg:         cfg,
		ScheduleJob: sj,
		ctx:         ctx,
		lock:        sync.Mutex{},
		c: cron.New(
			cron.WithParser(
				cron.NewParser(
					cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow,
				),
			),
			cron.WithChain(cron.SkipIfStillRunning(cl)),
			cron.WithLocation(time.FixedZone("UTC+7", 25200)),
		),
		asynqConsumer: asynqConsumer,
		ll:            ll,
	}

	return s
}

// Close ...
func (s *Scheduler) Close() {
	s.ll.Info("Scheduler closed")
}

// Set ...
func (s *Scheduler) WithHandler(
	name string, handler func(ctx context.Context, t *aq.Task) error,
) {
	fmt.Println("Set handler as", s.asynqConsumer)
	s.asynqConsumer.Handle(
		name, handler,
	)
}

// Start ...
func (s *Scheduler) Start() {
	if len(s.ScheduleJob) == 0 {
		return
	}
	s.ll.Info("Start scheduler")
	// add job to schedule
	for _, item := range s.ScheduleJob {
		jobCfg := item
		st := jobCfg.pl
		id, err := s.c.AddFunc(
			jobCfg.spec, func() {
				st.Before()
				s.ll.S.Infof("[%v] Running job ", st.GetNameWithSuffix())
				localMaxTry := jobCfg.retries

				if localMaxTry == 0 {
					localMaxTry = 1
				}

				if localMaxTry < 0 || localMaxTry > AllowedMaxRetries {
					localMaxTry = DefaultRetries
				}
				var retries int
				var shouldRetry = true

				for retries < localMaxTry && shouldRetry {
					retries++
					resp := st.Handle()
					code := processstatuscode.Status(resp.Code)
					switch code {
					case processstatuscode.Success:
						s.ll.S.Debugf(
							"[%v] Process job exit with [%v] return code at round number [%v]",
							st.GetNameWithSuffix(),
							code.String(),
							retries,
						)
						shouldRetry = false
					case processstatuscode.Drop:
						s.ll.S.Debugf(
							"[%v] Process job exit with [%v] return code at round number [%v]",
							st.GetNameWithSuffix(),
							code.String(),
							retries,
						)

						shouldRetry = false
					case processstatuscode.Retry:
						shouldRetry = true
					}
					s.ll.S.Debugf(
						"[%v] Process exit with code [%v] at number [%v]",
						st.GetNameWithSuffix(),
						code.String(),
						retries,
					)
				}
				if retries == localMaxTry && shouldRetry {
					s.ll.S.Debugf(
						"[%v] Process exit with error code at number [%v]. Give up",
						st.GetNameWithSuffix(),
						retries,
					)
				}

				s.ll.S.Debugf("[%v] Job finished", st.GetNameWithSuffix())
				st.After()
				jobCfg.finished = true
			},
		)
		if err != nil {
			s.ll.S.Debugf(
				"[%v] Cannot schedule job due to error [%v]",
				st.GetNameWithSuffix(), err,
			)
			jobCfg.finished = true
			jobCfg.status = ScheduleFailed // ready to remove from job list
		} else {
			jobCfg.id = id
		}
	}

	s.c.Start()

	tick := time.NewTicker(1 * time.Second)

	go func() {
		s.asynqConsumer.Start()
	}()

	defer s.c.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-tick.C:
			s.shouldTerminate()
		}
	}
}

func (s *Scheduler) shouldTerminate() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.ScheduleJob) == 0 {
		return
	}

	var removedItems []int
	for idx, v := range s.ScheduleJob {
		if v.status == ScheduleFailed {
			removedItems = append(removedItems, idx)
			continue
		}
		switch v.scheduleType {
		case OneTimeMode:
			if v.finished {
				removedItems = append(removedItems, idx)
			}
		default:
			break
		}
	}

	for _, idx := range removedItems {
		sj := s.ScheduleJob[idx]
		s.c.Remove(sj.id)
	}
	if len(removedItems) > 0 {
		s.ll.S.Infof("Removed %v jobs", len(removedItems))
	}
}
