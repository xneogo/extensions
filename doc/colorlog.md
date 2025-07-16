# 彩色日志工具 (colorlog)

## 简介

美观的彩色日志工具，支持：
- 彩色输出：不同级别的日志使用不同颜色
- 日志级别：Debug、Info、Warn、Error等级别
- 格式化输出：支持printf风格的格式化日志
- 自定义配置：可配置日志级别和颜色方案
- 上下文支持：支持context.Context的日志记录

## 安装

```bash
go get github.com/xneogo/extensions/colorlog
```

## 使用示例

### 基本用法

```go
import "github.com/xneogo/extensions/colorlog"

// 使用默认日志
ctx := context.Background()
colorlog.SetLevel(colorlog.InfoLevel)
colorlog.SetColorful(true)

colorlog.Debug(ctx, "调试信息")
colorlog.Info(ctx, "普通信息")
colorlog.Warn(ctx, "警告信息")
colorlog.Error(ctx, "错误信息")

// 格式化日志
colorlog.Infof(ctx, "用户 %s 登录成功", "张三")
```

### 自定义日志器

```go
// 创建自定义日志器
logger := colorlog.New()
logger.SetLevel(colorlog.WarnLevel)
logger.Warn(ctx, "这是警告日志")
```

## API 参考

### 日志级别

```go
const (
    DebugLevel LogLevel = iota
    InfoLevel
    WarnLevel
    ErrorLevel
)
```

### 日志接口

```go
type Logger interface {
    Debug(ctx context.Context, msg string)
    Info(ctx context.Context, msg string)
    Warn(ctx context.Context, msg string)
    Error(ctx context.Context, msg string)
    Debugf(ctx context.Context, format string, args ...interface{})
    Infof(ctx context.Context, format string, args ...interface{})
    Warnf(ctx context.Context, format string, args ...interface{})
    Errorf(ctx context.Context, format string, args ...interface{})
    SetLevel(level LogLevel)
    SetColorful(colorful bool)
}

func New() Logger
``` 