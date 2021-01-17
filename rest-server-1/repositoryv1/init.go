package repositoryv1

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/rest-server-1/gen/proto"
)

const (
	tableName = "staff"
)

var (
	sqlitedb *gorm.DB
)

func Init(ctx cornerstone.Context) {
	funcName := "repositoryv2.Init"

	dbFilePath := fmt.Sprintf(formatSqliteDatabasePath, dataFolder)
	if err := createDirIfNotExist(ctx, dataFolder); err != nil {
		panic(err)
	}
	if err := createFileIfNotExist(ctx, dbFilePath); err != nil {
		panic(err)
	}

	db, errOpen := gorm.Open(dialect, dbFilePath)
	if errOpen != nil {
		cornerstone.Panicf(ctx, "[%s] failed to connect: %s, database: %s, err: %+v", funcName, dialect, dbFilePath, errOpen)
	}
	sqlitedb = db

	task := &proto.Staff{}
	if hasTable := sqlitedb.HasTable(task); hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, tableName)
		db.AutoMigrate(&proto.Staff{})
		return
	}

	if errCreate := sqlitedb.CreateTable(task).Error; errCreate != nil {
		cornerstone.Panicf(ctx, "[%s] failed to create table: %s, err: %+v", funcName, tableName, errCreate)
	}
}

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
