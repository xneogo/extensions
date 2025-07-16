# 熔断器工具 (xbreaker)

## 简介

生产级的熔断器实现，提供：
- 熔断策略：基于错误率和QPS的熔断策略
- 状态管理：关闭、开启、半开三种状态
- 统计功能：实时统计请求成功率和响应时间
- 自动恢复：支持自动恢复和手动恢复
- 配置灵活：支持动态配置和热更新

## 安装

```bash
go get github.com/xneogo/extensions/xbreaker
```

## 使用示例

### 基本用法

```go
import "github.com/xneogo/extensions/xbreaker"

// 创建熔断器配置
config := xbreaker.CircuitBreakerConfig{
    MinQPS:               10,
    FailureRateThreshold: 50, // 50%错误率触发熔断
    OpenStatusDurationMs: 5000, // 5秒后尝试半开
}

// 创建熔断器
breaker := xbreaker.NewCircuitBreaker(config)

// 使用熔断器执行函数
err := breaker.Execute(ctx, func() error {
    // 业务逻辑
    return doSomething()
})
```

### 简单计数器熔断器

```go
// 简单计数器熔断器
xbreaker.StatBreaker("cluster1", "table1", err) // 统计错误
allowed := xbreaker.Entry("cluster1", "table1") // 检查是否允许请求
```

## API 参考

### 配置结构

```go
type CircuitBreakerConfig struct {
    MinQPS               int
    FailureRateThreshold float64
    OpenStatusDurationMs int64
}
```

### 熔断器接口

```go
type CircuitBreaker interface {
    Execute(ctx context.Context, fn func() error) error
    IsOpen() bool
    GetStats() Stats
}

func NewCircuitBreaker(config CircuitBreakerConfig) CircuitBreaker
```

### 简单熔断器

```go
func StatBreaker(cluster, table string, err error)
func Entry(cluster, table string) bool
``` 