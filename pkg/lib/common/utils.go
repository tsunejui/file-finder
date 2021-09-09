package common

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func GetFileInfo(path string) (fs.FileInfo, error) {
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get status: %v", err)
	}
	return fileinfo, nil
}

func GetFiles(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read the directory: %v", err)
	}
	return files, nil
}

func InArray(n string, array []string) bool {
	var inArray bool
	for _, v := range array {
		if v == n {
			inArray = true
			break
		}
	}
	return inArray
}
