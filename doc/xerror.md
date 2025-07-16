# 错误处理工具 (xerror)

## 简介

增强的错误处理工具，提供：
- 错误包装：支持错误链和上下文信息
- 堆栈跟踪：自动记录错误发生的堆栈信息
- 错误类型：自定义错误类型和错误码
- 错误匹配：支持错误类型的判断和转换
- 格式化输出：丰富的错误信息格式化

## 使用示例

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