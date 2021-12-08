package workerv1

import (
	"fmt"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

func typeBJobHandle(ctx cornerstone.Context, job *domainv1.Job) {
	funcName := "typeBJobHandle"
	ctx.Set("job.id", job.ID)
	ctx.Set("cid", job.CID)
	defer recoverTypeBJobHandle(ctx, job)
	var err error
	switch job.Status {
	case domainv1.JobGeneralStatusReady.ToInt():
		err = typeBJobHandleReadyToStarted(ctx, job)
	case domainv1.CustomStatusBStarted.ToInt():
		err = typeBJobHandleStartedToFinished(ctx, job)
	default:
		cornerstone.Errorf(ctx, "[%s] unsupport job status: %d, id: %s", funcName, job.Status, job.ID)
	}
	removeFromJobQueue(ctx, job, err)
}

func recoverTypeBJobHandle(ctx cornerstone.Context, job *domainv1.Job) {
	if err := recover(); err != nil {
		errPanic, _ := err.(error)
		removeFromJobQueue(ctx, job, errPanic)
	}
}

func typeBJobHandleReadyToStarted(ctx cornerstone.Context, job *domainv1.Job) (err error) {
	funcName := "typeBJobHandleReadyToStarted"

	// status
	if job.Status != domainv1.JobGeneralStatusReady.ToInt() {
		cornerstone.Infof(ctx, "[%s] invalid job status: %d", funcName, job.Status)
		err = fmt.Errorf("invalid job status")
		return
	}

	cornerstone.Debugf(ctx, "[%s] job type b: %s, from ready to started", funcName, job.ID)

	// update status
	job.Status = domainv1.CustomStatusBStarted.ToInt()
	err = jr.UpdateJobStatusStillOngoing(job)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to update job status to downloaded, err: %+v", funcName, err)
	}
	return
}

func typeBJobHandleStartedToFinished(ctx cornerstone.Context, job *domainv1.Job) (err error) {
	funcName := "typeBJobHandleStartedToFinished"

	// status
	if job.Status != domainv1.CustomStatusBStarted.ToInt() {
		cornerstone.Infof(ctx, "[%s] invalid job status: %d", funcName, job.Status)
		err = fmt.Errorf("invalid job status")
		return
	}

	cornerstone.Debugf(ctx, "[%s] job type b: %s, from started to finished", funcName, job.ID)

	// update status
	job.Status = domainv1.JobGeneralStatusFinished.ToInt()
	err = jr.UpdateJobStatusToStopped(job)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to update job status to finished, err: %+v", funcName, err)
	}
	return
}
