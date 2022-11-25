package logz

import (
	"github.com/zeromicro/go-queue/kq"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Kafka struct {
	Host  []string
	Topic string
}

type KafkaLogWriter struct {
	Pusher *kq.Pusher
}

func (w *KafkaLogWriter) Write(p []byte) (n int, err error) {
	if err := w.Pusher.Push(strings.TrimSpace(string(p))); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (w *KafkaLogWriter) Sync() error {
	return nil
}

func getKafkaWriter(c Kafka) zapcore.WriteSyncer {
	pusher := kq.NewPusher(c.Host, c.Topic)
	return &KafkaLogWriter{
		Pusher: pusher,
	}
}
