package workerv1

import (
	"time"

	"github.com/lovemew67/public-misc/cornerstone"
)

type scheduleConfig struct {
	BatchSize     int
	CheckInterval time.Duration
}

func getAutoDeletionScheduleInfo() scheduleConfig {
	return scheduleConfig{
		BatchSize:     10,
		CheckInterval: 10 * time.Second,
	}
}

func InitScheduler(ctx cornerstone.Context, executeCycle ...time.Duration) *time.Ticker {
	funcName := "InitScheduler"

	sfg := getAutoDeletionScheduleInfo()
	cornerstone.Debugf(ctx, "[%s] init scheduleConfig: %+v", funcName, sfg)

	cycle := sfg.CheckInterval
	if len(executeCycle) > 0 {
		cycle = executeCycle[0]
	}
	cornerstone.Debugf(ctx, "[%s] with executeCycle: %+v", funcName, cycle)

	tick := time.NewTicker(cycle)
	go func(tick *time.Ticker) {
		ctx := cornerstone.NewContext()
		ctx.Set("worker", "scheduler")
		cornerstone.Debugf(ctx, "[%s] ticker routine inited", funcName)
		for range tick.C {
			scheduleHandler(ctx)
		}
	}(tick)
	return tick
}

func scheduleHandler(ctx cornerstone.Context) {
	// list valid schedules
	// for each schedule
	//  block concurrent dispatching
	//  do handling: (eg, insert job)
	//  unblock
}
