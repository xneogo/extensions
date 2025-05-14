# Extensions

[![Go Report Card](https://goreportcard.com/badge/github.com/xneogo/extensions)](https://goreportcard.com/report/github.com/xneogo/extensions)
[![GoDoc](https://godoc.org/github.com/xneogo/extensions?status.svg)](https://godoc.org/github.com/xneogo/extensions)
[![License](https://img.shields.io/github/license/xneogo/extensions)](LICENSE)

Extensions 是一个 Go 语言工具库集合，提供了一系列实用的扩展功能，帮助开发者更高效地进行 Go 开发。

## 功能特性

- 🚀 高性能：所有功能都经过性能优化
- 🛠 易用性：提供简单直观的 API
- 📦 模块化：按功能分类的独立模块
- 🧪 测试覆盖：完善的单元测试
- 📚 详细文档：每个模块都有详细的使用说明和示例

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

// HMAC
key := []byte("secret")
data := []byte(text)
hmac := xcrypto.HMacSHA256(key, data)
```

### 时间工具 (xtime)

```go
import "github.com/xneogo/extensions/xtime"

// 获取时间范围
startOfDay := xtime.StartOfDay(time.Now())
endOfDay := xtime.EndOfDay(time.Now())

// 时间戳操作
timestamp := xtime.DayBeginStamp()
```

## 模块说明

- `xreflect`: 反射工具，提供结构体操作、指针操作等功能
- `xcrypto`: 加密工具，提供哈希、HMAC、签名等功能
- `xtime`: 时间工具，提供时间范围、时间戳操作等功能
- `xphone`: 手机号工具，提供手机号验证、格式化等功能
- `xstring`: 字符串工具，提供字符串处理功能
- `xjson`: JSON工具，提供JSON处理功能
- `xerror`: 错误处理工具，提供错误处理功能
- `xrand`: 随机数工具，提供随机数生成功能
- `xsync`: 同步工具，提供同步原语功能
- `xpage`: 分页工具，提供分页功能
- `xother`: 其他工具，提供其他实用功能

## 贡献指南

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情
