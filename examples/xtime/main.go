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
 @Time    : 2025/5/13 -- 14:17
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xtime examples/xtime/main.go
*/

package main

import (
	"fmt"
	"time"

	"github.com/xneogo/extensions/xtime"
)

func main() {
	// 示例1：基础时间操作
	fmt.Println("示例1：基础时间操作")
	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format(xtime.TimeFormatAll))
	fmt.Printf("本年开始: %s\n", xtime.StartOfYear(now).Format(xtime.TimeFormatAll))
	fmt.Printf("本年结束: %s\n", xtime.EndOfYear(now).Format(xtime.TimeFormatAll))
	fmt.Printf("本月开始: %s\n", xtime.StartOfMonth(now).Format(xtime.TimeFormatAll))
	fmt.Printf("本月结束: %s\n", xtime.EndOfMonth(now).Format(xtime.TimeFormatAll))
	fmt.Printf("本周开始: %s\n", xtime.StartOfWeek(now).Format(xtime.TimeFormatAll))
	fmt.Printf("本周结束: %s\n", xtime.EndOfWeek(now).Format(xtime.TimeFormatAll))
	fmt.Printf("当天开始: %s\n", xtime.StartOfDay(now).Format(xtime.TimeFormatAll))
	fmt.Printf("当天结束: %s\n", xtime.EndOfDay(now).Format(xtime.TimeFormatAll))

	// 示例2：时间戳操作
	fmt.Println("\n示例2：时间戳操作")
	nowStamp := time.Now().Unix()
	fmt.Printf("当前时间戳: %d\n", nowStamp)
	fmt.Printf("当天开始时间戳: %d\n", xtime.DayBeginStamp(nowStamp))
	fmt.Printf("当前小时开始时间戳: %d\n", xtime.HourBeginStamp(nowStamp))

	dayBegin, dayEnd := xtime.DayScope(nowStamp)
	fmt.Printf("当天时间范围: %d - %d\n", dayBegin, dayEnd)

	weekBegin, weekEnd := xtime.WeekScope(nowStamp)
	fmt.Printf("本周时间范围: %d - %d\n", weekBegin, weekEnd)

	monthBegin, monthEnd := xtime.MonthScope(nowStamp)
	fmt.Printf("本月时间范围: %d - %d\n", monthBegin, monthEnd)

	// 示例3：定时器功能
	fmt.Println("\n示例3：定时器功能")
	timer := xtime.NewTimer(2 * time.Second)
	timer.Start(func() {
		fmt.Printf("定时器触发: %s\n", time.Now().Format(xtime.TimeFormatAll))
	})

	// 等待5秒后停止定时器
	time.Sleep(5 * time.Second)
	timer.Stop()
	fmt.Println("定时器已停止")

	// 示例4：时间窗口检查
	fmt.Println("\n示例4：时间窗口检查")
	// 绝对时间检查
	needUpdate, err := xtime.CheckNeedUpdate(xtime.TimeTypeAbs, "", nowStamp)
	if err != nil {
		fmt.Printf("绝对时间检查错误: %v\n", err)
	} else {
		fmt.Printf("绝对时间检查结果: %v\n", needUpdate)
	}

	// 自然日检查
	needUpdate, err = xtime.CheckNeedUpdate(xtime.TimeTypeNatureDay, "", nowStamp)
	if err != nil {
		fmt.Printf("自然日检查错误: %v\n", err)
	} else {
		fmt.Printf("自然日检查结果: %v\n", needUpdate)
	}

	// 示例5：运行时间统计
	fmt.Println("\n示例5：运行时间统计")
	timeStat := xtime.NewTimeStat()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("运行时间: %d 毫秒\n", timeStat.Millisecond())
	fmt.Printf("运行时间: %d 微秒\n", timeStat.Microsecond())
	fmt.Printf("运行时间: %d 纳秒\n", timeStat.Nanosecond())
}
