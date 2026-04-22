package file

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
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

func IsJSONFile(path string) bool {
	extension := filepath.Ext(path)

	return extension == ".json"
}
