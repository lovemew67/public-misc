package sqlite

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lovemew67/project-misc/grpc-gateway-1/gen/proto"
	"github.com/lovemew67/public-misc/cornerstone"
)

type StaffV1SQLiteRepositorier struct{}

func (s *StaffV1SQLiteRepositorier) CreateStaff(staff *proto.StaffV1) (result *proto.StaffV1, err error) {
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

	result, err = s.GetStaff(id)
	return
}

func (s *StaffV1SQLiteRepositorier) CountTotalStaff() (result int, err error) {
	db := sqlitedb
	db = db.Model(proto.StaffV1{})
	db = db.Count(&result)
	err = db.Error
	return
}

func (s *StaffV1SQLiteRepositorier) GetStaff(id string) (staff *proto.StaffV1, err error) {
	staffList := make([]*proto.StaffV1, 1)
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

func (s *StaffV1SQLiteRepositorier) QueryAllStaffWithOffsetAndLimit(offset, limit int) (staffList []*proto.StaffV1, err error) {
	staffList = make([]*proto.StaffV1, limit)
	db := sqlitedb
	db = db.Offset(offset)
	db = db.Limit(limit)
	db = db.Find(&staffList)
	err = db.Error
	return
}

func (s *StaffV1SQLiteRepositorier) PatchStaff(id string, staff *proto.StaffV1) (err error) {
	now := time.Now().UnixNano()
	staff.Updated = now
	db := sqlitedb
	db = db.Model(proto.StaffV1{Id: id})
	db = db.Update(staff)
	err = db.Error
	return
}

func (s *StaffV1SQLiteRepositorier) DeleteStaff(id string) (err error) {
	staff := proto.StaffV1{Id: id}
	db := sqlitedb
	db = db.Delete(proto.StaffV1{}, &staff)
	err = db.Error
	return
}

func NewStaffV1SQLiteRepositorier(ctx cornerstone.Context) (result *StaffV1SQLiteRepositorier, err error) {
	funcName := "NewStaffV1SQLiteRepositorier"

	dbFilePath := fmt.Sprintf(formatSqliteDatabasePath, dataFolder)
	if err = createDirIfNotExist(ctx, dataFolder); err != nil {
		return
	}
	if err = createFileIfNotExist(ctx, dbFilePath); err != nil {
		return
	}

	db, err := gorm.Open(dialect, dbFilePath)
	if err != nil {
		return
	}
	sqlitedb = db

	task := &proto.StaffV1{}
	if hasTable := sqlitedb.HasTable(task); hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, tableName)
		db.AutoMigrate(&proto.StaffV1{})
		return
	}

	if err = sqlitedb.CreateTable(task).Error; err != nil {
		return
	}

	result = &StaffV1SQLiteRepositorier{}
	return
}
