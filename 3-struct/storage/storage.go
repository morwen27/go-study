package storage

import "work-with-bins/bins"

type Storage interface {
	SaveList(*bins.BinList) bool
	ReadList() ([]byte, error)
}

type StorageWithDb struct {
	Storage
	db *bins.BinList
}

func Init(storageType Storage) *StorageWithDb {
	return &StorageWithDb{
		Storage: storageType,
		db:      bins.ConstructBinList(),
	}
}
