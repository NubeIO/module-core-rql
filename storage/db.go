package storage

import (
	simdb "github.com/sonyarouje/simdb"
	"os"
	"path/filepath"
)

type db struct {
	DB *simdb.Driver
}

func New(dbFile string) (IStorage, error) {
	if dbFile == "" {
		dbFile = "data/data"
	}
	parentDir := filepath.Dir(dbFile)
	if parentDir != "" {
		err := os.MkdirAll(parentDir, 0755)
		if err != nil {
			panic("data directory creation issue")
		}
	}
	newDb, err := simdb.New(dbFile)

	return &db{DB: newDb}, err
}

//func New(dbFile string) (IStorage, error) {
//	if dbFile == "" {
//		dbFile = "data/data.db"
//	}
//	parentDir := filepath.Dir(dbFile)
//	if parentDir != "" {
//		err := os.MkdirAll(parentDir, 0755)
//		if err != nil {
//			panic("data directory creation issue")
//		}
//	}
//	newDb, err := buntdb.Open(dbFile)
//	if err != nil {
//		return &db{DB: newDb}, err
//	}
//	size := 10 * 1024 * 1024 //10mb
//	c := buntdb.Config{
//		SyncPolicy:           buntdb.EverySecond,
//		AutoShrinkPercentage: 30,
//		AutoShrinkMinSize:    size,
//	}
//	err = newDb.SetConfig(c)
//	if err != nil {
//		return &db{DB: newDb}, err
//	}
//	return &db{DB: newDb}, err
//}

//func (inst *db) Close() error {
//	//return inst.DB.Close()
//}

func matchRuleUUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "rql_" {
			return true
		}
	}
	return false
}

func matchVarUID(uuid string) bool {
	if len(uuid) == 16 {
		if uuid[0:4] == "var_" {
			return true
		}
	}
	return false
}
