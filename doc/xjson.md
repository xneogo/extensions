# JSON工具 (xjson)

## 简介

灵活的JSON数据处理工具，提供：
- 路径查询：支持点分割路径的JSON数据提取
- 数组索引：支持数组元素的直接访问
- 嵌套数据处理：深度遍历JSON结构
- 流式处理：返回io.Reader接口，支持大数据处理
- 错误处理：完善的错误信息和异常处理

## 安装

```bash
go get github.com/xneogo/extensions/xjson
```

## 使用示例

### JSON路径查询

```go
import "github.com/xneogo/extensions/xjson"

// JSON路径查询
jsonData := `{
    "user": {
        "name": "Alice",
        "age": 30,
        "hobbies": ["reading", "gaming", "coding"]
    }
}`

// 查询用户名
reader, err := xjson.WalkEndReader([]byte(jsonData), "user.name", ".")
if err == nil {
    var name string
    json.NewDecoder(reader).Decode(&name)
    fmt.Printf("用户名: %s\n", name)
}

// 查询爱好列表
reader, err = xjson.WalkEndReader([]byte(jsonData), "user.hobbies", ".")
if err == nil {
    var hobbies []string
    json.NewDecoder(reader).Decode(&hobbies)
    fmt.Printf("爱好: %v\n", hobbies)
}

// 查询数组元素
reader, err = xjson.WalkEndReader([]byte(jsonData), "user.hobbies.0", ".")
if err == nil {
    var firstHobby string
    json.NewDecoder(reader).Decode(&firstHobby)
    fmt.Printf("第一个爱好: %s\n", firstHobby)
}
```

## API 参考

### 核心函数

```go
func WalkEndReader(data []byte, path, separator string) (io.Reader, error)
``` 