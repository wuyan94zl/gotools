package logz

import (
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Default string
	Level   string
	Encoder string
	File    File
	Kafka   Kafka
}

func InitLog(c Config) {
	store := strings.Split(c.Default, ",")
	var writers []zapcore.WriteSyncer
	for _, v := range store {
		switch v {
		case "console":
			writers = append(writers, getConsoleWriter())
		case "file":
			writers = append(writers, getFileWriter(c.File))
		case "kafka":
			writers = append(writers, getKafkaWriter(c.Kafka))
		}
	}

	setWriter(zapcore.NewMultiWriteSyncer(writers...), c.Level, c.Encoder)
}

// 设置日志writers
func setWriter(writeSyncer zapcore.WriteSyncer, level string, encoder string) {
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		panic("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}
	// 初始化 core
	core := zapcore.NewCore(getEncoder(encoder), writeSyncer, logLevel)
	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)
	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

// 设置日志存储格式
func getEncoder(encoder string) zapcore.Encoder {
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "content",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}
	if encoder == "plain" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 终端输出的关键词高亮
		return zapcore.NewConsoleEncoder(encoderConfig)              // 本地设置内置的 Console 解码器（支持 stacktrace 换行）
	}
	return zapcore.NewJSONEncoder(encoderConfig) // 线上环境使用 JSON 编码器
}

// 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000Z"))
}
