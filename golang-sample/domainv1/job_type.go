package domainv1

import (
	"strconv"
)

/*
	skip 0, because gorm cannot update field with ZERO value:
		Update update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
		WARNING when update with struct, GORM will not update fields that with zero value
*/
type JobType int

const (
	JobTypeA JobType = 1
	JobTypeB JobType = 2
)

func (jt JobType) String() string {
	return strconv.Itoa(int(jt))
}

func isValidTaskType(tString string) (result int, valid bool) {
	tInt, err := strconv.Atoi(tString)
	if err != nil {
		return
	}
	result = tInt
	valid = tInt >= int(JobTypeA) && tInt <= int(JobTypeB)
	return
}

func ConvertToTaskType(tString string) (jt JobType, valid bool) {
	tInt, valid := isValidTaskType(tString)
	if !valid {
		return
	}
	jt = JobType(tInt)
	return
}
