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
 @Time    : 2024/10/28 -- 18:26
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: xrate.go
*/

package xrate

import (
	"context"
	"errors"
)

var ErrRateLimiterNotFound = errors.New("[x-rate] rate limiter not found")
var ErrRateLimiterRegistryNotInited = errors.New("[x-rate] rate limit registry not inited! Call `Init()` first! ")
var ErrRateLimited error = errors.New("[x-rate] rate limited")

type RateLimiter interface {
	Limit() error
	ChangeQpsThreshold(newQpsThreshold int64)
}

type InterfaceRateLimitRegistry interface {
	InterfaceRateLimit(ctx context.Context, interfaceName string, caller string) error
}

// just make sure
var _ RateLimiter = (*LeakyBucketRateLimiter)(nil)
var _ RateLimiter = (*SlidingWindowRateLimiter)(nil)

const (
	RateLimitModeBlocking = 1
	RateLimitModeDenying  = 2

	// OtherCallerKey 其他调用方的key：表示除了指定调用方的其余所有调用方，都会被此规则统计
	OtherCallerKey = "OTHER_CALLER"
	// DefaultMaxCap 默认最大承载力key：表示一个服务中配置的接口默认最大承载力
	DefaultMaxCap = "DEFAULT_MAX_CAP"
	// MaxCap 接口最大承载力：表示当前接口最大承载力的key
	MaxCap = "NULL"
)
