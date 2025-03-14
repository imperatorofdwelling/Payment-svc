package scheduler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler struct {
	cron *cron.Cron
	log  *zap.SugaredLogger
}

func NewScheduler(log *zap.SugaredLogger) *Scheduler {
	c := cron.New()
	return &Scheduler{c, log}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Create(pattern string) {
	s.cron.AddFunc(pattern, func() {
		fmt.Println("Scheduler works")
	})
}

func (s *Scheduler) Stop() {
	if err := s.cron.Stop(); err != nil {
		panic(err)
	} else {
		s.log.Info("Scheduler stopped")
	}
}

//"0 22 * * *" example cron
