package system

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"errors"
	"mime/multipart"
	"io"

)
var root string

func init(){
	root = "/home2/nginx"
}

func SysFileGetList()([]string, error){
	logger.Debug("Sys get list file")
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root{
			return nil
		}
		path = path[len(root)+1:len(path)]
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
	info, err := os.Stat(root + "/" + file_path)
	if os.IsNotExist(err) {
		return errors.New("File not existed")
	}
	if info.IsDir(){
		return errors.New("Folder not allow to download")
	}
	return nil
}


func SysCheckUploadPath(folder_path string) error {
	info, err := os.Stat(root + "/" + folder_path)
	if os.IsNotExist(err) {
		return errors.New("Folder not available")
	}
	if !info.IsDir(){
		return errors.New("Folder path upload is a file")
	}
	return nil
}

func SysFileUpload(folder_path string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(root + "/" + folder_path + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil{
		return err
	}
	return nil
}
