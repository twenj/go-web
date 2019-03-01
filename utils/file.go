package utils

import (
	"io/ioutil"
	"path/filepath"
)

func Traversing(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fileArr []string
	for _, file := range files {
		filename := file.Name()
		ext := filepath.Ext(filename)
		if ext == ".yaml" {
			fileArr = append(fileArr, filename)
		}
	}
	return fileArr, err
}
