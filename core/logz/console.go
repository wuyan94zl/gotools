package logz

import (
	"go.uber.org/zap/zapcore"
	"os"
)

func getConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
