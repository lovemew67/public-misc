package workerv1

import (
	"fmt"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

func typeAJobHandle(ctx cornerstone.Context, job *domainv1.Job) {
	funcName := "typeAJobHandle"
	ctx.Set("job.id", job.ID)
	ctx.Set("cid", job.CID)
	defer recoverTypeAJobHandle(ctx, job)
	var err error
	switch job.Status {
	case domainv1.JobGeneralStatusReady.ToInt():
		err = typeAJobHandleReadyToStarted(ctx, job)
	case domainv1.CustomStatusAStarted.ToInt():
		err = typeAJobHandleStartedToFinished(ctx, job)
	default:
		cornerstone.Errorf(ctx, "[%s] unsupport job status: %d, id: %s", funcName, job.Status, job.ID)
	}
	removeFromJobQueue(ctx, job, err)
}

func recoverTypeAJobHandle(ctx cornerstone.Context, job *domainv1.Job) {
	if err := recover(); err != nil {
		errPanic, _ := err.(error)
		removeFromJobQueue(ctx, job, errPanic)
	}
}

func typeAJobHandleReadyToStarted(ctx cornerstone.Context, job *domainv1.Job) (err error) {
	funcName := "typeAJobHandleReadyToStarted"

	// status
	if job.Status != domainv1.JobGeneralStatusReady.ToInt() {
		cornerstone.Infof(ctx, "[%s] invalid job status: %d", funcName, job.Status)
		err = fmt.Errorf("invalid job status")
		return
	}

	cornerstone.Debugf(ctx, "[%s] job type a: %s, from ready to started", funcName, job.ID)

	// update status
	job.Status = domainv1.CustomStatusAStarted.ToInt()
	err = jr.UpdateJobStatusStillOngoing(job)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to update job status to downloaded, err: %+v", funcName, err)
	}
	return
}

func typeAJobHandleStartedToFinished(ctx cornerstone.Context, job *domainv1.Job) (err error) {
	funcName := "typeAJobHandleStartedToFinished"

	// status
	if job.Status != domainv1.CustomStatusAStarted.ToInt() {
		cornerstone.Infof(ctx, "[%s] invalid job status: %d", funcName, job.Status)
		err = fmt.Errorf("invalid job status")
		return
	}

	cornerstone.Debugf(ctx, "[%s] job type a: %s, from started to finished", funcName, job.ID)

	// update status
	job.Status = domainv1.JobGeneralStatusFinished.ToInt()
	err = jr.UpdateJobStatusToStopped(job)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to update job status to finished, err: %+v", funcName, err)
	}
	return
}
