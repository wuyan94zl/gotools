package logz

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type File struct {
	FilePath  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
}

func getFileWriter(c File) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.FilePath,  // 日志文件路径： storage/logs/logs.log
		MaxSize:    c.MaxSize,   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: c.MaxBackup, // 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
		MaxAge:     c.MaxAge,    // 最多保存多少天，7 表示一周前的日志会被删除，0 表示不删
		Compress:   c.Compress,  // 是否压缩，压缩日志不方便查看，我们设置为 false（压缩可节省空间）
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
