package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Log to use application log configuration
var Log *zap.SugaredLogger

// InitializeLogger to initialize the application logger
func InitializeLogger() {
	logLevel := zapcore.InfoLevel
	if strings.ToLower(viper.GetString("log.level")) == "debug" {
		logLevel = zapcore.DebugLevel
	}

	// Setting log type as console
	core := zapcore.NewTee(zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), logLevel))
	if strings.ToLower(viper.GetString("log.type")) == "json" {
		core = zapcore.NewTee(zapcore.NewCore(getJSONEncoder(), zapcore.AddSync(os.Stdout), logLevel))
	} else if strings.ToLower(viper.GetString("log.type")) == "file" {
		core = zapcore.NewTee(zapcore.NewCore(getJSONEncoder(), getLogWriter(), logLevel))
	} else if strings.ToLower(viper.GetString("log.type")) == "console" {
		// Using default logger
	} else {
		fmt.Println("Using default console logging.")
	}
	Log = zap.New(core, zap.AddCaller()).Sugar()
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	logFile := viper.GetString("log.file")
	logFileMaxSize := viper.GetInt("log.file.maxsize")
	logFileMaxBackups := viper.GetInt("log.file.maxbackups")
	logFileMaxAge := viper.GetInt("log.file.maxage")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    logFileMaxSize, //megabytes
		MaxBackups: logFileMaxBackups,
		MaxAge:     logFileMaxAge, //days
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
