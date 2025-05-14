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
 @Time    : 2025/4/7 -- 11:43
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xreflect xreflect/slice.go
*/

package xreflect

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

var (
	errNotASlice = errors.New("not a slice")
)

// ToSliceInterface
// return can be nil
func ToSliceInterface(ctx context.Context, arr interface{}) ([]interface{}, error) {
	if arr == nil {
		return nil, nil
	}
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return nil, errNotASlice
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret, nil
}

// MustToSliceInterface
// panic if error; return can be nil
func MustToSliceInterface(ctx context.Context, arr interface{}) []interface{} {
	slice, err := ToSliceInterface(ctx, arr)
	if err != nil {
		panic(fmt.Sprintf("convert interface slice error:%v", err))
		return nil
	}
	return slice
}
