```yaml
# 日志配置
Log:
  # Default：日志方式。 console:控制台日志打印，file：文件日志存储，kafka：写入kafka队列，需要配置elk相关组件
  # 支持多个同时存在，多个以,隔开：console,file,kafka (配置file和kafka时需要配置参数信息)
  Default: "console"
  Level: "info" #  debug info warn error 
  Encoder: "json" # plain json
  File:
    FilePath: "storage/logs/logs.log"
    MaxSize: 64
    MaxBackup: 10
    MaxAge: 30
    Compress: false
  Kafka:
    Host:
      - "127.0.0.1:9092"
    topic: "my-topic"
```