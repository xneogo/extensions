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
 @Time    : 2025/4/27 -- 18:26
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xrand examples/xrand/main.go
*/

package main

import (
	"fmt"
	"github.com/xneogo/extensions/xrand"
	"time"
)

func main() {
	fmt.Println("--------------------------------")
	// Example()
	fmt.Println("--------------------------------")
	ExampleRandString()
	fmt.Println("--------------------------------")
	// ExampleRandTicker()
	fmt.Println("--------------------------------")
	// ExampleJackpot()
	fmt.Println("--------------------------------")
}

func Example() {
	fmt.Println(xrand.ChooseMN(10, 5))
	fmt.Println(xrand.RandBetween(1, 10))
	fmt.Println(xrand.Wave(100, 10, 20))
}

func ExampleRandString() {
	fmt.Println(xrand.RandString(10))
	fmt.Println(xrand.RandDigit(10))
	fmt.Println(xrand.RandAnythingOnceFrom([]string{"a", "b", "c"}))
	fmt.Println(xrand.RandAnythingSomeFrom([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
	fmt.Println(xrand.RandAnythingSomeFrom([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
	fmt.Println(xrand.RandAnythingSomeFrom([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
}

func ExampleRandTicker() {
	cnt := 0
	tkr := xrand.NewRandTicker(time.Second, time.Second)
	for {
		if cnt > 10 {
			tkr.Stop()
			break
		}
		select {
		case <-tkr.C:
			cnt++
			fmt.Println("tick at ", time.Now().Unix())
		}
	}
}

func ExampleJackpot() {
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(100))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(1000))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(10000))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(100000))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(1000000))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(10000000))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(100000000))
	fmt.Println(xrand.RanD.CheckProbabilityJackpot(1000000000))
	fmt.Println()
}
