package main

import (
	"gitee.com/svanrj/server/Databases"
	"gitee.com/svanrj/server/Router"
)


type UserInfo struct {
	ID uint
	Username string
	Password string
}
func main() {
	Databases.Mysql()
	defer Databases.DB.Close()
	Router.InitRouter()

}
