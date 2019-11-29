package system

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"errors"
	"strings"

)
var root string

func init(){
	root = "/home2/nginx"
}

func SysFileGetList()([]string, error){
	logger.Debug("Sys get list file")
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	for _, file := range files {
		logger.Debug(file)
	}
	return files, nil
}

func SysCheckFile(file_path string) error {
	info, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return errors.New("File not existed")
	}
	if !strings.HasPrefix(file_path, root){
		return errors.New("File not allow to download")
	}
	if info.IsDir(){
		return errors.New("Folder not allow to download")
	}
	return nil
}
