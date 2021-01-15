package repositoryv1

import (
	"fmt"

	"github.com/lovemew67/project-misc/rest-server-0/modelv1"
)

func CreateStaff(staff *modelv1.Staff) (err error) {
	db := sqlitedb
	db = db.Create(staff)
	err = db.Error
	return
}

func CountTotalStaff() (result int, err error) {
	db := sqlitedb
	db = db.Model(modelv1.Staff{})
	db = db.Count(&result)
	err = db.Error
	return
}

func GetStaff(id int) (staff *modelv1.Staff, err error) {
	staffList := make([]*modelv1.Staff, 1)
	db := sqlitedb
	db = db.Where("id = ?", id)
	db = db.Limit(1)
	db = db.Find(&staffList)
	err = db.Error
	if len(staffList) != 0 {
		staff = staffList[0]
	} else {
		err = fmt.Errorf("not found")
	}
	return
}

func QueryAllStaffWithOffsetAndLimit(offset, limit int) (staffList []modelv1.Staff, err error) {
	staffList = make([]modelv1.Staff, limit)
	db := sqlitedb
	db = db.Offset(offset)
	db = db.Limit(limit)
	db = db.Find(&staffList)
	err = db.Error
	return
}

func PatchStaff(id int, staff *modelv1.Staff) (err error) {
	db := sqlitedb
	db = db.Model(modelv1.Staff{ID: id})
	db = db.Update(staff)
	err = db.Error
	return
}

func DeleteStaff(id int) (err error) {
	staff := modelv1.Staff{ID: id}
	db := sqlitedb
	db = db.Delete(modelv1.Staff{}, &staff)
	err = db.Error
	return
}
