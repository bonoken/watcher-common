package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

// Param
type Param struct {
	LogPath           string
	LogFile           string
	LogLevel          string
	IsConsoleAppender bool

	MaxSize    int // megabytes
	MaxBackups int
	MaxDays    int //days
}

func InitZap(p *Param) *zap.SugaredLogger {

	// log level
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	switch {
	case p.LogLevel == "debug":
		level.SetLevel(zapcore.DebugLevel)
	case p.LogLevel == "info":
		level.SetLevel(zapcore.InfoLevel)
	case p.LogLevel == "warn":
		level.SetLevel(zapcore.WarnLevel)
	case p.LogLevel == "error":
		level.SetLevel(zapcore.ErrorLevel)
	case p.LogLevel == "fatal":
		level.SetLevel(zapcore.FatalLevel)
	}

	//
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Name",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "St",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logFile := p.LogFile
	if len(p.LogPath) > 0 {
		if strings.HasSuffix(p.LogPath, "/") {
			logFile = p.LogPath + p.LogFile
		} else {
			logFile = p.LogPath + "/" + p.LogFile
		}
	}

	maxSize := 100
	maxBackups := 30
	maxAge := 30

	if p.MaxSize > 0 {
		maxSize = p.MaxSize
	}

	if p.MaxBackups > 0 {
		maxBackups = p.MaxBackups
	}

	if p.MaxDays > 0 {
		maxAge = p.MaxDays
	}

	enc := zapcore.NewJSONEncoder(encoderConfig)
	sink := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    maxSize, // megabytes
			MaxBackups: maxBackups,
			MaxAge:     maxAge, //days
		},
	)

	// init zap
	logger, _ := zap.NewDevelopment()

	if p.IsConsoleAppender {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig = encoderConfig
		logger, _ = config.Build()
	} else {
		logger = zap.New(
			zapcore.NewCore(enc, sink, level),
		)
	}

	defer logger.Sync()
	//logger.Error("aaa",
	//	zap.String("eeef", "eefe"),
	//)

	sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", global.Config.Project.Url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", global.Config.Project.Url)

	return sugar
}
