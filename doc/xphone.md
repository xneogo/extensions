# 手机号工具 (xphone)

## 简介

国际化手机号处理工具，支持：
- 多国手机号验证：支持中国、美国、日本等多个国家
- 手机号格式化：标准格式、拨号格式转换
- 手机号匿名化：隐藏敏感信息，支持多种匿名方式
- 区号解析：自动识别和分离国家区号
- 手机号规范化：统一格式处理和验证

## 使用示例

```go
import "github.com/xneogo/extensions/xphone"

// 手机号验证
phone := "13812345678"
isValid := xphone.IsValidChinaPhone(phone)
fmt.Printf("手机号 %s 有效性: %v\n", phone, isValid)

// 手机号格式化
formatted := xphone.FormatPhone(phone, xphone.China)
fmt.Printf("格式化后: %s\n", formatted)

// 手机号匿名化
anonymous := xphone.AnonymizePhone(phone)
fmt.Printf("匿名化: %s\n", anonymous)
``` 