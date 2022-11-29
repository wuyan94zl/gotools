package logz

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
)

type Kafka struct {
	Host  []string
	Topic string
}

type KafkaLogWriter struct {
	Pusher *kq.Pusher
}

func (w *KafkaLogWriter) Write(p []byte) (n int, err error) {
	data := make(map[string]interface{})
	json.Unmarshal(p, &data)
	data["@timestamp"] = time.Now()
	marshal, _ := json.Marshal(data)
	if err := w.Pusher.Push(strings.TrimSpace(string(marshal))); err != nil {
		return 0, err
	}
	return len(p), nil
}
func (w *KafkaLogWriter) Sync() error {
	return nil
}

func getKafkaWriter(c Kafka) zapcore.WriteSyncer {
	pusher := kq.NewPusher(c.Host, c.Topic)
	writer := &KafkaLogWriter{
		Pusher: pusher,
	}
	return zapcore.AddSync(writer)
}
