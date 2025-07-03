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
 @Time    : 2025/4/7 -- 11:37
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xreflect xreflect/ptr.go
*/

package xreflect

import "reflect"

// PeelOffPtr
// []*Struct --> []Struct
// *Struct --> Struct
func PeelOffPtr(from interface{}) interface{} {
	fromType := reflect.TypeOf(from)
	fromVal := reflect.ValueOf(from)
	switch fromType.Kind() {
	case reflect.Slice:
		elemType := reflect.TypeOf(from).Elem()
		// 获取切片的元素类型
		switch elemType.Kind() {
		case reflect.Ptr:
			output := reflect.MakeSlice(reflect.SliceOf(elemType.Elem()), fromVal.Len(), fromVal.Len())

			// 遍历输入切片并转换每个元素
			for i := 0; i < fromVal.Len(); i++ {
				ele := fromVal.Index(i)
				// 检查元素是否是指针
				if ele.IsNil() {
					// 如果是指针为 nil，直接跳过
					continue
				}
				// 解引用指针并设置到目标切片
				output.Index(i).Set(reflect.Indirect(ele))
			}
			return output.Interface()
		}
	case reflect.Ptr:
		return reflect.ValueOf(from).Elem().Interface()
	default:
		// do nothing
		return from
		panic("unhandled default case")
	}
	return from
}

func ToPtr[elem any](from elem) func() (to *elem) {
	return func() (to *elem) {
		return &from
	}
}

func MultiToPtr[elem any](from []elem) func() (to []*elem) {
	return func() (to []*elem) {
		for i := range from {
			to = append(to, &from[i])
		}
		return to
	}
}

func ToPtrSafe[elem any](from elem, defaultV elem) func() (to *elem) {
	return func() (to *elem) {
		if from == nil {
			return &defaultV
		}
		return &from
	}
}
