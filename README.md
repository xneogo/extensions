# Extensions

[![Go Report Card](https://goreportcard.com/badge/github.com/xneogo/extensions)](https://goreportcard.com/report/github.com/xneogo/extensions)
[![GoDoc](https://godoc.org/github.com/xneogo/extensions?status.svg)](https://godoc.org/github.com/xneogo/extensions)
[![License](https://img.shields.io/github/license/xneogo/extensions)](LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/xneogo/extensions)

Extensions 是一个 Go 语言工具库集合，提供了一系列实用的扩展功能，帮助开发者更高效地进行 Go 开发。

## 功能特性

- 🚀 **高性能**：所有功能都经过性能优化，支持高并发场景
- 🛠 **易用性**：提供简单直观的 API，开箱即用
- 📦 **模块化**：按功能分类的独立模块，支持按需引入
- 🧪 **测试覆盖**：完善的单元测试，保证代码质量
- 📚 **详细文档**：每个模块都有详细的使用说明和示例
- 🔐 **安全可靠**：内置加密、签名、熔断等安全机制
- ⚡ **并发安全**：提供原子操作、信号量等并发控制工具
- 🎯 **生产就绪**：支持限流、熔断、日志等生产环境必需功能

## 安装

```bash
go get github.com/xneogo/extensions
```

## 使用示例

### 反射工具 (xreflect)

```go
import "github.com/xneogo/extensions/xreflect"

// 结构体转Map
type User struct {
    Name string `structs:"name"`
    Age  int    `structs:"age,omitempty"`
}

user := &User{Name: "张三", Age: 25}
s := xreflect.New(user)
m := s.Map()
```

### 加密工具 (xcrypto)

```go
import "github.com/xneogo/extensions/xcrypto"

// 哈希计算
text := "Hello, World!"
md5 := xcrypto.MD5.DoString(text)
sha256 := xcrypto.SHA256.DoString(text)
crc32 := xcrypto.CRC32.DoString(text)

// HMAC签名
key := []byte("secret")
data := []byte(text)
hmacSHA256 := xcrypto.HMacSHA256(key, data)
hmacMD5 := xcrypto.HMacMD5(key, data)

// AES加密
import "github.com/xneogo/extensions/xcrypto/aescipher"
ctx := context.Background()
plaintext := []byte("sensitive data")
secretKey := []byte("1234567890123456") // 16字节密钥
iv := []byte("1234567890123456")        // 16字节IV

// ECB模式加密
encrypted, err := aescipher.ECB.Encrypt(ctx, plaintext, secretKey, iv, nil)
decrypted, err := aescipher.ECB.Decrypt(ctx, encrypted, secretKey, iv, nil)

// Base编码
import "github.com/xneogo/extensions/xcrypto/base"
encoded := base.Base64.SEncode("hello world")
decoded, err := base.Base64.SDecode(encoded)
encoded62 := base.Base62.SEncode("hello world")
```

### 时间工具 (xtime)

```go
import "github.com/xneogo/extensions/xtime"

// 获取时间范围
now := time.Now()
startOfDay := xtime.StartOfDay(now)
endOfDay := xtime.EndOfDay(now)
startOfWeek := xtime.StartOfWeek(now)
startOfMonth := xtime.StartOfMonth(now)

// 时间戳操作
nowStamp := time.Now().Unix()
dayBegin := xtime.DayBeginStamp(nowStamp)
hourBegin := xtime.HourBeginStamp(nowStamp)

// 时间范围
dayStart, dayEnd := xtime.DayScope(nowStamp)
weekStart, weekEnd := xtime.WeekScope(nowStamp)

// 定时器
timer := xtime.NewTimer(2 * time.Second)
timer.Start(func() {
    fmt.Println("定时器触发")
})
timer.Stop()

// 运行时间统计
timeStat := xtime.NewTimeStat()
// ... 执行业务逻辑 ...
duration := timeStat.Millisecond() // 获取运行时间(毫秒)
```

### 命令行工具 (xshell)

```go
import "github.com/xneogo/extensions/xshell"

// 执行命令
ctx := context.Background()
output, err := xshell.Exec(ctx, "", "ls -la")

// 远程执行命令
output, err = xshell.Exec(ctx, "192.168.1.100", "ps aux")

// 文件操作
err = xshell.CP(ctx, "", "/tmp/file1.txt", "/tmp/file2.txt")
err = xshell.MV(ctx, "", "/tmp/old.txt", "/tmp/new.txt")
```

### 文件工具 (xfile)

```go
import "github.com/xneogo/extensions/xfile"

// 文件检查
exists := xfile.FileExists("/path/to/file.txt")
isFile, err := xfile.IsFile("/path/to/file.txt")
size, err := xfile.FileSize("/path/to/file.txt")

// 文件操作
err = xfile.CopyFile("/src/file.txt", "/dst/file.txt")
content, err := xfile.ReadAllString("/path/to/file.txt")
err = xfile.WriteAll("/path/to/file.txt", []byte("content"))

// 行处理
err = xfile.EachLine("/path/to/file.txt", func(line string) bool {
    fmt.Println(line)
    return true // 继续读取
}, true) // 跳过空行
```

### 熔断器 (xbreaker)

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

// 简单计数器熔断器
xbreaker.StatBreaker("cluster1", "table1", err) // 统计错误
allowed := xbreaker.Entry("cluster1", "table1") // 检查是否允许请求
```

### 限流器 (xrate)

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

// 滑动窗口限流器
windowLimiter := xrate.NewSlidingWindowRateLimiter(
    1000,       // QPS阈值
    10,         // 窗口大小
    100,        // 子窗口间隔(ms)
)

// 动态修改限流阈值
limiter.ChangeQpsThreshold(200)
```

### 彩色日志 (colorlog)

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

// 创建自定义日志器
logger := colorlog.New()
logger.SetLevel(colorlog.WarnLevel)
logger.Warn(ctx, "这是警告日志")
```

### 版本比较 (versioncmp)

```go
import "github.com/xneogo/extensions/versioncmp"

// 创建版本比较器
current := versioncmp.New("2.3.5.10239")

// 版本比较
fmt.Println(current.Gt("2.2"))        // true: 2.3.5.10239 > 2.2
fmt.Println(current.Lt("3.0.0"))      // true: 2.3.5.10239 < 3.0.0
fmt.Println(current.Eq("2.3.5.10239")) // true: 版本相等
fmt.Println(current.Gte("2.3.5"))     // true: 2.3.5.10239 >= 2.3.5

// 获取格式化后的版本字符串
formatted := current.GetFormatVersion()
```

### 随机数工具 (xrand)

```go
import "github.com/xneogo/extensions/xrand"

// 基础随机数生成
num := xrand.RandBetween(1, 100)    // 生成1-100之间的随机数
fmt.Printf("随机数: %d\n", num)

// 波动函数
base := int64(1000)
waveValue := xrand.Wave(base, -20, 30) // 基于1000，波动-20%到+30%
fmt.Printf("波动值: %d\n", waveValue)

// 随机字符串
randomStr := xrand.RandString(8)     // 生成8位随机字符串(字母+数字)
randomDigit := xrand.RandDigit(6)    // 生成6位随机数字字符串
fmt.Printf("随机字符串: %s, 随机数字: %s\n", randomStr, randomDigit)

// 随机选择
items := []string{"apple", "banana", "cherry", "date", "elderberry"}
selected := xrand.RandAnythingOnceFrom(items)        // 随机选择一个
multiple := xrand.RandAnythingSomeFrom(items, 3)     // 随机选择3个
fmt.Printf("随机选择: %s, 多选: %v\n", selected, multiple)

// 从N个数中选择M个
chosen := xrand.ChooseMN(10, 5)  // 从0-9中选择5个数
fmt.Printf("选择的索引: %v\n", chosen)

// 概率抽奖
prob := int32(3000) // 30%概率
isWin := xrand.RanD.CheckProbabilityJackpot(prob)
fmt.Printf("是否中奖: %v\n", isWin)

// 随机定时器
ticker := xrand.NewRandTicker(2*time.Second, 500*time.Millisecond)
go func() {
    for t := range ticker.C {
        fmt.Printf("随机定时器触发: %v\n", t)
    }
}()
// 记住调用 ticker.Stop() 停止定时器
```

### 切片工具 (xslice)

```go
import "github.com/xneogo/extensions/xslice"

// 去重
nums := []int{1, 2, 2, 3, 3, 4, 5}
unique := xslice.Distinct(nums)
fmt.Printf("去重后: %v\n", unique)

// 集合操作
set1 := []int{1, 2, 3, 4, 5}
set2 := []int{4, 5, 6, 7, 8}
sets := xslice.NewXSets(set1, set2)

intersection := sets.Intersection()   // 交集
union := sets.Union()                // 并集
distinctUnion := sets.DistinctUnion() // 去重并集
fmt.Printf("交集: %v, 并集: %v, 去重并集: %v\n", intersection, union, distinctUnion)

// 差集
diff := xslice.Except(set1, set2)  // 在set1但不在set2中的元素
fmt.Printf("差集: %v\n", diff)

// 栈操作
stack := []string{"a", "b", "c", "d"}
last := xslice.RPop(stack)           // 右弹出
first := xslice.LPop(stack)          // 左弹出
fmt.Printf("右弹出: %s, 左弹出: %s\n", last, first)

// 切片计算
numbers := []int{1, 2, 3, 4, 5}
sum := xslice.Sum(numbers)           // 求和
fmt.Printf("求和: %d\n", sum)

// 安全切片
safeCut := xslice.Cut(numbers, 1, 4)  // 安全切片[1:4)
fmt.Printf("安全切片: %v\n", safeCut)
```

### 字符串工具 (xstring)

```go
import "github.com/xneogo/extensions/xstring"

// UTF-8字符处理
text := "Hello, 世界! 🌍"
utf8Chars := xstring.GetUtf8Chars(text, 5)  // 获取前5个UTF-8字符
fmt.Printf("前5个字符: %s\n", utf8Chars)

// 字符串长度（按字符计算）
length := xstring.StrLen(text)
fmt.Printf("字符长度: %d\n", length)

// Emoji处理
hasEmoji := xstring.HaveEmoji(text)
cleanText := xstring.FilterEmoji(text)
replacedText := xstring.ReplaceEmoji(text, "[emoji]")
fmt.Printf("有Emoji: %v, 清理后: %s, 替换后: %s\n", hasEmoji, cleanText, replacedText)

// 全角半角转换
fullWidth := "１２３４５"
halfWidth := xstring.SBC2DBC(fullWidth)
fmt.Printf("全角转半角: %s -> %s\n", fullWidth, halfWidth)

// 首字母大写
name := "alice"
capitalized := xstring.UCFirst(name)
fmt.Printf("首字母大写: %s -> %s\n", name, capitalized)

// 高效字符串拼接
result := xstring.Concat("Hello", " ", "World", "!")
fmt.Printf("拼接结果: %s\n", result)

// 任意类型转字符串
data := map[string]interface{}{"name": "Alice", "age": 30}
jsonStr := xstring.Any2String(data)
fmt.Printf("转换为字符串: %s\n", jsonStr)
```

### JSON工具 (xjson)

```go
import "github.com/xneogo/extensions/xjson"

// JSON路径查询
jsonData := `{
    "user": {
        "name": "Alice",
        "age": 30,
        "hobbies": ["reading", "gaming", "coding"]
    }
}`

// 查询用户名
reader, err := xjson.WalkEndReader([]byte(jsonData), "user.name", ".")
if err == nil {
    var name string
    json.NewDecoder(reader).Decode(&name)
    fmt.Printf("用户名: %s\n", name)
}

// 查询爱好列表
reader, err = xjson.WalkEndReader([]byte(jsonData), "user.hobbies", ".")
if err == nil {
    var hobbies []string
    json.NewDecoder(reader).Decode(&hobbies)
    fmt.Printf("爱好: %v\n", hobbies)
}

// 查询数组元素
reader, err = xjson.WalkEndReader([]byte(jsonData), "user.hobbies.0", ".")
if err == nil {
    var firstHobby string
    json.NewDecoder(reader).Decode(&firstHobby)
    fmt.Printf("第一个爱好: %s\n", firstHobby)
}
```

### 同步工具 (xsync)

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

### 分页工具 (xpage)

```go
import "github.com/xneogo/extensions/xpage"

// 页码转换为偏移量和限制
page := 2
size := 10
offset, limit := xpage.PageSize2OffsetLimit(page, size)
fmt.Printf("第%d页，每页%d条 -> 偏移量: %d, 限制: %d\n", page, size, offset, limit)

// 偏移量转换为页码
offset = 20
limit = 10
page, size = xpage.OffsetLimit2PageSize(offset, limit)
fmt.Printf("偏移量: %d, 限制: %d -> 第%d页，每页%d条\n", offset, limit, page, size)

// 实际使用示例
func getUsers(page, size int) {
    offset, limit := xpage.PageSize2OffsetLimit(page, size)
    // 使用offset和limit进行数据库查询
    // SELECT * FROM users LIMIT offset, limit
    fmt.Printf("查询用户: LIMIT %d, %d\n", offset, limit)
}
```

### 序列化工具 (xmarshal)

```go
import "github.com/xneogo/extensions/xmarshal"

// 定义结构体
type Config struct {
    Host     string `properties:"host"`
    Port     int    `properties:"port"`
    Username string `properties:"username"`
    Password string `properties:"password"`
}

// 序列化为属性格式
config := Config{
    Host:     "localhost",
    Port:     8080,
    Username: "admin",
    Password: "secret",
}

data, err := xmarshal.Marshal(config)
if err == nil {
    fmt.Printf("序列化结果:\n%s\n", string(data))
}

// 反序列化
var newConfig Config
err = xmarshal.Unmarshal(data, &newConfig)
if err == nil {
    fmt.Printf("反序列化结果: %+v\n", newConfig)
}

// 从键值对反序列化
kv := map[string]string{
    "host":     "192.168.1.100",
    "port":     "3306",
    "username": "root",
    "password": "123456",
}
var dbConfig Config
err = xmarshal.UnmarshalKV(kv, &dbConfig)
if err == nil {
    fmt.Printf("从KV反序列化: %+v\n", dbConfig)
}
```

### 队列工具 (xqueue)

```go
import "github.com/xneogo/extensions/xqueue"

// 创建队列
queue := xqueue.NewQueue()

// 生产者
go func() {
    for i := 0; i < 5; i++ {
        err := queue.Push(fmt.Sprintf("消息-%d", i))
        if err != nil {
            fmt.Printf("推送失败: %v\n", err)
        } else {
            fmt.Printf("推送消息: 消息-%d\n", i)
        }
        time.Sleep(100 * time.Millisecond)
    }
}()

// 消费者
go func() {
    for {
        msg, err := queue.Pop()
        if err != nil {
            if err == xqueue.ErrQueueClosed {
                fmt.Println("队列已关闭")
                break
            }
            fmt.Printf("弹出失败: %v\n", err)
            continue
        }
        fmt.Printf("消费消息: %s (ID: %s)\n", msg.Body, msg.ID)
    }
}()

// 队列大小
size := queue.Size()
fmt.Printf("队列大小: %d\n", size)

// 关闭队列
time.Sleep(1 * time.Second)
queue.Close()
```

### 命令行参数工具 (xflag)

```go
import (
    "flag"
    "github.com/xneogo/extensions/xflag"
)

// 字符串数组参数
var hosts xflag.ArrayStringFlags
flag.Var(&hosts, "host", "服务器地址列表，逗号分隔")

// 解析命令行参数
flag.Parse()

// 使用参数
for _, host := range hosts {
    fmt.Printf("连接服务器: %s\n", host)
}

// 命令行示例: ./app -host=localhost,127.0.0.1,192.168.1.100
```

### 邮件工具 (xmail)

```go
import "github.com/xneogo/extensions/xmail"

// 创建邮件客户端
email := xmail.NewEmail("smtp.example.com", "sender@example.com", "password", "发送者名称")

// 发送邮件
err := email.Send(
    []string{"recipient@example.com"},  // 收件人
    []string{"cc@example.com"},         // 抄送
    []string{"bcc@example.com"},        // 密送
    "邮件主题",                           // 主题
    "邮件内容",                           // 内容
)
if err != nil {
    fmt.Printf("发送邮件失败: %v\n", err)
} else {
    fmt.Println("邮件发送成功")
}

// 带附件的邮件
err = email.SendWithAttachment(
    []string{"recipient@example.com"},
    []string{},
    []string{},
    "带附件的邮件",
    "请查看附件。",
    []string{"/path/to/attachment.pdf"}, // 附件路径
)
```

### 错误处理工具 (xerror)

```go
import "github.com/xneogo/extensions/xerror"

// 创建错误
err := xerror.New("业务错误")
fmt.Printf("错误: %v\n", err)

// 带堆栈的错误
stackErr := xerror.NewWithStack("带堆栈的错误")
fmt.Printf("错误堆栈: %v\n", stackErr.Stack())

// 错误包装
wrappedErr := xerror.Wrap(err, "上层错误")
fmt.Printf("包装错误: %v\n", wrappedErr)

// 格式化错误
formattedErr := xerror.Errorf("格式化错误: %s", "参数")
fmt.Printf("格式化错误: %v\n", formattedErr)

// 错误链
if xerror.Is(wrappedErr, err) {
    fmt.Println("错误匹配")
}

// 错误类型断言
var xErr *xerror.XError
if xerror.As(wrappedErr, &xErr) {
    fmt.Printf("错误类型: %T\n", xErr)
}
```

## 模块说明

### 🔍 反射工具 (xreflect)
提供强大的反射操作功能，包括：
- 结构体到Map转换，支持自定义tag
- 指针操作：解引用、转换、批量处理
- 切片操作：类型转换、interface{}切片处理
- 递归解引用和类型匹配检查
- 零值检测和字段访问

### 🔐 加密工具 (xcrypto)
完整的加密工具库，包含：
- 哈希算法：MD5、SHA1、SHA256、SHA512、CRC32
- HMAC签名：支持多种哈希算法的HMAC
- 对称加密：AES加密，支持ECB、CBC等模式
- 编码解码：Base64、Base62、Base58等编码
- 填充方案：PKCS5、PKCS7、Zero padding等
- 证书处理：PEM格式证书加载和验证

### ⏰ 时间工具 (xtime)
全面的时间处理工具，提供：
- 时间范围：获取日、周、月、年的开始和结束时间
- 时间戳操作：Unix时间戳的各种转换和计算
- 时间格式化：多种预定义格式和自定义格式
- 定时器：可配置的定时器和运行时间统计
- 时间范围查询：支持绝对时间和相对时间范围

### 📱 手机号工具 (xphone)
国际化手机号处理工具，支持：
- 多国手机号验证：支持中国、美国、日本等多个国家
- 手机号格式化：标准格式、拨号格式转换
- 手机号匿名化：隐藏敏感信息，支持多种匿名方式
- 区号解析：自动识别和分离国家区号
- 手机号规范化：统一格式处理和验证

### 🔤 字符串工具 (xstring)
高效的字符串处理工具，包含：
- UTF-8字符处理：正确处理Unicode字符
- Emoji处理：检测、过滤、替换Emoji表情
- 字符转换：全角半角转换、大小写转换
- 字符串拼接：高效的字符串连接
- 类型转换：任意类型到字符串的安全转换
- 字符串长度：按字符计算的真实长度

### 📊 JSON工具 (xjson)
灵活的JSON数据处理工具，提供：
- 路径查询：支持点分割路径的JSON数据提取
- 数组索引：支持数组元素的直接访问
- 嵌套数据处理：深度遍历JSON结构
- 流式处理：返回io.Reader接口，支持大数据处理
- 错误处理：完善的错误信息和异常处理

### 🎲 随机数工具 (xrand)
全功能随机数生成工具，包含：
- 基础随机数：指定范围内的随机数生成
- 波动函数：基于基础值的百分比波动
- 随机字符串：支持纯数字、字母数字组合
- 随机选择：从集合中随机选择单个或多个元素
- 概率抽奖：基于概率的中奖判断
- 随机定时器：带随机间隔的定时器

### 🔄 原语工具 (xsync)
并发安全的同步工具，提供：
- 自定义Mutex：支持TryLock的互斥锁
- 原子操作：Int32、Int64、Bool、String的原子操作
- 线程安全：无锁的并发安全数据结构
- 性能优化：基于channel的高效锁实现
- 兼容性：与标准库sync包完全兼容

### 📄 分页工具 (xpage)
简单实用的分页计算工具，支持：
- 页码转换：页码和偏移量的双向转换
- 数据库分页：直接适用于SQL LIMIT查询
- 边界处理：自动处理边界条件和异常情况
- 灵活配置：支持不同的分页大小设置

### 🔧 序列化工具 (xmarshal)
Properties格式序列化工具，提供：
- 结构体序列化：支持tag标记的属性序列化
- 键值对处理：Map和KV对的序列化和反序列化
- 嵌套结构：支持嵌套结构体和数组的处理
- 类型转换：自动处理基本类型的转换
- 兼容性：与Java Properties格式完全兼容

### 📬 队列工具 (xqueue)
线程安全的内存队列，包含：
- FIFO队列：先进先出的消息队列
- 阻塞操作：支持阻塞式的Push和Pop操作
- 消息ID：自动生成唯一的消息标识
- 队列状态：支持队列大小查询和关闭操作
- 并发安全：完全线程安全的队列实现

### 🏷️ 命令行参数工具 (xflag)
增强的命令行参数处理，提供：
- 字符串数组：支持逗号分隔的字符串数组参数
- 参数解析：自动处理空格和特殊字符
- 标准兼容：与Go标准flag包完全兼容
- 灵活配置：支持多种参数格式和分隔符

### 📧 邮件工具 (xmail)
功能完整的邮件发送工具，支持：
- SMTP发送：支持各种SMTP服务器
- 多收件人：支持收件人、抄送、密送
- 邮件编码：RFC2047编码支持
- 附件发送：支持多种格式的附件
- 错误处理：完善的邮件发送状态检查

### ❌ 错误处理工具 (xerror)
增强的错误处理工具，提供：
- 错误包装：支持错误链和上下文信息
- 堆栈跟踪：自动记录错误发生的堆栈信息
- 错误类型：自定义错误类型和错误码
- 错误匹配：支持错误类型的判断和转换
- 格式化输出：丰富的错误信息格式化

### 🔪 切片工具 (xslice)
强大的切片操作工具，包含：
- 去重操作：支持自定义比较函数的去重
- 集合运算：交集、并集、差集等集合操作
- 栈操作：左右弹出、压入操作
- 切片计算：求和、平均值等数学计算
- 安全切片：边界检查的安全切片操作
- 随机排序：Fisher-Yates算法的随机排序

### 🛡️ 熔断器工具 (xbreaker)
生产级的熔断器实现，提供：
- 熔断策略：基于错误率和QPS的熔断策略
- 状态管理：关闭、开启、半开三种状态
- 统计功能：实时统计请求成功率和响应时间
- 自动恢复：支持自动恢复和手动恢复
- 配置灵活：支持动态配置和热更新

### 🚦 限流工具 (xrate)
多种限流算法实现，包含：
- 漏桶算法：平滑限流，适合突发流量控制
- 滑动窗口：精确限流，适合QPS控制
- 动态配置：支持运行时修改限流阈值
- 多维度限流：支持按不同维度的限流策略
- 性能优化：高性能的限流实现

### 📁 文件工具 (xfile)
全面的文件操作工具，提供：
- 文件检查：文件存在性、类型、大小检查
- 文件操作：复制、移动、删除等基本操作
- 内容读写：支持字符串和字节数组的读写
- 行处理：按行读取和处理文件内容
- 流式处理：支持大文件的流式读取

### 🌈 彩色日志工具 (colorlog)
美观的彩色日志工具，支持：
- 彩色输出：不同级别的日志使用不同颜色
- 日志级别：Debug、Info、Warn、Error等级别
- 格式化输出：支持printf风格的格式化日志
- 自定义配置：可配置日志级别和颜色方案
- 上下文支持：支持context.Context的日志记录

### 📊 版本比较工具 (versioncmp)
智能的版本号比较工具，提供：
- 版本解析：自动解析各种格式的版本号
- 版本比较：支持大于、小于、等于等比较操作
- 格式化输出：统一的版本号格式化
- 兼容性：支持语义化版本和自定义版本格式

### 📦 命令行工具 (xshell)
强大的命令执行工具，支持：
- 本地命令：执行本地系统命令
- 远程命令：支持SSH远程命令执行
- 文件操作：远程文件复制、移动等操作
- 超时控制：支持命令执行超时设置
- 错误处理：完善的命令执行状态检查

### 🔧 其他工具 (xother)
实用的辅助工具集合，包含：
- 设备信息：移动设备标识符处理
- 通用工具：常用的辅助函数和工具
- 平台适配：跨平台的兼容性处理
- 性能优化：各种性能优化的工具函数

## 贡献指南

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情
