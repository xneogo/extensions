/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2025/5/13 -- 14:34
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xreflect examples/xreflect/main.go
*/

package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/xneogo/extensions/xreflect"
)

// User 示例结构体
type User struct {
	Name      string    `structs:"name"`
	Age       int       `structs:"age,omitempty"`
	Email     string    `structs:"email,omitempty"`
	CreatedAt time.Time `structs:"created_at"`
	Address   *Address  `structs:"address,omitnested"`
}

// Address 示例嵌套结构体
type Address struct {
	City    string `structs:"city"`
	Country string `structs:"country"`
}

func main() {
	// 示例1：结构体操作
	fmt.Println("示例1：结构体操作")
	user := &User{
		Name:      "张三",
		Age:       25,
		CreatedAt: time.Now(),
		Address: &Address{
			City:    "北京",
			Country: "中国",
		},
	}

	// 结构体转Map
	s := xreflect.New(user)
	m := s.Map()
	fmt.Printf("结构体转Map: %+v\n", m)

	// 获取字段
	field, ok := s.FieldOk("Name")
	if ok {
		fmt.Printf("字段Name的值: %v\n", field.Value())
	}

	// 检查零值
	fmt.Printf("结构体是否为零值: %v\n", s.IsZero())
	fmt.Printf("结构体是否有零值字段: %v\n", s.HasZero())

	// 示例2：指针操作
	fmt.Println("\n示例2：指针操作")
	// 解引用指针
	userValue := xreflect.PeelOffPtr(user)
	fmt.Printf("解引用指针: %+v\n", userValue)

	// 值转指针
	ptrFunc := xreflect.ToPtr(userValue.(User))
	userPtr := ptrFunc()
	fmt.Printf("值转指针: %+v\n", userPtr)

	// 切片指针转换
	users := []User{
		{Name: "张三", Age: 25},
		{Name: "李四", Age: 30},
	}
	ptrSliceFunc := xreflect.MultiToPtr(users)
	userPtrs := ptrSliceFunc()
	fmt.Printf("切片转指针切片: %+v\n", userPtrs)

	// 示例3：切片操作
	fmt.Println("\n示例3：切片操作")
	// 切片转interface{}切片
	ctx := context.Background()
	interfaceSlice, err := xreflect.ToSliceInterface(ctx, users)
	if err != nil {
		fmt.Printf("转换错误: %v\n", err)
	} else {
		fmt.Printf("切片转interface{}切片: %+v\n", interfaceSlice)
	}

	// 示例4：反射工具
	fmt.Println("\n示例4：反射工具")
	// 递归解引用
	value := xreflect.RecursiveIndirect(reflect.ValueOf(user))
	fmt.Printf("递归解引用: %+v\n", value.Interface())

	// 类型匹配检查
	srcType := reflect.TypeOf(user)
	tarType := reflect.TypeOf(&User{})
	if srcType == tarType {
		fmt.Println("类型匹配")
	} else {
		fmt.Println("类型不匹配")
	}
}
