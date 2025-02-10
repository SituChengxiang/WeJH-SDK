## WeJH-SDK

### 工具

- AES 加密解密
- MinIO 对象服务操作
- Zap 日志初始化
- Redis 初始化
- 基于 Redis 的 Session 初始化
- 基于 Redis 的 Wechat 初始化
- 邮件功能
- Fetch

### 使用

```
go get github.com/SituChengxiang/WeJH-SDK
```

### 代码规范检查

需要安装 [gci](https://github.com/daixiang0/gci) 和 [golangci-lint](https://golangci-lint.run/)

```
gofmt -w .
gci write . -s standard -s default
golangci-lint run --config .golangci.yml
```
