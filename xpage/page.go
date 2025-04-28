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
 @Time    : 2025/4/28 -- 18:03
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xpage xpage/page.go
*/

package xpage

import (
	"math"
)

func PageSize2OffsetLimit(page, size int) (offset, limit int) {
	if page <= 0 || size <= 0 {
		return 0, size
	}
	return size * (page - 1), size
}

func OffsetLimit2PageSize(offset, limit int) (page, size int) {
	if limit <= 0 || offset < 0 {
		return 1, 0
	}
	return int(math.Floor(float64(offset)/float64(limit))) + 1, limit
}
