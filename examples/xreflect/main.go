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
 @Time    : 2025/4/7 -- 14:34
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xreflect examples/xreflect/main.go
*/

package main

import (
	"fmt"
	"github.com/xneogo/extensions/xreflect"
	"reflect"
)

func main() {
	fmt.Println("-------------------------\nMirroring:")
	Mirroring()

	fmt.Println("-------------------------\nStruct2Map:")
	Struct2Map()

	fmt.Println("-------------------------\nPanicTypeMissMatch:")
	PanicTypeMissMatch()

	fmt.Println("-------------------------\nPtr:")
	Ptr()

	fmt.Println("-------------------------")

	fmt.Println("-------------------------")
}

func Mirroring() {
	type Sa struct {
		A string
	}
	type Sb struct {
		B string
	}

	type Ss struct {
		S int
	}

	src := &Sa{A: "a"}
	sb := new(Sb)
	ss := new(Ss)
	sa := new(Sa)
	sa2 := Sa{}

	err := xreflect.Mirroring(src, sb)
	fmt.Println(src, sb, err)
	// &{a} &{} src and tar must have the same struct type

	err = xreflect.Mirroring(src, ss)
	fmt.Println(src, ss, err)
	// &{a} &{0} src and tar must have the same struct type

	err = xreflect.Mirroring(src, sa)
	fmt.Println(src, sa, err)
	// &{a} &{a} <nil>
	// Success!!

	err = xreflect.Mirroring(src, sa2)
	fmt.Println(src, sa2, err)
	// &{a} {} src must be a pointer to a struct

}

func Struct2Map() {
	type S struct {
		A string
		B int
		C bool
		D float64
	}
	s := S{
		A: "a",
		B: 1,
		C: true,
		D: 1.1,
	}
	m, err := xreflect.Struct2Map(s, "")
	fmt.Println(m, err)
}

func PanicTypeMissMatch() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	xreflect.PanicTypeMissMatch("test", nil, nil)
}

func Ptr() {
	type S struct {
		A string
	}
	s := S{
		A: "a",
	}
	fmt.Println(fmt.Sprintf("type of %v is %s", s, reflect.TypeOf(s).Kind()))
	// type of {a} is struct
	fmt.Println(fmt.Sprintf("type of %v is %s", &s, reflect.TypeOf(&s).Kind()))
	// type of &{a} is ptr
	fmt.Println(fmt.Sprintf("type of %v is %s", *&s, reflect.TypeOf(*&s).Kind()))
	// type of {a} is struct s = *&s
	fmt.Println(fmt.Sprintf("original: %v - %s, after peel off : %v - %s", s, reflect.TypeOf(s).Kind(), xreflect.PeelOffPtr(s).(S), reflect.TypeOf(xreflect.PeelOffPtr(s).(S)).Kind()))
	// original: {a} - struct, after peel off : {a} - struct
	fmt.Println(fmt.Sprintf("original: %v - %s, after peel off : %v - %s", &s, reflect.TypeOf(&s).Kind(), xreflect.PeelOffPtr(&s).(S), reflect.TypeOf(xreflect.PeelOffPtr(&s).(S)).Kind()))
	// original: &{a} - ptr, after peel off : {a} - struct

	fmt.Println("Slice:")
	ss := []*S{
		&S{
			A: "a",
		},
		&S{
			A: "b",
		},
	}
	for _, v := range ss {
		fmt.Println(fmt.Sprintf("original: %v - %s", v, reflect.TypeOf(v).Kind()))
	}
	fmt.Println(fmt.Sprintf("original: %v - %s", ss, reflect.TypeOf(ss).Kind()))
	for _, v := range xreflect.PeelOffPtr(ss).([]S) {
		fmt.Println(fmt.Sprintf("after peel off: %v - %s", v, reflect.TypeOf(v).Kind()))
	}
	fmt.Println(fmt.Sprintf("after peel off: %v - %s", xreflect.PeelOffPtr(ss), reflect.TypeOf(xreflect.PeelOffPtr(ss)).Kind()))
	// original: [0x140001322c0 0x140001322d0] - slice

	fmt.Println(fmt.Sprintf("original: %v - %s", s, reflect.TypeOf(s).Kind()))
	// original: {a} - struct
	fmt.Println(fmt.Sprintf("toPtr: %v - %s", xreflect.ToPtr(s)(), reflect.TypeOf(xreflect.ToPtr(s)()).Kind()))
	// toPtr: &{a} - ptr
}
