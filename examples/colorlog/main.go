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
 @Time    : 2025/7/3 -- 11:25
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: colorlog examples/colorlog/main.go
*/

package colorlog

import (
	"context"
	"log"

	"github.com/xneogo/extensions/colorlog"
)

func main() {
	// Cyan, "青色 debug", Reset, CyanBright, "高亮 debug", Reset, "恢复默认颜色", CyanDelLine, "删除线", Reset, CyanUnderLine, "下划线", Reset, CyanBevel, "斜体 debug", Reset, CyanBg, "背景", Reset
	// default log
	colorlog.SetLevel(colorlog.PanicLevel)
	colorlog.SetColorful(true)
	ctx := context.Background()

	colorlog.Debug(ctx, "Debug 日志")
	colorlog.Info(ctx, "Info 日志")
	colorlog.Warn(ctx, "Warn 日志")
	colorlog.Error(ctx, "Error 日志")
	colorlog.Panic(ctx, "Panic 日志")

	log.Println("================")
	colorlog.Debugf(ctx, "Debugf %s, %s, %d", "日志", "asdasda", 4)
	colorlog.Infof(ctx, "Infof %s", "日志")
	colorlog.Warnf(ctx, "Warnf %s", "日志")
	colorlog.Errorf(ctx, "Errorf %s", "日志")
	colorlog.Panicf(ctx, "Panicf %s", "日志")

	log.Println("================")

	xLogger := colorlog.New()
	// xLogger.Color(Red)
	xLogger.SetLevel(colorlog.WarnLevel)

	xLogger.Debug(ctx, "new logger Debug 日志")
	xLogger.Info(ctx, "new logger Info 日志")
	xLogger.Warn(ctx, "new logger Warn 日志")
	xLogger.Error(ctx, "new logger Error 日志")

	log.Println("================")

	xLogger.Debugf(ctx, "new logger Debugf %s", "日志")
	xLogger.Infof(ctx, "new logger Infof %s", "日志")
	xLogger.Warnf(ctx, "new logger Warnf %s", "日志")
	xLogger.Errorf(ctx, "new logger Errorf %s", "日志")

	// color log
}
