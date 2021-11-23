package workerv1

import (
	"runtime/debug"
	"sync"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

const (
	logWorkerPool       = "-WorkerPool"
	logJobWorker        = "-JobWorker"
	logJobWorkerPool    = "-JobWorkerLoop"
	logRecoverJobWorker = "-RecoverJobWorker"
)

func startJobWorkerLoop(ctx cornerstone.Context, jobType domainv1.JobType, maxJobWorkers int, jobHandle func(cornerstone.Context, *domainv1.Job), jobWorkerWaitGroup *sync.WaitGroup, jobWorkerQueue chan *domainv1.Job) {
	funcName := jobType.String() + logJobWorkerPool
	cornerstone.Infof(ctx, "[%s] start job workers: %d", funcName, maxJobWorkers)
	jobWorkerWaitGroup.Add(maxJobWorkers)
	for i := 0; i < maxJobWorkers; i++ {
		go jobWorker(ctx, jobType, jobHandle, jobWorkerWaitGroup, jobWorkerQueue)
	}
	cornerstone.Infof(ctx, "[%s] all job workers started", funcName)
}

func recoverJobWorker(ctx cornerstone.Context, jobType domainv1.JobType, jobHandle func(cornerstone.Context, *domainv1.Job), jobWorkerWaitGroup *sync.WaitGroup, jobWorkerQueue chan *domainv1.Job) {
	funcName := jobType.String() + logRecoverJobWorker
	if err := recover(); err != nil {
		cornerstone.Errorf(ctx, "[%s] error: %+v\n %s", funcName, err, debug.Stack())
		jobWorker(ctx, jobType, jobHandle, jobWorkerWaitGroup, jobWorkerQueue)
	}
}

func jobWorker(ctx cornerstone.Context, jobType domainv1.JobType, jobHandle func(cornerstone.Context, *domainv1.Job), jobWorkerWaitGroup *sync.WaitGroup, jobWorkerQueue chan *domainv1.Job) {
	funcName := jobType.String() + logJobWorker
	defer recoverJobWorker(ctx, jobType, jobHandle, jobWorkerWaitGroup, jobWorkerQueue)
	for task := range jobWorkerQueue {
		taskCtx := ctx.CopyContext()
		taskCtx.Set("name", jobType.String()+logWorkerPool)
		taskCtx.Set("cid", task.CID)
		jobHandle(taskCtx, task)
	}
	cornerstone.Infof(ctx, "[%s] worker finished", funcName)
	jobWorkerWaitGroup.Done()
}
