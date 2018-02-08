package storage

import (
	"io"
	uid "github.com/segmentio/ksuid"
	"os"
	"fmt"
)

const (
	filesRootDir = "files"
	filesRootDirPerm os.FileMode = 0666
	storedFilesPerm os.FileMode = 0666
)

type Item struct {
	Key string
	Url string
}

func Remove(key string) {
	os.Remove(getItemPath(key))
}

func Save(data io.Reader) (*Item, error) {
	err := ensureDirectoryExists()
	if err != nil {
		return nil, err
	}

	fileKey := uid.New().String()
	filePath := getItemPath(fileKey)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, storedFilesPerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, data)
	if err != nil {
		Remove(fileKey)
		return nil, err
	}

	return &Item{fileKey, filePath}, nil
}
func getItemPath(fileKey string) string {
	return fmt.Sprintf("%s/%s", filesRootDir, fileKey)
}

func ensureDirectoryExists() error {
	return os.MkdirAll(filesRootDir, filesRootDirPerm)
}