package workerv1

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
)

const (
	minDispatcherCheckInterval = 200 * time.Millisecond
	defaultJobWorkerSize       = 10
	defaultJobQueueSize        = 10
)

var (
	jobTypeAWorkerWaitGroup sync.WaitGroup
	jobTypeAWorkerQueue     chan *domainv1.Job
)

var (
	jobTypeBWorkerWaitGroup sync.WaitGroup
	jobTypeBWorkerQueue     chan *domainv1.Job
)

func InitAndBlocking(ctx cornerstone.Context, workerStop chan os.Signal) {
	funcName := "workerv1.InitAndBlocking"

	// validate config
	dispatcherCheckInterval := 10 * time.Second
	if dispatcherCheckInterval < minDispatcherCheckInterval {
		cornerstone.Infof(ctx, "[%s] dispatcher check interval less than: %d ns, will set to: %d ns", funcName, minDispatcherCheckInterval, minDispatcherCheckInterval)
		dispatcherCheckInterval = minDispatcherCheckInterval
	}

	// validate config
	jobWorkerSize := defaultJobWorkerSize
	jobQueueSize := defaultJobQueueSize
	if jobWorkerSize == 0 || jobQueueSize == 0 {
		cornerstone.Panicf(ctx, "[%s] empty app.chatroom_audio.task setting", funcName)
	}
	if jobWorkerSize > jobQueueSize {
		cornerstone.Panicf(ctx, "[%s] worker size should not less than job size", funcName)
	}

	// init task queue
	jobTypeAWorkerQueue = make(chan *domainv1.Job, jobQueueSize)
	jobTypeBWorkerQueue = make(chan *domainv1.Job, jobQueueSize)

	// start workers
	startJobWorkerLoop(
		ctx,
		domainv1.JobTypeA,
		jobWorkerSize,
		typeAJobHandle,
		&jobTypeAWorkerWaitGroup,
		jobTypeAWorkerQueue,
	)
	startJobWorkerLoop(
		ctx,
		domainv1.JobTypeB,
		jobWorkerSize,
		typeBJobHandle,
		&jobTypeBWorkerWaitGroup,
		jobTypeBWorkerQueue,
	)

	// start dispatcher
	startDispatcherLoop(
		ctx,
		jobDispatch,
		dispatcherCheckInterval,
	)

	// blocking
	sig := <-workerStop
	cornerstone.Infof(ctx, "[%s] receive exit signal: %+v", funcName, sig)

	// stop dispatcher
	close(dispatcherStop)
	<-dispatcherDone

	// stop workers
	close(jobTypeAWorkerQueue)
	close(jobTypeBWorkerQueue)
	jobTypeAWorkerWaitGroup.Wait()
	jobTypeBWorkerWaitGroup.Wait()
}

func updateInternalDataFailedReasons(ctx cornerstone.Context, job *domainv1.Job, message string) (err error) {
	funcName := "updateInternalDataFailedReasons"

	internalData := domainv1.JobInternalData{}
	if err = json.Unmarshal([]byte(job.InternalData), &internalData); err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to unmarshal interenal data: %+v, err: %+v", funcName, job.InternalData, err)
		return
	}
	internalData.FailedReasons[job.RetryCount] = message

	b, err := json.Marshal(internalData)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to marshal interenal data: %+v, err: %+v", funcName, internalData, err)
		return
	}
	job.InternalData = string(b)

	return
}

func removeFromTaskQueue(ctx cornerstone.Context, job *domainv1.Job, err error) {
	funcName := "removeFromTaskQueue"
	if err := updateInternalDataFailedReasons(ctx, job, err.Error()); err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to update internal data for job: %+v, err: %+v", funcName, job, err)
	}
	// if err := repositoryv2.RemoveFromTaskQueue(job); err != nil {
	// 	cornerstone.Errorf(ctx, "[%s] failed to remove job: %+v, from queue, err: %+v", funcName, job, err)
	// }
}
