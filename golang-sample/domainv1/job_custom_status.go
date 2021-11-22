package domainv1

// different job type need to define their status
// and start from 2, end with TaskGeneralStatusFinished (65535)

type JobCustomStatus int

const (
	CustomStatusAStarted JobCustomStatus = 2
	CustomStatusAEnded   JobCustomStatus = 3
)

const (
	CustomStatusBStarted JobCustomStatus = 2
	CustomStatusBEnded   JobCustomStatus = 3
)

func (jcs JobCustomStatus) ToInt() int {
	return int(jcs)
}
