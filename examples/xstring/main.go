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
 @Time    : 2025/4/2 -- 15:13
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xstring examples/xstring/main.go
*/

package main

import (
	"fmt"
	"github.com/xneogo/extensions/xstring"
)

func main() {
	Example()
}

func Example() {
	fmt.Println(xstring.GetInvalidUtf8String("你好世界", 3))
	fmt.Println(xstring.Bytes2Str([]byte("你好世界")))
	fmt.Println(xstring.Str2Bytes("你好世界"))
	fmt.Println(xstring.Concat("你好", "世界"))
	fmt.Println(xstring.Any2String("你好世界"))
	fmt.Println(xstring.FilterEmoji("你好世界 🌲"))
	fmt.Println(xstring.HaveEmoji("你好世界 🌲"))
	fmt.Println(xstring.ReplaceEmoji("你好世界 🌲", "☀️"))
	fmt.Println(xstring.UCFirst("你好世界"))
	fmt.Println(xstring.SBC2DBC("你好１２［　］６７８"))
	fmt.Println(xstring.GetUtf8Chars("你好世界", 3))
	fmt.Println(xstring.StrLen("你好世界"))
	fmt.Println(xstring.GetInvalidUtf8String("asdf你好世界", 3))
}
