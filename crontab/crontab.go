package crontab

import (
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func newCron() *cron.Cron {
	if Cron != nil {
		Cron = cron.New(cron.WithSeconds())
	}
	return Cron
}

func AddJob(spec string, job cron.Job) {
	newCron().AddJob(spec, job)
}

func NewInstance() *cron.Cron {
	return Cron
}
