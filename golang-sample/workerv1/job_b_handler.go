package workerv1

import (
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
	default:
		cornerstone.Errorf(ctx, "[%s] unsupport job status: %d, id: %s", funcName, job.Status, job.ID)
	}
	if err != nil {
		removeFromJobQueue(ctx, job, err)
	}
}

func recoverTypeBJobHandle(ctx cornerstone.Context, job *domainv1.Job) {
	if err := recover(); err != nil {
		errPanic, _ := err.(error)
		removeFromJobQueue(ctx, job, errPanic)
	}
}
