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
 @Time    : 2025/7/9 -- 16:40
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xerror xerror/stack.go
*/

package xerror

import (
	"runtime"
)

type (
	Stack []uintptr
)

const (
	maxStackDepth = 32
)

func Callers(skip ...int) Stack {
	var depth int
	var spot [maxStackDepth]uintptr
	if len(skip) > 0 {
		depth = skip[0]
	}
	return spot[:runtime.Callers(depth, spot[:])]
}
