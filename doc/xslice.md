# 切片工具 (xslice)

## 简介

强大的切片操作工具，包含：
- 去重操作：支持自定义比较函数的去重
- 集合运算：交集、并集、差集等集合操作
- 栈操作：左右弹出、压入操作
- 切片计算：求和、平均值等数学计算
- 安全切片：边界检查的安全切片操作
- 随机排序：Fisher-Yates算法的随机排序

## 安装

```bash
go get github.com/xneogo/extensions/xslice
```

## 使用示例

### 去重操作

```go
import "github.com/xneogo/extensions/xslice"

// 去重
nums := []int{1, 2, 2, 3, 3, 4, 5}
unique := xslice.Distinct(nums)
fmt.Printf("去重后: %v\n", unique)
```

### 集合操作

```go
// 集合操作
set1 := []int{1, 2, 3, 4, 5}
set2 := []int{4, 5, 6, 7, 8}
sets := xslice.NewXSets(set1, set2)

intersection := sets.Intersection()   // 交集
union := sets.Union()                // 并集
distinctUnion := sets.DistinctUnion() // 去重并集
fmt.Printf("交集: %v, 并集: %v, 去重并集: %v\n", intersection, union, distinctUnion)

// 差集
diff := xslice.Except(set1, set2)  // 在set1但不在set2中的元素
fmt.Printf("差集: %v\n", diff)
```

### 栈操作

```go
// 栈操作
stack := []string{"a", "b", "c", "d"}
last := xslice.RPop(stack)           // 右弹出
first := xslice.LPop(stack)          // 左弹出
fmt.Printf("右弹出: %s, 左弹出: %s\n", last, first)
```

### 数学计算

```go
// 切片计算
numbers := []int{1, 2, 3, 4, 5}
sum := xslice.Sum(numbers)           // 求和
fmt.Printf("求和: %d\n", sum)

// 安全切片
safeCut := xslice.Cut(numbers, 1, 4)  // 安全切片[1:4)
fmt.Printf("安全切片: %v\n", safeCut)
```

## API 参考

### 基础函数

```go
func Distinct[T comparable](slice []T) []T
func Except[T comparable](slice1, slice2 []T) []T
func RPop[T any](slice []T) T
func LPop[T any](slice []T) T
func Sum[T constraints.Ordered](slice []T) T
func Cut[T any](slice []T, start, end int) []T
```

### 集合类型

```go
type XSets[T comparable] struct{}

func NewXSets[T comparable](sets ...[]T) *XSets[T]
func (xs *XSets[T]) Intersection() []T
func (xs *XSets[T]) Union() []T
func (xs *XSets[T]) DistinctUnion() []T
``` 