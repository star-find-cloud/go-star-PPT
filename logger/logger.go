package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"localhost/business_code_test/config"
	"sync"
)

type Logger struct {
	coreLogger *zap.SugaredLogger
	config     *config.Config
	model      string
	encoder    zapcore.Encoder
	fileSync   zapcore.WriteSyncer
	mu         sync.Mutex
}

var globalLogger *Logger

func init() {
	conf := config.GetConfig()
	encoderConig := zap.NewProductionEncoderConfig()
	encoderConig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConig.EncodeLevel = zapcore.CapitalLevelEncoder

	globalLogger = &Logger{
		config:  conf,
		encoder: zapcore.NewConsoleEncoder(encoderConig),
	}
}

func (l *Logger) rotateFile() {
	if l.fileSync != nil {
		_ = l.fileSync.Sync()
	}

	fileName := fmt.Sprintf("%s/app-%s-%s-%s.log", l.config.Log.Path, l.config.Log.Version, l.config.Log.Level, l.model)
	lumberJack := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    l.config.Log.MaxSize,
		MaxBackups: l.config.Log.MaxBackups,
		MaxAge:     l.config.Log.MaxAge,
		Compress:   l.config.Log.Compress,
	}
	l.fileSync = zapcore.AddSync(lumberJack)
	core := zapcore.NewCore(l.encoder, l.fileSync, zap.DebugLevel)

	if l.coreLogger != nil {
		_ = l.coreLogger.Sync()
	}
	l.coreLogger = zap.New(core).Sugar()
}

func (l *Logger) spawn(modelName string) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	newLogger := &Logger{
		config:  l.config,
		model:   modelName,
		encoder: l.encoder,
	}
	newLogger.rotateFile()
	return newLogger
}

func GetLog(model string) *zap.SugaredLogger {
	logger := globalLogger.spawn(model)
	return logger.coreLogger
}
