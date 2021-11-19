package workerv1

import (
	"time"

	"github.com/lovemew67/public-misc/cornerstone"
)

type scheduleConfig struct {
	BatchSize     int
	CheckInterval time.Duration
}

type scheduler struct {
	ErrorHandler func(cornerstone.CodeError)
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
	cornerstone.Debugf(ctx, "[InitAutoDeletionScheduler] init scheduleConfig: %+v", funcName, sfg)

	// default execute metrics funcs every 500 ms
	cycle := sfg.CheckInterval
	if len(executeCycle) > 0 {
		cycle = executeCycle[0]
	}
	cornerstone.Debugf(ctx, "[%s] with executeCycle: %+v", funcName, cycle)

	tick := time.NewTicker(cycle)
	go func(tick *time.Ticker) {
		ctx := cornerstone.Background()
		ctx.Set("worker", "scheduler")
		for range tick.C {
			scheduleHandler(ctx)
		}
	}(tick)
	return tick
}

func scheduleHandler(ctx cornerstone.Context) {
	// funcName := "scheduleHandler"

	// autoDeletionScheduleOperator := dbaccessv1.GetAutoDeletionScheduleOperator()
	// qa := &dbaccessv1.QueryArgs{
	// 	StatusRunning: false,
	// 	Enable:        true,
	// }
	// sfs := dbaccessv1.GetSelectFields([]string{}, []string{})
	// autoDeletionScheduleList, total := autoDeletionScheduleOperator.Find(ctx, qa, sfs, 0, autoDeletionScheduleConfig.BatchSize)
	// if total == 0 {
	// 	return
	// }

	// for idx := range autoDeletionScheduleList {
	// autoDeletionSchedule := &autoDeletionScheduleList[idx]

	// == block concurrent dispatching
	// qa := &dbaccessv1.QueryArgs{
	// 	StatusRunning: false,
	// 	Enable:        true,
	// 	Service:       autoDeletionSchedule.Service,
	// }
	// patchMap := bson.M{
	// 	modelv1.FdStatusRunning: true,
	// }
	// err := autoDeletionScheduleOperator.Patch(ctx, qa, patchMap)
	// if err != nil {
	// 	m800log.Infof(ctx, "[%s] failed to patch auto deletion schedule, qa: %+v, err: %+v", funcName, qa, err)
	// 	continue
	// }

	// == do something

	// == unblock
	// qa.StatusRunning = true
	// patchMap = bson.M{
	// 	modelv1.FdStatusRunning: false,
	// }
	// err = autoDeletionScheduleOperator.Patch(ctx, qa, patchMap)
	// if err != nil {
	// 	m800log.Errorf(ctx, "[%s] failed to patch auto deletion schedule, err: %+v", funcName, err)
	// }
	// }
}
