# 限流器工具 (xrate)

## 简介

多种限流算法实现，包含：
- 漏桶算法：平滑限流，适合突发流量控制
- 滑动窗口：精确限流，适合QPS控制
- 动态配置：支持运行时修改限流阈值
- 多维度限流：支持按不同维度的限流策略
- 性能优化：高性能的限流实现

## 安装

```bash
go get github.com/xneogo/extensions/xrate
```

## 使用示例

### 漏桶限流器

```go
import "github.com/xneogo/extensions/xrate"

// 漏桶限流器
limiter := xrate.NewLeakyBucketRateLimiter(100) // 每秒100个请求

// 检查是否被限流
err := limiter.Limit()
if err == xrate.ErrRateLimited {
    // 被限流，处理逻辑
    return
}
```

### 滑动窗口限流器

```go
// 滑动窗口限流器
windowLimiter := xrate.NewSlidingWindowRateLimiter(
    1000,       // QPS阈值
    10,         // 窗口大小
    100,        // 子窗口间隔(ms)
)

// 动态修改限流阈值
limiter.ChangeQpsThreshold(200)
```

## API 参考

### 限流器接口

```go
type RateLimiter interface {
    Limit() error
    ChangeQpsThreshold(qps int)
}

var ErrRateLimited = errors.New("rate limited")
```

### 具体实现

```go
func NewLeakyBucketRateLimiter(qps int) RateLimiter
func NewSlidingWindowRateLimiter(qps, windowSize, intervalMs int) RateLimiter
``` 