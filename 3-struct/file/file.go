package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"work-with-bins/bins"
)

type FileStorage struct {
	name string
}

var fileName = ""

func Init(path string) *FileStorage {
	_, checkExistingError := os.Stat(path)

	if os.IsNotExist(checkExistingError) {
		createFile(path)
	}

	return &FileStorage{
		name: path,
	}
}

func (fileStorage FileStorage) SaveList(data *bins.BinList) bool {
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

func (fileStorage FileStorage) ReadList() ([]byte, error) {
	readingResult, readingError := os.ReadFile(fileName)

	if readingError != nil {
		fmt.Println("Error: file reading was interrupt")

		return nil, readingError
	}

	return readingResult, nil
}

func readFile(path string) ([]byte, error) {
	_, checkExistingError := os.Stat(path)

	if os.IsNotExist(checkExistingError) {
		return nil, errors.New("Error: file is not exist")
	}

	if checkExistingError != nil {
		return nil, errors.New("Error: checking file was finished with error")
	}

	file, readFileError := os.ReadFile(path)

	if readFileError != nil {
		return nil, readFileError
	}

	return file, nil
}

func createFile(name string) {
	fileName = name
	file, fileCreateError := os.Create(fileName)

	if fileCreateError != nil {
		fmt.Println("Error: during creation file error accured: ", fileCreateError)
	}

	defer file.Close()
}

func isJSONFile(path string) bool {
	extension := filepath.Ext(path)

	return extension == ".json"
}
