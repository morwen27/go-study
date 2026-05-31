package api

import (
	"work-with-bins/bins"
	"work-with-bins/config"
)

type CloudStorage struct {
	url string
}

var appConfig config.Config

func Init(path string) *CloudStorage {
	appConfig = *config.InitConfig()

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
