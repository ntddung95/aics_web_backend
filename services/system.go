package services

import (
	"AICS_WebBackend/system"
	"mime/multipart"
)

// LoginUser login user service
func SysFileGetList() (result []string, err error) {
	result, err = system.SysFileGetList()
	return result, err
}

func SysCheckFile(file_path string) error {
	return system.SysCheckFile(file_path)
}

func SysCheckUploadPath(folder_path string) error {
	return system.SysCheckUploadPath(folder_path)
}

func SysFileUpload(folder_path string, file *multipart.FileHeader) error {
	return system.SysFileUpload(folder_path, file)
}
