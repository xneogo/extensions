# 时间工具 (xtime)

## 简介

全面的时间处理工具，提供：
- 时间范围：获取日、周、月、年的开始和结束时间
- 时间戳操作：Unix时间戳的各种转换和计算
- 时间格式化：多种预定义格式和自定义格式
- 定时器：可配置的定时器和运行时间统计
- 时间范围查询：支持绝对时间和相对时间范围

## 安装

```bash
go get github.com/xneogo/extensions/xtime
```

## 使用示例

### 时间范围获取

```go
import "github.com/xneogo/extensions/xtime"

// 获取时间范围
now := time.Now()
startOfDay := xtime.StartOfDay(now)
endOfDay := xtime.EndOfDay(now)
startOfWeek := xtime.StartOfWeek(now)
endOfWeek := xtime.EndOfWeek(now)
startOfMonth := xtime.StartOfMonth(now)
endOfMonth := xtime.EndOfMonth(now)

fmt.Printf("今天开始: %v\n", startOfDay)
fmt.Printf("今天结束: %v\n", endOfDay)
fmt.Printf("本周开始: %v\n", startOfWeek)
fmt.Printf("本周结束: %v\n", endOfWeek)
fmt.Printf("本月开始: %v\n", startOfMonth)
fmt.Printf("本月结束: %v\n", endOfMonth)
```

### 时间戳操作

```go
// 时间戳操作
nowStamp := time.Now().Unix()
dayBegin := xtime.DayBeginStamp(nowStamp)
dayEnd := xtime.DayEndStamp(nowStamp)
hourBegin := xtime.HourBeginStamp(nowStamp)
hourEnd := xtime.HourEndStamp(nowStamp)

fmt.Printf("当前时间戳: %d\n", nowStamp)
fmt.Printf("今天开始时间戳: %d\n", dayBegin)
fmt.Printf("今天结束时间戳: %d\n", dayEnd)
fmt.Printf("当前小时开始: %d\n", hourBegin)
fmt.Printf("当前小时结束: %d\n", hourEnd)

// 时间范围
dayStart, dayEnd := xtime.DayScope(nowStamp)
weekStart, weekEnd := xtime.WeekScope(nowStamp)
monthStart, monthEnd := xtime.MonthScope(nowStamp)

fmt.Printf("日范围: %d - %d\n", dayStart, dayEnd)
fmt.Printf("周范围: %d - %d\n", weekStart, weekEnd)
fmt.Printf("月范围: %d - %d\n", monthStart, monthEnd)
```

### 时间格式化

```go
// 预定义格式常量
now := time.Now()

// 使用预定义格式
dateStr := now.Format(xtime.DateFormat)      // 2006-01-02
timeStr := now.Format(xtime.TimeFormat)      // 15:04:05
datetimeStr := now.Format(xtime.DateTimeFormat) // 2006-01-02 15:04:05

fmt.Printf("日期: %s\n", dateStr)
fmt.Printf("时间: %s\n", timeStr)
fmt.Printf("日期时间: %s\n", datetimeStr)

// 解析时间字符串
parsedTime, err := time.Parse(xtime.DateTimeFormat, "2023-12-25 15:30:45")
if err != nil {
    fmt.Printf("解析失败: %v\n", err)
} else {
    fmt.Printf("解析时间: %v\n", parsedTime)
}
```

### 定时器

```go
// 创建定时器
timer := xtime.NewTimer(2 * time.Second)

// 启动定时器
timer.Start(func() {
    fmt.Println("定时器触发")
})

// 等待一段时间后停止
time.Sleep(10 * time.Second)
timer.Stop()
fmt.Println("定时器已停止")

// 重置定时器间隔
timer.Reset(5 * time.Second)
timer.Start(func() {
    fmt.Println("重置后的定时器触发")
})
```

### 运行时间统计

```go
// 创建时间统计器
timeStat := xtime.NewTimeStat()

// 执行业务逻辑
time.Sleep(100 * time.Millisecond)
fmt.Println("执行了一些业务逻辑...")

// 获取运行时间
nanoseconds := timeStat.Nanosecond()     // 纳秒
microseconds := timeStat.Microsecond()   // 微秒  
milliseconds := timeStat.Millisecond()   // 毫秒
seconds := timeStat.Second()             // 秒

fmt.Printf("运行时间 - 纳秒: %d, 微秒: %d, 毫秒: %d, 秒: %.2f\n",
    nanoseconds, microseconds, milliseconds, seconds)
```

### 时间窗口

