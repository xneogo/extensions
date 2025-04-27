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
 @Time    : 2024/7/13 -- 15:50
 @Author  : bishop ❤️ MONEY
 @Description: xrandstring.go
*/

package xrand

import "math/rand"

const letterDigit = "0123456789"

const letterAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int, ctype int) string {

	randSet := letterAlpha
	if ctype == 1 {
		randSet = letterDigit

	} else if ctype == 2 {
		randSet = letterAlpha
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = randSet[rand.Intn(len(randSet))]
	}
	return string(b)
}

// RandString Generate a string of digits and letters of length n
func RandString(n int) string {
	return randString(n, 2)
}

// RandDigit Generate a string of digits of length n
func RandDigit(n int) string {
	return randString(n, 1)
}
