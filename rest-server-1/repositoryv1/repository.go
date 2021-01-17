package repositoryv1

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lovemew67/project-misc/rest-server-1/gen/proto"
)

func CreateStaff(staff *proto.Staff) (result *proto.Staff, err error) {
	id := uuid.New().String()
	staff.Id = id
	now := time.Now().UnixNano()
	staff.Created = now
	staff.Updated = now
	db := sqlitedb
	db = db.Create(staff)
	err = db.Error
	if err != nil {
		return
	}

	result, err = GetStaff(id)
	return
}

func CountTotalStaff() (result int, err error) {
	db := sqlitedb
	db = db.Model(proto.Staff{})
	db = db.Count(&result)
	err = db.Error
	return
}

func GetStaff(id string) (staff *proto.Staff, err error) {
	staffList := make([]*proto.Staff, 1)
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

func QueryAllStaffWithOffsetAndLimit(offset, limit int) (staffList []proto.Staff, err error) {
	staffList = make([]proto.Staff, limit)
	db := sqlitedb
	db = db.Offset(offset)
	db = db.Limit(limit)
	db = db.Find(&staffList)
	err = db.Error
	return
}

func PatchStaff(id string, staff *proto.Staff) (err error) {
	now := time.Now().UnixNano()
	staff.Updated = now
	db := sqlitedb
	db = db.Model(proto.Staff{Id: id})
	db = db.Update(staff)
	err = db.Error
	return
}

func DeleteStaff(id string) (err error) {
	staff := proto.Staff{Id: id}
	db := sqlitedb
	db = db.Delete(proto.Staff{}, &staff)
	err = db.Error
	return
}
