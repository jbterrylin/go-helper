package zaphelper

import (
	"os"

	filehelper "github.com/jbterrylin/go-helper/fileHelper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// format: "json" | ""
// encodeLevel: "LowercaseLevelEncoder" |  "LowercaseColorLevelEncoder" | "CapitalLevelEncoder" |  "CapitalColorLevelEncoder" | ""
// level: "debug" | "info" | "warn" | "error" | "dpanic" | "panic" | "fatal"
func NewZap(director, format, stacktraceKey, encodeLevel, prefix, level string, logInConsole, showLine bool) (logger *zap.Logger) {
	if ok, _ := filehelper.DirExist(director); !ok {
		_ = os.Mkdir(director, os.ModePerm)
	}

	zapCore := NewZapCore(director, format, stacktraceKey, encodeLevel, prefix, level, logInConsole)
	cores := zapCore.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if showLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
