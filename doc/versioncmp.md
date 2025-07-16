# 版本比较工具 (versioncmp)

## 简介

智能的版本号比较工具，提供：
- 版本解析：自动解析各种格式的版本号
- 版本比较：支持大于、小于、等于等比较操作
- 格式化输出：统一的版本号格式化
- 兼容性：支持语义化版本和自定义版本格式

## 安装

```bash
go get github.com/xneogo/extensions/versioncmp
```

## 使用示例

### 基本用法

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

## API 参考

### 版本比较器

```go
type Version interface {
    Gt(other string) bool
    Lt(other string) bool
    Eq(other string) bool
    Gte(other string) bool
    Lte(other string) bool
    GetFormatVersion() string
}

func New(version string) Version
``` 