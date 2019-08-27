package example

import (
	db "github.com/bonoken/watcher-common/common/gorm"
	"github.com/bonoken/watcher-common/common/gormzap"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Orm        *gorm.DB
	connection string
	Logger     *zap.SugaredLogger
)

func mainTest() {

	// init GORM
	Orm = db.ConnectGORM("mysql", connection)
	Orm.LogMode(true)
	Orm.SetLogger(gormzap.New(Logger.Desugar(), gormzap.WithLevel(zapcore.DebugLevel)))
	defer Orm.Close()

}
