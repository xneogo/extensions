# 随机数工具 (xrand)

## 简介

全功能随机数生成工具，包含：
- 基础随机数：指定范围内的随机数生成
- 波动函数：基于基础值的百分比波动
- 随机字符串：支持纯数字、字母数字组合
- 随机选择：从集合中随机选择单个或多个元素
- 概率抽奖：基于概率的中奖判断
- 随机定时器：带随机间隔的定时器

## 安装

```bash
go get github.com/xneogo/extensions/xrand
```

## 使用示例

### 基础随机数生成

```go
import "github.com/xneogo/extensions/xrand"

// 基础随机数生成
num := xrand.RandBetween(1, 100)    // 生成1-100之间的随机数
fmt.Printf("随机数: %d\n", num)

// 波动函数
base := int64(1000)
waveValue := xrand.Wave(base, -20, 30) // 基于1000，波动-20%到+30%
fmt.Printf("波动值: %d\n", waveValue)
```

### 随机字符串

```go
// 随机字符串
randomStr := xrand.RandString(8)     // 生成8位随机字符串(字母+数字)
randomDigit := xrand.RandDigit(6)    // 生成6位随机数字字符串
fmt.Printf("随机字符串: %s, 随机数字: %s\n", randomStr, randomDigit)
```

### 随机选择

```go
// 随机选择
items := []string{"apple", "banana", "cherry", "date", "elderberry"}
selected := xrand.RandAnythingOnceFrom(items)        // 随机选择一个
multiple := xrand.RandAnythingSomeFrom(items, 3)     // 随机选择3个
fmt.Printf("随机选择: %s, 多选: %v\n", selected, multiple)

// 从N个数中选择M个
chosen := xrand.ChooseMN(10, 5)  // 从0-9中选择5个数
fmt.Printf("选择的索引: %v\n", chosen)
```

### 概率抽奖

```go
// 概率抽奖
prob := int32(3000) // 30%概率
isWin := xrand.RanD.CheckProbabilityJackpot(prob)
fmt.Printf("是否中奖: %v\n", isWin)
```

### 随机定时器

```go
// 随机定时器
ticker := xrand.NewRandTicker(2*time.Second, 500*time.Millisecond)
go func() {
    for t := range ticker.C {
        fmt.Printf("随机定时器触发: %v\n", t)
    }
}()
// 记住调用 ticker.Stop() 停止定时器
```

## API 参考

### 基础函数

```go
func RandBetween(min, max int) int
func Wave(base int64, minPercent, maxPercent float64) int64
func RandString(length int) string
func RandDigit(length int) string
```

### 随机选择

```go
func RandAnythingOnceFrom[T any](items []T) T
func RandAnythingSomeFrom[T any](items []T, count int) []T
func ChooseMN(n, m int) []int
```

### 抽奖系统

```go
type Jackpot interface {
    CheckProbabilityJackpot(probability int32) bool
}

var RanD Jackpot
``` 