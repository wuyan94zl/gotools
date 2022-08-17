集成第三方对象存储 oss、cos、七牛、本地

创建配置文件 yaml
```yaml
Store: local
MaxFileSize: 1073741824

OSS:
  BucketName: ""
  Endpoint: ""
  AccessId: ""
  AccessSecret: ""

Local:
  BaseDir: "upload/images"
  BaseUrl: "http://localhost:8888/static"

COS:
  AppId: ""
  BucketName: ""
  Region: ""
  SecretId: ""
  SecretKey: ""

QN:
  AccessKey: ""
  SecretKey: ""
  SiteUrl: ""
  BucketName: ""
```

使用
```go
// 支持 local、cos、oss 、qn（七牛）存储
// parse to config
type conf struct {
    Store       string
    MaxFileSize int64
    OSS         storage.OSSConfig
    Local       storage.LocalConfig
    COS         storage.COSConfig
    QN          storage.QNConfig
}
storage.Parse(conf.Store,conf.MaxFileSize,conf.OSS)
// conf.Store 与conf.OSS 对应

// form表单上传文件 r : *http.Request
storage.Put(l.r, "myFile", "test", "test1.png") // 指定文件名
storage.Put(l.r, "myFile", "test") // 随机生成文件名，文件扩展一致

// 本地文件路径上传
storage.PutFromFile(filePath, "test", "test2.png") // 指定文件名
storage.PutFromFile(filePath, "test") // 随机生成文件名，文件扩展一致

// 删除文件
storage.Delete("test/test1.png") // 相对路径

```
