package services

import "AICS_WebBackend/system"

// LoginUser login user service
func SysFileGetList() (result []string, err error) {
	result, err = system.SysFileGetList()
	return result, err
}

func SysCheckFile(file_path string) error {
	return system.SysCheckFile(file_path)
}
