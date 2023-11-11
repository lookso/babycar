package zlog

import (
	"babycare/internal/common"
	"context"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ContextKey string

var (
	instanceOfStdLog    *zap.Logger
	instanceOfSugardLog *zap.SugaredLogger

	zapContextKey ContextKey = "zapLogger"
)

func Init(project string, env, filename string, maxSize, maxBackUps, maxAge int, compress bool) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.MessageKey = ""
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		getLogWriter(filename, maxSize, maxBackUps, maxAge, compress),
		zap.InfoLevel,
	)
	if env == common.EnvDev {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	}
	//core := zapcore.NewTee(consoleCore,fileCore)
	instanceOfStdLog = zap.New(core, zap.Fields(
		zap.String("project", project),
	),
		zap.AddCaller(),
		zap.AddCallerSkip(4),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	instanceOfSugardLog = instanceOfStdLog.Sugar()
}

func getLogWriter(filename string, maxSize, maxBackUps, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackUps,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Sync() error {
	return instanceOfStdLog.Sync()
}

func SugaredInstance() *zap.SugaredLogger {
	return instanceOfSugardLog
}

func STDInstance() *zap.Logger {
	return instanceOfStdLog
}

func WithGoContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return STDInstance()
	}
	if ctxLogger, ok := ctx.Value(zapContextKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return STDInstance()
}

func WithGoContextSugared(ctx context.Context) *zap.SugaredLogger {
	stdLog := WithGoContext(ctx)
	return stdLog.Sugar()
}