```go
// 创建时间窗口
window := xtime.NewTimeWindow(5 * time.Minute) // 5分钟窗口

// 添加事件
window.Add(time.Now(), "event1")
window.Add(time.Now().Add(-2*time.Minute), "event2")
window.Add(time.Now().Add(-10*time.Minute), "event3") // 超出窗口

// 获取窗口内的事件
events := window.GetEvents()
fmt.Printf("窗口内事件数量: %d\n", len(events))

// 清理过期事件
window.Cleanup()
```

### 退避算法

```go
// 创建指数退避
backoff := xtime.NewExponentialBackoff(
    100*time.Millisecond, // 初始延迟
    10*time.Second,       // 最大延迟
    2.0,                  // 乘数
    0.1,                  // 抖动因子
)

// 使用退避重试
for i := 0; i < 5; i++ {
    fmt.Printf("尝试第 %d 次\n", i+1)
    
    // 模拟操作失败
    if err := doSomeOperation(); err != nil {
        delay := backoff.NextDelay()
        fmt.Printf("操作失败，等待 %v 后重试\n", delay)
        time.Sleep(delay)
        continue
    }
    
    fmt.Println("操作成功")
    break
}

// 重置退避
backoff.Reset()
```

### 随机定时器

```go
// 创建随机间隔定时器
baseDuration := 1 * time.Second
jitter := 200 * time.Millisecond
randTicker := xtime.NewRandTicker(baseDuration, jitter)

// 启动定时器
go func() {
    for t := range randTicker.C {
        fmt.Printf("随机定时器触发: %v\n", t)
    }
}()

// 运行10秒后停止
time.Sleep(10 * time.Second)
randTicker.Stop()
```

## API 参考

### 时间常量

```go
// 预定义时间格式
const (
    DateFormat     = "2006-01-02"
    TimeFormat     = "15:04:05"
    DateTimeFormat = "2006-01-02 15:04:05"
    RFC3339Format  = time.RFC3339
)

// 时间单位常量
const (
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
    Hour        = time.Hour
    Day         = 24 * Hour
    Week        = 7 * Day
)
```

### 时间范围函数

```go
// 获取时间范围
func StartOfDay(t time.Time) time.Time
func EndOfDay(t time.Time) time.Time
func StartOfWeek(t time.Time) time.Time
func EndOfWeek(t time.Time) time.Time
func StartOfMonth(t time.Time) time.Time
func EndOfMonth(t time.Time) time.Time
func StartOfYear(t time.Time) time.Time
func EndOfYear(t time.Time) time.Time
```

### 时间戳函数

```go
// 时间戳操作
func DayBeginStamp(timestamp int64) int64
func DayEndStamp(timestamp int64) int64
func HourBeginStamp(timestamp int64) int64
func HourEndStamp(timestamp int64) int64
func DayScope(timestamp int64) (int64, int64)
func WeekScope(timestamp int64) (int64, int64)
func MonthScope(timestamp int64) (int64, int64)
```

### 定时器

```go
type Timer struct {
    // ...
}

func NewTimer(duration time.Duration) *Timer
func (t *Timer) Start(callback func())
func (t *Timer) Stop()
func (t *Timer) Reset(duration time.Duration)
func (t *Timer) IsRunning() bool
```

### 时间统计

```go
type TimeStat struct {
    // ...
}

func NewTimeStat() *TimeStat
func (ts *TimeStat) Nanosecond() int64
func (ts *TimeStat) Microsecond() int64
func (ts *TimeStat) Millisecond() int64
func (ts *TimeStat) Second() float64
func (ts *TimeStat) Reset()
```

### 退避算法

```go
type Backoff interface {
    NextDelay() time.Duration
    Reset()
}

func NewExponentialBackoff(initial, max time.Duration, multiplier, jitter float64) Backoff
func NewLinearBackoff(initial, increment, max time.Duration) Backoff
func NewConstantBackoff(duration time.Duration) Backoff
```

## 最佳实践

1. **时区处理**：
   - 始终明确时区，避免时区混乱
   - 服务器端统一使用UTC时间
   - 客户端展示时转换为本地时区

2. **性能优化**：
   - 缓存频繁使用的时间计算结果
   - 避免在循环中重复调用时间函数
   - 使用时间戳进行比较操作

3. **定时器使用**：
   - 及时停止不需要的定时器
   - 避免定时器泄露导致内存问题
   - 合理设置定时器间隔

4. **退避策略**：
   - 根据业务场景选择合适的退避算法
   - 设置合理的最大重试次数
   - 添加适当的抖动避免雷群效应

## 注意事项

- **时间精度**：Go的时间精度可能受操作系统影响
- **闰秒处理**：系统时间可能因闰秒调整而出现异常
- **并发安全**：某些操作需要考虑并发安全性
- **内存泄露**：定时器和时间窗口需要正确清理 