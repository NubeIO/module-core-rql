package storage

import (
	"github.com/sonyarouje/simdb"
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
