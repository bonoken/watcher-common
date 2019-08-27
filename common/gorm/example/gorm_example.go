package example

import (
	db "github.com/bonoken/watcher-common/common/gorm"
)

func mainTest() {

	// init GORM
	Orm := db.ConnectGORM("mysql", "username:password@tcp(**.**.**.***:3306)/database?charset=utf8&parseTime=True&loc=Local")
	//Orm.LogMode(true)
	defer Orm.Close()

}
