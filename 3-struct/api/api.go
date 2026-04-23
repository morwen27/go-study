package api

import (
	"work-with-bins/bins"
)

type CloudStorage struct {
	url string
}

func Init(path string) *CloudStorage {
	return &CloudStorage{
		url: path,
	}
}

func (cloudStorage CloudStorage) SaveList(data *bins.BinList) bool {
	return true
}

func (cloudStorage CloudStorage) ReadList() ([]byte, error) {
	return nil, nil
}
