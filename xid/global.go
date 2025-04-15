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
 @Time    : 2025/4/15 -- 13:57
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xid xid/global.go
*/

package xid

import "context"

var defaultIdConf = NewSnowFlakeModeConfig("_default_", "_default_")
var defaultIdGen, _ = CreateSnowFlakeIdGen(context.Background(), &defaultIdConf, 0)

func GetID(ctx context.Context) (int64, error) {
	return defaultIdGen.GetId(ctx)
}
