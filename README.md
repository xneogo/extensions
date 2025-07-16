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

## 组件文档

### 🔍 核心工具
- **[反射工具 (xreflect)](./doc/xreflect.md)** - 强大的反射操作功能，结构体转Map、指针操作等
- **[时间工具 (xtime)](./doc/xtime.md)** - 全面的时间处理工具，时间范围、格式化、定时器等
- **[字符串工具 (xstring)](./doc/xstring.md)** - 高效的字符串处理，UTF-8处理、Emoji操作等
- **[切片工具 (xslice)](./doc/xslice.md)** - 强大的切片操作，去重、集合运算、栈操作等
- **[JSON工具 (xjson)](./doc/xjson.md)** - 灵活的JSON数据处理，路径查询、流式处理等

### 🔐 安全与加密
- **[加密工具 (xcrypto)](./doc/xcrypto.md)** - 完整的加密工具库，哈希、HMAC、AES、RSA等
- **[错误处理工具 (xerror)](./doc/xerror.md)** - 增强的错误处理，错误包装、堆栈跟踪等

### ⚡ 并发与性能
- **[同步工具 (xsync)](./doc/xsync.md)** - 并发安全的同步工具，自定义锁、原子操作等
- **[熔断器工具 (xbreaker)](./doc/xbreaker.md)** - 生产级熔断器实现，多种熔断策略
- **[限流器工具 (xrate)](./doc/xrate.md)** - 多种限流算法，漏桶、滑动窗口等
- **[队列工具 (xqueue)](./doc/xqueue.md)** - 线程安全的内存队列，FIFO、阻塞操作等

### 🎲 随机与算法
- **[随机数工具 (xrand)](./doc/xrand.md)** - 全功能随机数生成，波动函数、概率抽奖等
- **[版本比较工具 (versioncmp)](./doc/versioncmp.md)** - 智能版本号比较，语义化版本支持

### 📁 文件与系统
- **[文件工具 (xfile)](./doc/xfile.md)** - 全面的文件操作，检查、读写、流式处理等
- **[命令行工具 (xshell)](./doc/xshell.md)** - 强大的命令执行，本地、远程、文件操作等

### 📊 数据处理
- **[序列化工具 (xmarshal)](./doc/xmarshal.md)** - Properties格式序列化，结构体转换等
- **[分页工具 (xpage)](./doc/xpage.md)** - 简单实用的分页计算，页码转换等

### 🌐 网络与通信
- **[邮件工具 (xmail)](./doc/xmail.md)** - 功能完整的邮件发送，SMTP、附件支持等
- **[手机号工具 (xphone)](./doc/xphone.md)** - 国际化手机号处理，验证、格式化等

### 🔧 实用工具
- **[命令行参数工具 (xflag)](./doc/xflag.md)** - 增强的命令行参数处理，数组参数支持
- **[彩色日志工具 (colorlog)](./doc/colorlog.md)** - 美观的彩色日志，多级别、格式化支持
- **[其他工具 (xother)](./doc/xother.md)** - 实用的辅助工具集合，设备信息、通用函数等

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/xneogo/extensions/xstring"
    "github.com/xneogo/extensions/xtime"
    "github.com/xneogo/extensions/xrand"
)

func main() {
    // 字符串处理
    text := "Hello, 世界! 🌍"
    length := xstring.StrLen(text)
    fmt.Printf("字符长度: %d\n", length)

    // 时间操作
    now := time.Now()
    startOfDay := xtime.StartOfDay(now)
    fmt.Printf("今天开始: %v\n", startOfDay)

    // 随机数生成
    randomNum := xrand.RandBetween(1, 100)
    fmt.Printf("随机数: %d\n", randomNum)
}
```

### 常用组合

```go
// 加密和编码
import (
    "github.com/xneogo/extensions/xcrypto"
    "github.com/xneogo/extensions/xcrypto/base"
)

// 并发控制
import (
    "github.com/xneogo/extensions/xsync"
    "github.com/xneogo/extensions/xbreaker"
    "github.com/xneogo/extensions/xrate"
)

// 数据处理
import (
    "github.com/xneogo/extensions/xjson"
    "github.com/xneogo/extensions/xreflect"
    "github.com/xneogo/extensions/xslice"
)
```

## 贡献指南

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情
