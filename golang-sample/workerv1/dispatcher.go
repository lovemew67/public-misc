package workerv1

import (
	"time"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

var (
	dispatcherCtx  cornerstone.Context
	dispatcherStop = make(chan struct{}, 1)
	dispatcherDone = make(chan struct{}, 1)
)

func startDispatcherLoop(ctx cornerstone.Context, jobDispatch func(cornerstone.Context), checkInterval time.Duration) {
	dispatcherCtx = ctx.CopyContext()
	dispatcherCtx.Set("name", "dispatcher")
	go dispatcherLoop(jobDispatch, dispatcherDone, checkInterval)
}

func dispatcherLoop(jobDispatch func(cornerstone.Context), dispatcherDone chan struct{}, checkInterval time.Duration) {
	funcName := "dispatcherLoop"

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	defer cornerstone.Infof(dispatcherCtx, "[%s] closed dispatcher", funcName)
	defer func(d chan struct{}) { d <- struct{}{} }(dispatcherDone)

	cornerstone.Infof(dispatcherCtx, "[%s] start dispatcher", funcName)
	run := true
	for run {
		select {
		case <-dispatcherStop:
			cornerstone.Infof(dispatcherCtx, "[%s] dispatcher terminating", funcName)
			run = false
		case <-ticker.C:
			jobCtx := dispatcherCtx.CopyContext()
			jobDispatch(jobCtx)
		}
	}
	cornerstone.Infof(dispatcherCtx, "[%s] dispatcher done", funcName)
}

func jobDispatch(ctx cornerstone.Context) {
	funcName := "jobDispatch"
	cornerstone.Debugf(ctx, "[%s] triggered", funcName)

	// get available jobs
	acceptableJobs := listAvailableJobs(ctx)
	if len(acceptableJobs) == 0 {
		cornerstone.Debugf(ctx, "[%s] oh no", funcName)
		return
	}
}

func listJobs(ctx cornerstone.Context) (jobs []*domainv1.Job, err error) {
	// db operations
	return
}

func listAvailableJobs(ctx cornerstone.Context) (acceptableJobs []*domainv1.Job) {
	funcName := "listAvailableJobs"

	acceptableJobs, err := listJobs(ctx)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to list jobs, err: %+v", funcName, err)
		return
	}

	// list jobs
	return
}
