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
 @Time    : 2024/11/5 -- 15:51
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: reflect.go
*/

package xreflect

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// Struct2Map
// function by name
// converts a struct to a map
// type for each field of the struct must be built-in type
func Struct2Map(target interface{}, useTag string) (map[string]interface{}, error) {
	if nil == target {
		return nil, nil
	}
	v := reflect.ValueOf(target)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("target must be a struct type")
	}
	t := v.Type()
	result := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		keyName := getKey(t.Field(i), useTag)
		if "" == keyName {
			continue
		}
		result[keyName] = v.Field(i).Interface()
	}
	return result, nil
}

func getKey(field reflect.StructField, useTag string) string {
	if !isExportedField(field.Name) {
		return ""
	}
	if field.Type.Kind() == reflect.Ptr {
		return ""
	}
	if "" == useTag {
		return field.Name
	}
	tag, ok := field.Tag.Lookup(useTag)
	if !ok {
		return ""
	}
	return ResolveTagName(tag)
}

func isExportedField(name string) bool {
	return strings.ToUpper(name) == name
}

func ResolveTagName(tag string) string {
	idx := strings.IndexByte(tag, ',')
	if -1 == idx {
		return tag
	}
	return tag[:idx]
}

// Mirroring
// copy what`s in src to tar like a mirror
func Mirroring(src, tar any) error {
	srcType := reflect.TypeOf(src)
	tarType := reflect.TypeOf(tar)
	if tarType.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("tar:%s must be a pointer to a struct", tarType.Kind()))
	}

	srcValue := reflect.ValueOf(src)
	tarValue := reflect.ValueOf(tar)
	// recalculate type based on Value`s type kind
	// if Value is ptr, we need to get it`s real value by calling Elem()
	if srcValue.Type().Kind() == reflect.Ptr {
		srcType = srcType.Elem()
		srcValue = srcValue.Elem()
	}
	if tarValue.Type().Kind() == reflect.Ptr {
		tarType = tarType.Elem()
		tarValue = tarValue.Elem()
	}

	if !match(srcType, tarType) {
		return errors.New(fmt.Sprintf("src:%s and tar:%s must have the same struct type", srcType, tarType))
	}

	tarValue.Set(srcValue)
	return nil
}

func match(srcT, tarT reflect.Type) bool {
	return srcT == tarT
}

func PanicTypeMissMatch(msg string, srcT, tarT reflect.Type) {
	if srcT.Kind() != tarT.Kind() {
		panic(fmt.Sprintf("%s: %v != %v", msg, srcT, tarT))
	}
}

func RecursiveIndirect(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}

func RecursiveIndirectType(p reflect.Type) reflect.Type {
	for p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	return p
}
