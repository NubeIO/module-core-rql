package storage

import (
	"github.com/sonyarouje/simdb"
)

type db struct {
	DB *simdb.Driver
}

func New(dbFile string) (IStorage, error) {
	newDb, err := simdb.New(dbFile)
	return &db{DB: newDb}, err
}
