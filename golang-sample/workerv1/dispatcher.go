package workerv1

import (
	"time"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

const (
	defaultMaxQueryReadyJobs = 10
	minRetryCount            = 1
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

	acceptableJobs := listAvailableJobs(ctx)
	if len(acceptableJobs) == 0 {
		return
	}

	for idx := range acceptableJobs {
		acceptableJob := &acceptableJobs[idx]
		err := jr.UpdateProcessStatusToOngoing(acceptableJob.ID)
		if err != nil {
			cornerstone.Errorf(ctx, "[%s] failed to make job: %s, to ongoing, err: %+v", funcName, acceptableJob.ID, err)
			continue
		}
		cornerstone.Debugf(ctx, "[%s] going to dispatch job: %s", funcName, acceptableJob.ID)
		switch acceptableJob.Type {
		case domainv1.JobTypeA:
			jobTypeAWorkerQueue <- acceptableJob
		case domainv1.JobTypeB:
			jobTypeBWorkerQueue <- acceptableJob
		default:
			cornerstone.Errorf(ctx, "[%s] invalid job type: %d", funcName, acceptableJob.Type)
			_ = jr.CancelJobByID(acceptableJob.ID)
		}
	}
}

func listJobs(ctx cornerstone.Context) (jobs []domainv1.Job, err error) {
	jobs, err = jr.QueryReadyJobs(defaultMaxQueryReadyJobs)
	return
}

func listAvailableJobs(ctx cornerstone.Context) (acceptableJobs []domainv1.Job) {
	funcName := "listAvailableJobs"

	readyJobs, err := listJobs(ctx)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to list jobs, err: %+v", funcName, err)
		return
	}

	for i := range readyJobs {
		readyJob := readyJobs[i]
		if readyJob.RetryCount == minRetryCount {
			_ = jr.CancelJobByID(readyJob.ID)
			continue
		}
		acceptableJobs = append(acceptableJobs, readyJob)
	}

	return
}
