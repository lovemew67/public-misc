package domainv1

/*
	skip 0, because gorm cannot update field with ZERO value:
		Update update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
		WARNING when update with struct, GORM will not update fields that with zero value
*/
type JobGeneralStatus int

const (
	JobGeneralStatusReady     JobGeneralStatus = 1
	JobGeneralStatusFinished  JobGeneralStatus = 65535
	JobGeneralStatusCancelled JobGeneralStatus = -1
)

func (jgs JobGeneralStatus) ToInt() int {
	return int(jgs)
}
