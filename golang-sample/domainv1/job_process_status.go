package domainv1

type JobProcessStatus int

const (
	JobProcessStatusStopped JobProcessStatus = -1
	JobProcessStatusOngoing JobProcessStatus = 1
)

func (jps JobProcessStatus) ToInt() int {
	return int(jps)
}
