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

func (s *Scheduler) Run() {
	s.cron.AddFunc("0 22 * * *", func() {
		fmt.Println("Scheduler works")
	})

	s.cron.Start()
}

func (s *Scheduler) Stop() {
	if err := s.cron.Stop(); err != nil {
		panic(err)
	} else {
		s.log.Info("Scheduler stopped")
	}
}
