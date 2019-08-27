package main // import github.com/bonoken/watcher-common

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/spf13/viper"

	_ "github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/gommon/color"
	_ "github.com/valyala/fasttemplate"

	_ "go.uber.org/zap"
	_ "go.uber.org/zap/zapcore"

	_ "gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	// read config.yml

	// init Zap Logger

	// init GORM

	// echo

}
