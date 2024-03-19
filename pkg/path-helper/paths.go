package pathhelper

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func SanitizePath(path string) string {
	sPath := path

	usr, _ := user.Current()
	dir := usr.HomeDir

	if path == "~" || path == "$HOME" {
		sPath = dir
	} else if strings.HasPrefix(path, "~/") {
		sPath = filepath.Join(dir, path[2:])
	} else if strings.HasPrefix(path, "$HOME/") {
		sPath = filepath.Join(dir, path[5:])
	}
	return sPath
}

func FileExists(path string) bool {
	if info, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return !info.IsDir()
	}
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func PathIsFile(path string) (bool, error) {
	var err error
	var info os.FileInfo

	if info, err = os.Stat(path); err != nil {
		return false, err
	}
	if info.IsDir() {
		return false, nil
	}
	return true, nil
}

func CreatePathIfNotExist(path string) error {
	// check if we deal with a path or a filepath
	if len(strings.Split(GetFileName(path), ".")) > 1 {
		path = strings.Join(strings.Split(path, "/")[:len(strings.Split(path, "/"))-1], "/")
	}
	if PathExists(path) {
		return nil
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func RemovePath(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func GetFileNameWithoutExtension(path string) string {
	return strings.Split(GetFileName(path), ".")[0]
}

func GetFileName(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
}

func GetFilesInDir(path string) ([]string, error) {
	var files []string

	if !PathExists(path) {
		return files, fmt.Errorf("path '%s' does not exist", path)
	}

	// recursivly search for files in provided path
	err := filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if !dir.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return files, err
	}

	return files, nil
}
