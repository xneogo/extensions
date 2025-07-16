# 邮件工具 (xmail)

## 简介

功能完整的邮件发送工具，支持：
- SMTP发送：支持各种SMTP服务器
- 多收件人：支持收件人、抄送、密送
- 邮件编码：RFC2047编码支持
- 附件发送：支持多种格式的附件
- 错误处理：完善的邮件发送状态检查

## 使用示例

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