package domainv1

type jobProcessStatus int

const (
	jobProcessStatusStopped jobProcessStatus = -1
	jobProcessStatusOngoing jobProcessStatus = 1
)

func (jps jobProcessStatus) ToInt() int {
	return int(jps)
}
