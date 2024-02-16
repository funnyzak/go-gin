package singleton

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	Cron     *cron.Cron
	CronLock sync.RWMutex
)

func InitCronTask() {
	Cron = cron.New(cron.WithSeconds(), cron.WithLocation(Loc))
}

func LoadCronTasks() {
	InitCronTask()

	Cron.Start()
}
