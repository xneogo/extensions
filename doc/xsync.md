# 同步工具 (xsync)

## 简介

并发安全的同步工具，提供：
- 自定义Mutex：支持TryLock的互斥锁
- 原子操作：Int32、Int64、Bool、String的原子操作
- 线程安全：无锁的并发安全数据结构
- 性能优化：基于channel的高效锁实现
- 兼容性：与标准库sync包完全兼容

## 安装

```bash
go get github.com/xneogo/extensions/xsync
```

## 使用示例

### 自定义互斥锁

```go
import "github.com/xneogo/extensions/xsync"

// 自定义Mutex
var mutex xsync.Mutex
mutex.Lock()
// 执行临界区代码
mutex.Unlock()

// 尝试加锁
if mutex.TryLock() {
    fmt.Println("获取锁成功")
    // 执行代码
    mutex.Unlock()
} else {
    fmt.Println("获取锁失败")
}
```

### 原子操作

```go
// 原子操作
counter := xsync.NewAtomicInt32(0)
counter.Add(1)
counter.Set(100)
value := counter.Get()
swapped := counter.CompareAndSwap(100, 200)
fmt.Printf("计数器值: %d, 交换成功: %v\n", value, swapped)

// 原子字符串
atomicStr := xsync.NewAtomicString("hello")
atomicStr.Set("world")
str := atomicStr.Get()
fmt.Printf("原子字符串: %s\n", str)

// 原子布尔
atomicBool := xsync.NewAtomicBool(false)
atomicBool.Set(true)
isTrue := atomicBool.Get()
fmt.Printf("原子布尔: %v\n", isTrue)
```

## API 参考

### 互斥锁

```go
type Mutex struct{}

func (m *Mutex) Lock()
func (m *Mutex) Unlock()
func (m *Mutex) TryLock() bool
```

### 原子类型

```go
type AtomicInt32 struct{}
func NewAtomicInt32(v int32) *AtomicInt32
func (a *AtomicInt32) Get() int32
func (a *AtomicInt32) Set(v int32)
func (a *AtomicInt32) Add(delta int32) int32
func (a *AtomicInt32) CompareAndSwap(old, new int32) bool

type AtomicString struct{}
func NewAtomicString(v string) *AtomicString
func (a *AtomicString) Get() string
func (a *AtomicString) Set(v string)

type AtomicBool struct{}
func NewAtomicBool(v bool) *AtomicBool
func (a *AtomicBool) Get() bool
func (a *AtomicBool) Set(v bool)
``` 