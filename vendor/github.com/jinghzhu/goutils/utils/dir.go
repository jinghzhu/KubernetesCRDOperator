package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CreateTempDir creates a new temporary directory in the directory dir with a name beginning with prefix
// and returns the path of the new directory. If dir doesn't exist, CreateTempDir will create it at first.
// It is the caller's responsibility to remove the directory when no longer needed.
func CreateTempDir(dirName, dirPrefix string) (string, error) {
	flag, err := HasDir(dirName)
	if err != nil {
		return "", err
	}
	if !flag {
		if err = CreateDir(dirName); err != nil {
			return "", err
		}
	}

	return ioutil.TempDir(dirName, dirPrefix)
}

// CreateDir creates a directory named path, along with any necessary parents, and returns nil, or else
// returns an error. The permission bits perm (before umask) are used for all directories that MkdirAll
// creates. If path is already a directory, MkdirAll does nothing and returns nil.
func CreateDir(dirName string) error {
	return os.MkdirAll(dirName, os.ModePerm)
}

// HasDir checks wheter the directory exists or not.
func HasDir(dirPath string) (bool, error) {
	_, err := os.Stat(dirPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, fmt.Errorf("Fail to check dir %s because of %v", dirPath, err)
}

// DeleteDir removes path and any children it contains. It removes everything it can but returns the
// first error it encounters. If the path does not exist, DeleteDir returns nil (no error).
func DeleteDir(dirName string) error {
	return os.RemoveAll(dirName)
}
