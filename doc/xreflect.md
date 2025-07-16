# 反射工具 (xreflect)

## 简介

提供强大的反射操作功能，包括：
- 结构体到Map转换，支持自定义tag
- 指针操作：解引用、转换、批量处理
- 切片操作：类型转换、interface{}切片处理
- 递归解引用和类型匹配检查
- 零值检测和字段访问

## 安装

```bash
go get github.com/xneogo/extensions/xreflect
```

## 使用示例

### 基本用法

```go
import "github.com/xneogo/extensions/xreflect"

// 结构体转Map
type User struct {
    Name string `structs:"name"`
    Age  int    `structs:"age,omitempty"`
}

user := &User{Name: "张三", Age: 25}
s := xreflect.New(user)
m := s.Map()
```

### 指针操作

```go
// 指针解引用
var ptr *string
str := "Hello World"
ptr = &str

deref := xreflect.DerefPtr(ptr)
fmt.Printf("解引用结果: %v\n", deref)

// 获取指针指向的值
value := xreflect.GetPtrValue(ptr)
fmt.Printf("指针值: %v\n", value)
```

### 切片操作

```go
// 切片类型转换
slice := []interface{}{1, 2, 3, 4, 5}
intSlice := xreflect.ConvertSlice(slice, reflect.TypeOf(int(0)))
fmt.Printf("转换后的切片: %v\n", intSlice)

// 处理interface{}切片
data := []interface{}{"hello", 123, true}
processed := xreflect.ProcessInterfaceSlice(data)
fmt.Printf("处理结果: %v\n", processed)
```

### 类型检查

```go
// 零值检测
var empty string
var filled = "not empty"

isEmpty1 := xreflect.IsZero(empty)  // true
isEmpty2 := xreflect.IsZero(filled) // false

fmt.Printf("空字符串是零值: %v\n", isEmpty1)
fmt.Printf("非空字符串是零值: %v\n", isEmpty2)
```

### 字段访问

```go
type Person struct {
    Name    string
    Age     int
    Address string
}

person := &Person{
    Name:    "Alice",
    Age:     30,
    Address: "New York",
}

s := xreflect.New(person)

// 获取字段名列表
fieldNames := s.Names()
fmt.Printf("字段名: %v\n", fieldNames)

// 获取字段值列表
fieldValues := s.Values()
fmt.Printf("字段值: %v\n", fieldValues)

// 获取指定字段
nameField := s.Field("Name")
fmt.Printf("Name字段: %v\n", nameField)
```

### 结构体标签处理

```go
type Config struct {
    Host     string `json:"host" structs:"server_host"`
    Port     int    `json:"port" structs:"server_port,omitempty"`
    Username string `json:"username" structs:"-"`
    Password string `json:"password" structs:"pwd,omitempty"`
}

config := &Config{
    Host:     "localhost",
    Port:     8080,
    Username: "admin",
    Password: "secret",
}

s := xreflect.New(config)

// 使用structs标签生成Map
structsMap := s.Map()
fmt.Printf("structs标签Map: %v\n", structsMap)

// 获取带标签的字段
taggedFields := s.Fields()
for _, field := range taggedFields {
    fmt.Printf("字段: %s, 标签: %s, 值: %v\n", 
        field.Name(), field.Tag("structs"), field.Value())
}
```

## API 参考

### 主要类型

#### Struct
```go
type Struct struct {
    // ...
}

func New(s interface{}) *Struct
func (s *Struct) Map() map[string]interface{}
func (s *Struct) Names() []string
func (s *Struct) Values() []interface{}
func (s *Struct) Field(name string) *Field
func (s *Struct) Fields() []*Field
```

#### Field
```go
type Field struct {
    // ...
}

func (f *Field) Name() string
func (f *Field) Value() interface{}
func (f *Field) Tag(key string) string
func (f *Field) IsZero() bool
```

### 工具函数

```go
// 指针操作
func DerefPtr(ptr interface{}) interface{}
func GetPtrValue(ptr interface{}) interface{}

// 类型检查
func IsZero(v interface{}) bool
func TypeOf(v interface{}) reflect.Type

// 切片操作
func ConvertSlice(slice interface{}, targetType reflect.Type) interface{}
func ProcessInterfaceSlice(slice []interface{}) []interface{}
```

## 最佳实践

1. **结构体标签使用**：合理使用标签控制字段的序列化行为
2. **零值检测**：在处理可选字段时使用零值检测
3. **类型安全**：在类型转换前进行类型检查
4. **性能考虑**：反射操作相对较慢，避免在热点路径中频繁使用
5. **错误处理**：处理反射操作可能出现的panic

## 注意事项

- 反射操作会降低性能，应谨慎使用
- 某些操作可能会panic，建议使用recover机制
- 私有字段无法通过反射访问
- 复杂嵌套结构需要特殊处理 