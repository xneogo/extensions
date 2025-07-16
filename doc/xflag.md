# 命令行参数工具 (xflag)

## 简介

增强的命令行参数处理，提供：
- 字符串数组：支持逗号分隔的字符串数组参数
- 参数解析：自动处理空格和特殊字符
- 标准兼容：与Go标准flag包完全兼容
- 灵活配置：支持多种参数格式和分隔符

## 使用示例

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