package gencontext

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("15:04:05.000") + "]")
}

func createLogDir(logDir string) error {
	err := os.MkdirAll(logDir, 0744)
	if err != nil {
		return err
	}
	return nil
}

func newConsoleEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "",
		MessageKey:     "msg",
		StacktraceKey:  "",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     SyslogTimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}
}

func newFileEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     SyslogTimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func initLogger(logDir string) *zap.Logger {
	err := createLogDir(logDir)
	if err != nil {
		panic(err)
	}

	consoleLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})
	fileLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.DebugLevel
	})

	consoleEncoder := zapcore.NewConsoleEncoder(newConsoleEncoderConfig())
	fileEncoder := zapcore.NewConsoleEncoder(newFileEncoderConfig())

	logFile, err := os.OpenFile(filepath.Join(logDir, "mify.log"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	cores := []zapcore.Core{
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stderr), consoleLevelEnabler),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), fileLevelEnabler),
	}

	core := zapcore.NewTee(cores...)
	return zap.New(core,
		zap.AddStacktrace(zap.ErrorLevel),
		zap.WithCaller(true),
	)
}
