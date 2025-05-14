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
 @Time    : 2025/4/3 -- 14:17
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xphone examples/xphone/main.go
*/

package main

import (
	"context"
	"fmt"

	"github.com/xneogo/extensions/xphone"
)

func main() {
	// 示例1：验证中国大陆手机号
	fmt.Println("示例1：验证中国大陆手机号")
	verifyChinesePhone("86-19216882555")
	verifyChinesePhone("86-19116882555")
	verifyChinesePhone("86-18512818932")

	// 示例2：验证其他国家手机号
	fmt.Println("\n示例2：验证其他国家手机号")
	verifyInternationalPhone("852-51234567")  // 香港
	verifyInternationalPhone("1-4155552671")  // 美国
	verifyInternationalPhone("81-9012345678") // 日本

	// 示例3：手机号格式化
	fmt.Println("\n示例3：手机号格式化")
	formatPhone("+86-13812345678")
	formatPhone("0086-13812345678")
	formatPhone("13812345678")

	// 示例4：手机号匿名化
	fmt.Println("\n示例4：手机号匿名化")
	anonymizePhone("86-13812345678", false)
	anonymizePhone("86-13812345678", true)

	// 示例5：区号解析
	fmt.Println("\n示例5：区号解析")
	parsePhoneNumber("86-13812345678")
	parsePhoneNumber("+86-13812345678")
	parsePhoneNumber("13812345678")
}

// 验证中国大陆手机号
func verifyChinesePhone(input string) {
	phoneInfo, err := xphone.RegexpPhoneVerify(&xphone.PhoneVerifyReq{
		Phone:  input,
		Region: xphone.ChinaMainland,
	})
	if err != nil {
		fmt.Printf("验证失败 %s: %v\n", input, err)
		return
	}
	fmt.Printf("验证成功 %s: 区号=%s, 地区=%s, 标准格式=%s\n",
		input, phoneInfo.AreaNumber, phoneInfo.Region, phoneInfo.RegularPhone)
}

// 验证国际手机号
func verifyInternationalPhone(input string) {
	phoneInfo, err := xphone.RegexpPhoneVerify(&xphone.PhoneVerifyReq{
		Phone: input,
	})
	if err != nil {
		fmt.Printf("验证失败 %s: %v\n", input, err)
		return
	}
	fmt.Printf("验证成功 %s: 区号=%s, 地区=%s, 标准格式=%s\n",
		input, phoneInfo.AreaNumber, phoneInfo.Region, phoneInfo.RegularPhone)
}

// 格式化手机号
func formatPhone(input string) {
	// 标准格式化
	standard := xphone.GetStandardPhone(input, false)
	fmt.Printf("标准格式化 %s -> %s\n", input, standard)

	// 用于拨号的格式
	callFormat := xphone.FormatPhoneForCall(context.Background(), input, "+", false)
	fmt.Printf("拨号格式 %s -> %s\n", input, callFormat)
}

// 匿名化手机号
func anonymizePhone(input string, withoutArea bool) {
	anonymized := xphone.AnonymityPhone(input, withoutArea)
	fmt.Printf("匿名化 %s (withoutArea=%v) -> %s\n", input, withoutArea, anonymized)
}

// 解析手机号
func parsePhoneNumber(input string) {
	area, number := xphone.ParsePhone(input, true)
	fmt.Printf("解析 %s -> 区号=%s, 号码=%s\n", input, area, number)
}
