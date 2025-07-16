# 文件工具 (xfile)

## 简介

全面的文件操作工具，提供：
- 文件检查：文件存在性、类型、大小检查
- 文件操作：复制、移动、删除等基本操作
- 内容读写：支持字符串和字节数组的读写
- 行处理：按行读取和处理文件内容
- 流式处理：支持大文件的流式读取

## 使用示例

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