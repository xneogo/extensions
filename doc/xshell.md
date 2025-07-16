# 命令行工具 (xshell)

## 简介

强大的命令执行工具，支持：
- 本地命令：执行本地系统命令
- 远程命令：支持SSH远程命令执行
- 文件操作：远程文件复制、移动等操作
- 超时控制：支持命令执行超时设置
- 错误处理：完善的命令执行状态检查

## 使用示例

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