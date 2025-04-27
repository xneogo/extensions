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
 @Time    : 2024/7/13 -- 14:53
 @Author  : bishop ❤️ MONEY
 @Description: xrand.go
*/

package xrand

import (
	"math/rand"
)

type numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Wave(base int64, lowRate int64, upRate int64) int64 {
	rate := rand.Intn(int((upRate-lowRate)+1)) + int(lowRate)

	return base + int64(float64(base)*float64(rate)/100)
}

// RandBetween choose a number between start and end.
func RandBetween[Unit numeric](start Unit, end Unit) Unit {
	if start > end {
		return 0
	}
	return Unit(rand.Intn(int(end-start)+1)) + start
}

// ChooseMN choose n from m eg: we have m=10 and n=5, then we will get [1,3,5,7,8]
func ChooseMN(m, n int) []int {
	chosen := make([]int, 0, n)
	if m <= n {
		for i := 0; i < m; i++ {
			chosen = append(chosen, i)
		}
		return chosen
	}
	hadChosen := map[int]bool{}
	for {
		if len(chosen) >= n {
			break
		}
		idx := rand.Intn(m)
		if !hadChosen[idx] {
			chosen = append(chosen, idx)
			hadChosen[idx] = true
		}
	}
	return chosen
}

// RandAnythingOnceFrom Get a random item from the slice only once
func RandAnythingOnceFrom[S ~[]T, T any](from S) T {
	length := len(from)
	r := rand.New(rand.NewSource(int64(length)))
	return from[r.Intn(length)]
}

// RandAnythingSomeFrom Get some random item from the slice. any item should show only once
func RandAnythingSomeFrom[S ~[]T, T any](from S, tar int) S {
	length := len(from)
	choice := ChooseMN(length, tar)
	var result S
	for _, idx := range choice {
		result = append(result, from[idx])
	}
	return result
}
