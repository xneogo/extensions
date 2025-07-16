# 字符串工具 (xstring)

## 简介

高效的字符串处理工具，包含：
- UTF-8字符处理：正确处理Unicode字符
- Emoji处理：检测、过滤、替换Emoji表情
- 字符转换：全角半角转换、大小写转换
- 字符串拼接：高效的字符串连接
- 类型转换：任意类型到字符串的安全转换
- 字符串长度：按字符计算的真实长度

## 安装

```bash
go get github.com/xneogo/extensions/xstring
```

## 使用示例

### UTF-8字符处理

```go
import "github.com/xneogo/extensions/xstring"

// UTF-8字符处理
text := "Hello, 世界! 🌍"
utf8Chars := xstring.GetUtf8Chars(text, 5)  // 获取前5个UTF-8字符
fmt.Printf("前5个字符: %s\n", utf8Chars)

// 字符串长度（按字符计算）
length := xstring.StrLen(text)
fmt.Printf("字符长度: %d\n", length)
```

### Emoji处理

```go
// Emoji处理
hasEmoji := xstring.HaveEmoji(text)
cleanText := xstring.FilterEmoji(text)
replacedText := xstring.ReplaceEmoji(text, "[emoji]")
fmt.Printf("有Emoji: %v, 清理后: %s, 替换后: %s\n", hasEmoji, cleanText, replacedText)
```

### 字符转换

```go
// 全角半角转换
fullWidth := "１２３４５"
halfWidth := xstring.SBC2DBC(fullWidth)
fmt.Printf("全角转半角: %s -> %s\n", fullWidth, halfWidth)

// 首字母大写
name := "alice"
capitalized := xstring.UCFirst(name)
fmt.Printf("首字母大写: %s -> %s\n", name, capitalized)
```

### 字符串操作

```go
// 高效字符串拼接
result := xstring.Concat("Hello", " ", "World", "!")
fmt.Printf("拼接结果: %s\n", result)

// 任意类型转字符串
data := map[string]interface{}{"name": "Alice", "age": 30}
jsonStr := xstring.Any2String(data)
fmt.Printf("转换为字符串: %s\n", jsonStr)
```

## API 参考

### 核心函数

```go
func GetUtf8Chars(s string, count int) string
func StrLen(s string) int
func HaveEmoji(s string) bool
func FilterEmoji(s string) string
func ReplaceEmoji(s, replacement string) string
func SBC2DBC(s string) string
func UCFirst(s string) string
func Concat(strs ...string) string
func Any2String(v interface{}) string
``` 