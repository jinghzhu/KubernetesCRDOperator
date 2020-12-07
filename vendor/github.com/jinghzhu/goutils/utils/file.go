package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// CreateFile creates the file at the given path. It some directory doesn't exist in the path,
// it will create them at first.
func CreateFile(file, data string) error {
	dirName, fileName := path.Split(file)
	_, err := CreateFileWithDirFile(dirName, fileName, data)

	return err
}

// CreateTempFile creates a new temporary file in the directory dir with a name beginning with prefix,
// opens the file for reading and writing, and returns the resulting *os.File. The caller can use f.Name()
// to find the pathname of the file. It is the caller's responsibility to remove the file when no longer needed.
// If some directory doesn't exist in the path, it will create them at first.
func CreateTempFile(dirName, filePrefix string) (*os.File, error) {
	flag, err := HasDir(dirName)
	if err != nil {
		return nil, err
	}
	if !flag {
		if err = CreateDir(dirName); err != nil {
			return nil, err
		}
	}

	return ioutil.TempFile(dirName, filePrefix)
}

// CreateFileWithDirFile creates the file with given directory and file name.
func CreateFileWithDirFile(dirName, fileName, data string) (string, error) {
	filePath := path.Join(dirName, fileName)
	exist, err := HasDir(dirName)
	if err != nil {
		return "", err
	}
	if !exist {
		if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
			return "", fmt.Errorf("Fail to create file %s because of %v", filePath, err)
		}
	}
	exist, err = HasFile(filePath)
	if err != nil {
		return "", err
	}
	if exist {
		return "", fmt.Errorf("Fail to create file %s because it already exists", filePath)
	}
	if _, err = os.Create(filePath); err != nil {
		return "", err
	}
	if err = ioutil.WriteFile(filePath, []byte(data), 0777); err != nil {
		return "", fmt.Errorf("Fail to write data into file %s because of %v", filePath, err)
	}

	return filePath, nil
}

// HasFile checks whether the file exists or not.
func HasFile(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// HasFileWithDirFile checks whether the file exists or not.
func HasFileWithDirFile(dirName, fileName string) (bool, error) {
	filePath := path.Join(dirName, fileName)
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// DeleteFile removes the named file or directory. If there is an error, it will be of type *PathError.
// No error is thrown if we can't find the file.
func DeleteFile(file string) error {
	_, err := Retry(800*time.Millisecond, 3, func() (bool, error) { // Retry to delete local kubeconfig file
		err1 := os.Remove(file)
		if err1 == nil || os.IsNotExist(err1) {
			return true, nil
		}

		return false, err1
	})

	return err
}

// DeleteFileWithDirFile removes files with given path name and file name.
func DeleteFileWithDirFile(dirName, fileName string) error {
	return DeleteFile(path.Join(dirName, fileName))
}
