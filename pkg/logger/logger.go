package logger

import (
	"data-export/pkg/app"
	"data-export/pkg/console"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Config struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var Logger *zap.Logger

func InitLogger(cfg *Config) {
	ws := getLogWriter(cfg)
	level := zap.DebugLevel
	if !app.IsDebug() {
		err := level.Set(cfg.Level)
		console.ExitIf(err)
	}
	core := zapcore.NewCore(getEncoder(), ws, level)
	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),              // one layer is encapsulated, and one layer is removed from the calling file
		zap.AddStacktrace(zap.ErrorLevel), // stacktrace will be displayed only when error occurs
	)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder

	if app.IsDebug() {
		// keyword highlight of terminal output
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// json encoder used in online environment
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(cfg *Config) zapcore.WriteSyncer {
	if app.IsDebug() {
		return zapcore.AddSync(os.Stdout)
	}

	logger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	return zapcore.AddSync(logger)
}

// custom date format
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}
