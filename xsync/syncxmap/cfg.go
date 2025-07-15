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
 @Time    : 2025/7/15 -- 12:26
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: syncxmap xsync/syncxmap/cfg.go
*/

package syncxmap

import (
	"errors"
	"time"
)

// 企业级TTLMap错误定义
var (
	ErrMemoryLimit    = errors.New("memory limit exceeded")
	ErrItemCountLimit = errors.New("item count limit exceeded")
	ErrKeySizeLimit   = errors.New("key size limit exceeded")
	ErrValueSizeLimit = errors.New("value size limit exceeded")
	ErrMapClosed      = errors.New("map is closed")
)

// Config 企业级TTLMap配置
type Config struct {
	// 性能配置
	SegmentCount   int           // 分段数量，默认32
	CleanInterval  time.Duration // 清理间隔
	CleanBatchSize int           // 批处理清理大小

	// 安全限制
	MaxMemoryMB  int64 // 最大内存限制(MB)
	MaxItemCount int64 // 最大条目数限制
	MaxKeySize   int   // 单个key最大长度
	MaxValueSize int   // 单个value最大长度

	// 清理策略
	EnableLRU       bool    // 是否启用LRU淘汰
	LRUEvictPercent float64 // LRU淘汰比例，默认0.1(10%)
}

// DefaultConfig 默认企业级配置
func DefaultConfig() *Config {
	return &Config{
		SegmentCount:    32,
		CleanInterval:   5 * time.Minute,
		CleanBatchSize:  100,
		MaxMemoryMB:     1024,        // 1GB
		MaxItemCount:    1000000,     // 100万条目
		MaxKeySize:      1024,        // 1KB
		MaxValueSize:    1024 * 1024, // 1MB
		EnableLRU:       true,
		LRUEvictPercent: 0.1,
	}
}

func HighConcurrencyConfig() *Config {
	return &Config{
		SegmentCount:    64,
		CleanInterval:   30 * time.Second,
		CleanBatchSize:  200,
		MaxMemoryMB:     2048,
		MaxItemCount:    10000000,
		MaxKeySize:      2048,
		MaxValueSize:    10 * 1024 * 1024,
		EnableLRU:       true,
		LRUEvictPercent: 0.1,
	}
}

func LowLatencyConfig() *Config {
	return &Config{
		SegmentCount:    128,
		CleanInterval:   5 * time.Second,
		CleanBatchSize:  20,
		MaxMemoryMB:     1024,
		MaxItemCount:    1000000,
		MaxKeySize:      512,
		MaxValueSize:    4096,
		EnableLRU:       false, // 禁用LRU以减少延迟
		LRUEvictPercent: 0.1,
	}
}

func MemoryConstrainedConfig() *Config {
	return &Config{
		SegmentCount:    16,
		CleanInterval:   10 * time.Second,
		CleanBatchSize:  50,
		MaxMemoryMB:     256,
		MaxItemCount:    100000,
		MaxKeySize:      256,
		MaxValueSize:    1024,
		EnableLRU:       true,
		LRUEvictPercent: 0.2,
	}
}
