package main

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(path string) string {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return ""
	}
	data, err := ioutil.ReadFile(path)
	CheckErr(err)

	return string(data)
}

func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	CheckErr(err)
}

func FindFolderInParent(path string, folder string) (string, error) {
	_, err := os.Stat(filepath.Join(path, folder))
	if os.IsNotExist(err) {
		if path == "/" {
			return path, errors.New("no git repository found")
		}
		parent := filepath.Dir(path)
		return FindFolderInParent(parent, folder)
	}
	return filepath.Join(path, folder), nil
}

func CreateFolderIfNotExists(path string, mode fs.FileMode) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Mkdir(path, mode)
	}
	return nil
}
