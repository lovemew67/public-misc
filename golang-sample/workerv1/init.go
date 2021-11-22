package workerv1

import (
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
	jobWorkerWaitGroup sync.WaitGroup
	jobWorkerTaskQueue chan *domainv1.Job
	jobWorkerDone      = make(chan struct{}, 1)
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
}
