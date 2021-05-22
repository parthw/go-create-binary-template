package logger

import (
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log to use application log configuration
var Log *zap.SugaredLogger

// InitializeLogger to initialize the application logger
func InitializeLogger() {
	logLevel := zapcore.InfoLevel
	if strings.ToLower(viper.GetString("log.level")) == "debug" {
		logLevel = zapcore.DebugLevel
	}
	core := zapcore.NewTee(zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), logLevel))
	Log = zap.New(core, zap.AddCaller()).Sugar()
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
