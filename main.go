package main

import (
	oroServer "github.com/ntwarijoshua/orostreams/http"
	"github.com/ntwarijoshua/orostreams/utils"
)

var (
	err error
)

func main() {
	//databaseSetup
	if utils.ActiveConnection, err = utils.NewDatabaseConnector("mysql", "root:root@/orostreams?charset=utf8&parseTime=True&loc=Local"); err != nil {
		panic(err)
	}
	utils.RunMigrations()
	oroServer.Start(":8083")
}
