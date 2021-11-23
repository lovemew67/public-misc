package sqlite

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lovemew67/public-misc/cornerstone"
)

const (
	dialect = "sqlite3"
)

const (
	staffV1TableName = "staff_v1"
	jobV1TableName   = "job_v1"
)

const (
	dataFolder               = "./data"
	formatSqliteDatabasePath = "%s/db.sqlite"
)

var (
	sqlitedb *gorm.DB
)

func createDirIfNotExist(ctx cornerstone.Context, path string) (err error) {
	funcName := "createDirIfNotExist"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cornerstone.Debugf(ctx, "[%s] creating dir: %s", funcName, path)
		if err := os.MkdirAll(path, 0755); err != nil {
			cornerstone.Errorf(ctx, "[%s] failed to create dir for path: %s, err: %+v", funcName, path, err)
		}
	}
	return
}

func createFileIfNotExist(ctx cornerstone.Context, path string) (err error) {
	funcName := "createFileIfNotExist"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cornerstone.Debugf(ctx, "[%s] creating file: %s", funcName, path)
		if _, err := os.Create(path); err != nil {
			cornerstone.Errorf(ctx, "[%s] failed to create file for path: %s, err: %+v", funcName, path, err)
		}
	}
	return
}
