package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"work-with-bins/bins"
)

var fileName = ""

type Storage interface {
	InitStorage()
	SaveList(*bins.BinList)
	ReadList() ([]byte, error)
}

type FileStorage struct {
	fileName string
}

func InitStorage(name string) {
	fileName = name
	_, checkExistingError := os.Stat(fileName)

	if os.IsNotExist(checkExistingError) {
		createFile()
	}
}

func createFile() {
	file, fileCreateError := os.Create(fileName)

	if fileCreateError != nil {
		fmt.Println("Error: during creation file error accured: ", fileCreateError)
	}

	defer file.Close()
}

func SaveList(data *bins.BinList) bool {
	convertedData, convertingError := json.Marshal((*data).Bins)

	if convertingError != nil {
		fmt.Println("Error: convertation to JSON was interrupt")

		return false
	}

	fmt.Println(convertedData)

	recordError := os.WriteFile(fileName, convertedData, 0644)

	if recordError != nil {
		fmt.Println("Error: recording to file was interrupt")

		return false
	}

	return true
}

func ReadList() ([]byte, error) {
	readingResult, readingError := os.ReadFile(fileName)

	if readingError != nil {
		fmt.Println("Error: file reading was interrupt")

		return nil, readingError
	}

	return readingResult, nil
}
