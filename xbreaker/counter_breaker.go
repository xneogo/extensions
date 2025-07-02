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
 @Time    : 2024/11/4 -- 10:40
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: counter_breaker.go
*/

package xbreaker

import (
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xneogo/extensions/xstring"
)

const (
	checkTick          = time.Millisecond * 25
	defaultGranularity = time.Second * 1
	defaultThreshold   = 10
	defaultBreakerGap  = 10 // 单位: seconds
)

// CounterBreakerManager todo 简单计数法实现熔断操作，后续改为滑动窗口或三方组件的方式
type CounterBreakerManager struct {
	lock     sync.Mutex
	Breakers map[string]*Breaker
}

var cm *CounterBreakerManager

func init() {
	cm = &CounterBreakerManager{Breakers: make(map[string]*Breaker)}
}

func StatBreaker(cluster, table string, err error) {
	cm.StatBreaker(cluster, table, err)
}

func Entry(cluster, table string) bool {
	return cm.Entry(cluster, table)
}

// StatBreaker state errors for breaker
func (b *CounterBreakerManager) StatBreaker(cluster, table string, err error) {
	if err != nil && (strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "invalid connection")) {
		key := xstring.Concat(cluster, "_", table)
		cm.lock.Lock()
		if _, ok := cm.Breakers[key]; !ok {
			breaker := new(Breaker)
			breaker.run()
			cm.Breakers[key] = breaker
		}
		breaker := cm.Breakers[key]
		cm.lock.Unlock()
		atomic.AddInt32(&breaker.Count, 1)
	}
}

// Entry check if allow request
func (b *CounterBreakerManager) Entry(cluster, table string, ) bool {
	key := xstring.Concat(cluster, "_", table)
	cm.lock.Lock()
	breaker := cm.Breakers[key]
	cm.lock.Unlock()
	if breaker != nil {
		return atomic.LoadInt32(&breaker.Rejected) != 1
	}
	return true
}

type Breaker struct {
	Rejected      int32
	RejectedStart int64
	Count         int32
}

func (b *Breaker) run() {
	go func() {
		granularityTickC := time.Tick(defaultGranularity)
		checkTickC := time.Tick(checkTick)
		for {
			select {
			case <-granularityTickC:
				atomic.StoreInt32(&b.Count, 0)
				// check 1s/checkTick times in 1s
			case <-checkTickC:
				threshold := defaultThreshold
				breakerGap := defaultBreakerGap
				if atomic.LoadInt32(&b.Count) > int32(threshold) {
					atomic.StoreInt32(&b.Rejected, 1)
					b.RejectedStart = time.Now().Unix()
				} else {
					now := time.Now().Unix()
					if now-b.RejectedStart > int64(breakerGap) {
						atomic.StoreInt32(&b.Rejected, 0)
					}
				}
			}
		}
	}()
}
