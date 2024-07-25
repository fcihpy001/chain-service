// @Author fcihpy
// @Date 2024/7/17 18:12:00
// @Desc
//

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.SugaredLogger

func init() {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	syncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     28,
		Compress:   true,
		LocalTime:  true,
	})
	zapLogger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "line",
		MessageKey:     "msg",
		FunctionKey:    "F",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}),
		zapcore.NewMultiWriteSyncer(syncer, zapcore.AddSync(os.Stdout)), atomicLevel), zap.AddCaller())

	zap.ReplaceGlobals(zapLogger)
	log := zap.S()
	logger = log.WithOptions(zap.AddCallerSkip(1))
}

func Debug(msg string, kv ...interface{}) {
	logger.Debugw(msg, kv...)
}

func Info(msg string, kv ...interface{}) {
	logger.Infow(msg, kv...)
}

func Warn(msg string, kv ...interface{}) {
	logger.Warnw(msg, kv...)
}

func Error(msg string, kv ...interface{}) {
	logger.Errorw(msg, kv...)
}

func Panic(msg string, kv ...interface{}) {
	logger.Panic(msg, kv)
}

func PanicF(template string, kv ...interface{}) {
	logger.Panicf(template, kv)
}

func DebugF(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func InfoF(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func WarnF(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func ErrorF(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}
