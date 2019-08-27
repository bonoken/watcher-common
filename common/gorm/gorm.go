package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectGORM(dialect string, connection string) *gorm.DB {
	//fmt.Print(dialect +", " +connection)
	// Connect
	db, err := gorm.Open(dialect, connection)
	if err != nil {
		panic(err)
	}
	return db
}
