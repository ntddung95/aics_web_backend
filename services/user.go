package services

import "AICS_WebBackend/database"

// LoginUser login user service
func LoginUser(user, password string) (success bool, permission string) {
	success, permission = database.LoginUser(user, password)
	return success, permission
}
