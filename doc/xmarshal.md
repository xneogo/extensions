# 序列化工具 (xmarshal)

## 简介

Properties格式序列化工具，提供：
- 结构体序列化：支持tag标记的属性序列化
- 键值对处理：Map和KV对的序列化和反序列化
- 嵌套结构：支持嵌套结构体和数组的处理
- 类型转换：自动处理基本类型的转换
- 兼容性：与Java Properties格式完全兼容

## 使用示例

```go
import "github.com/xneogo/extensions/xmarshal"

// 定义结构体
type Config struct {
    Host     string `properties:"host"`
    Port     int    `properties:"port"`
    Username string `properties:"username"`
    Password string `properties:"password"`
}

// 序列化为属性格式
config := Config{
    Host:     "localhost",
    Port:     8080,
    Username: "admin",
    Password: "secret",
}

data, err := xmarshal.Marshal(config)
if err == nil {
    fmt.Printf("序列化结果:\n%s\n", string(data))
}

// 反序列化
var newConfig Config
err = xmarshal.Unmarshal(data, &newConfig)
if err == nil {
    fmt.Printf("反序列化结果: %+v\n", newConfig)
}

// 从键值对反序列化
kv := map[string]string{
    "host":     "192.168.1.100",
    "port":     "3306",
    "username": "root",
    "password": "123456",
}
var dbConfig Config
err = xmarshal.UnmarshalKV(kv, &dbConfig)
if err == nil {
    fmt.Printf("从KV反序列化: %+v\n", dbConfig)
}
``` 