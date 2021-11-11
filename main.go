package main

import (
	"gitee.com/svanrj/server/Databases"
	"gitee.com/svanrj/server/Router"
)

func main() {
	Databases.Mysql()
	defer Databases.DB.Close()
	Router.InitRouter()

}
