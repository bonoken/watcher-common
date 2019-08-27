package example

import (
	logger "github.com/bonoken/watcher-common/common/zap"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

func mainTest() {

	// init Zap Logger

	//isConsole := false
	//if global.Config.Project.Env == "local" {
	//	isConsole = true
	//}

	// init Zap Logger
	zapParam := logger.Param{
		LogPath:           "/logs",
		LogFile:           "test.log",
		LogLevel:          "debug",
		IsConsoleAppender: true,

		MaxSize:    100,
		MaxBackups: 30,
		MaxDays:    30,
	}
	Logger = logger.InitZap(&zapParam)
}
