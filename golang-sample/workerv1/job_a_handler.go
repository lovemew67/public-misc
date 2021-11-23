package workerv1

import (
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

func typeAJobHandle(ctx cornerstone.Context, job *domainv1.Job) {
	funcName := "typeAJobHandle"
	ctx.Set("job.id", job.ID)
	ctx.Set("cid", job.CID)
	switch job.Status {
	default:
		cornerstone.Errorf(ctx, "[%s] unsupport job status: %d, id: %d", funcName, job.Status, job.ID)
	}
}
