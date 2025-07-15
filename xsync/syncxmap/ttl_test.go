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
 @Time    : 2025/7/11 -- 13:37
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xsync xsync/map_test.go
*/

package syncxmap

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 企业级TTLMap测试
func TestTTLMapBasicOperations(t *testing.T) {
	em, err := NewTTLMap(nil)
	if err != nil {
		t.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	// 测试Store和Load
	err = em.Store("key1", "value1", 1*time.Second)
	if err != nil {
		t.Errorf("Failed to store: %v", err)
	}

	if value, ok := em.Load("key1"); !ok || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	// 测试Len
	if em.Len() != 1 {
		t.Errorf("Expected length 1, got %d", em.Len())
	}

	// 等待过期
	time.Sleep(1100 * time.Millisecond)

	// 测试惰性清理
	if _, ok := em.Load("key1"); ok {
		t.Error("Expected key1 to be expired and deleted")
	}

	if em.Len() != 0 {
		t.Errorf("Expected length 0 after expiry, got %d", em.Len())
	}
}

func TestTTLMapLimits(t *testing.T) {
	config := &Config{
		SegmentCount:    4,
		CleanInterval:   1 * time.Second,
		CleanBatchSize:  10,
		MaxMemoryMB:     1,   // 1MB限制
		MaxItemCount:    100, // 100条目限制
		MaxKeySize:      50,  // 50字节key限制（考虑string头部开销）
		MaxValueSize:    100, // 100字节value限制
		EnableLRU:       true,
		LRUEvictPercent: 0.1,
	}

	em, err := NewTTLMap(config)
	if err != nil {
		t.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	// 测试key大小限制
	longKey := "this_is_a_very_long_key_that_should_definitely_exceed_the_limit_of_fifty_characters"
	err = em.Store(longKey, "value", 1*time.Hour)
	if err != ErrKeySizeLimit {
		t.Errorf("Expected ErrKeySizeLimit, got %v", err)
	}

	// 测试value大小限制
	largeValue := make([]byte, 200)
	err = em.Store("key", largeValue, 1*time.Hour)
	if err != ErrValueSizeLimit {
		t.Errorf("Expected ErrValueSizeLimit, got %v", err)
	}

	// 测试条目数限制
	for i := 0; i < 150; i++ {
		key := fmt.Sprintf("k%d", i) // 短key，应该不会触发key大小限制
		err = em.Store(key, "value", 1*time.Hour)
		if err != nil && err != ErrItemCountLimit {
			t.Errorf("Unexpected error at iteration %d: %v", i, err)
		}
	}
}

func TestTTLMapConcurrency(t *testing.T) {
	em, err := NewTTLMap(nil)
	if err != nil {
		t.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	var wg sync.WaitGroup
	numGoroutines := 50
	numOperations := 100

	// 并发写入
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				value := fmt.Sprintf("value_%d_%d", id, j)
				em.Store(key, value, 500*time.Millisecond)
			}
		}(i)
	}

	// 并发读取
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				em.Load(key)
			}
		}(i)
	}

	// 并发删除
	for i := 0; i < numGoroutines/2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations/2; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)
				em.Delete(key)
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Final length: %d", em.Len())
}

func TestTTLMapCleanup(t *testing.T) {
	config := &Config{
		SegmentCount:    4,
		CleanInterval:   100 * time.Millisecond,
		CleanBatchSize:  10,
		MaxMemoryMB:     1024,
		MaxItemCount:    1000,
		MaxKeySize:      1024,
		MaxValueSize:    1024,
		EnableLRU:       true,
		LRUEvictPercent: 0.1,
	}

	em, err := NewTTLMap(config)
	if err != nil {
		t.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	// 添加一些会很快过期的数据
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("short_key_%d", i)
		em.Store(key, "value", 50*time.Millisecond)
	}

	// 添加一些长期数据
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("long_key_%d", i)
		em.Store(key, "value", 1*time.Hour)
	}

	initialLen := em.Len()
	t.Logf("Initial length: %d", initialLen)

	// 等待清理
	time.Sleep(300 * time.Millisecond)

	finalLen := em.Len()
	t.Logf("Final length after cleanup: %d", finalLen)

	if finalLen >= initialLen {
		t.Error("Expected cleanup to reduce the map size")
	}
}

func TestTTLMapStats(t *testing.T) {
	em, err := NewTTLMap(nil)
	if err != nil {
		t.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	// 添加一些数据
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		em.Store(key, "value", 1*time.Hour)
	}

	stats := em.Stats()
	t.Logf("Stats: %+v", stats)

	if stats["total_items"] != int64(10) {
		t.Errorf("Expected 10 items, got %v", stats["total_items"])
	}

	if stats["status"] != "running" {
		t.Errorf("Expected running status, got %v", stats["status"])
	}
}

func TestTTLMapStop(t *testing.T) {
	em, err := NewTTLMap(nil)
	if err != nil {
		t.Fatalf("Failed to create TTLMap: %v", err)
	}

	// 添加一些数据
	em.Store("key1", "value1", 1*time.Hour)

	// 停止
	em.Stop()

	// 验证已关闭
	if !em.IsClosed() {
		t.Error("Expected map to be closed")
	}

	// 尝试操作已关闭的map
	err = em.Store("key2", "value2", 1*time.Hour)
	if err != ErrMapClosed {
		t.Errorf("Expected ErrMapClosed, got %v", err)
	}

	if _, ok := em.Load("key1"); ok {
		t.Error("Expected Load to fail on closed map")
	}
}

// 性能基准测试
func BenchmarkTTLMapStore(b *testing.B) {
	em, err := NewTTLMap(nil)
	if err != nil {
		b.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key_%d", i)
			em.Store(key, "value", 1*time.Hour)
			i++
		}
	})
}

func BenchmarkTTLMapMixed(b *testing.B) {
	em, err := NewTTLMap(nil)
	if err != nil {
		b.Fatalf("Failed to create TTLMap: %v", err)
	}
	defer em.Stop()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key_%d", i)
			if i%3 == 0 {
				em.Store(key, "value", 1*time.Hour)
			} else if i%3 == 1 {
				em.Load(key)
			} else {
				em.Delete(key)
			}
			i++
		}
	})
}
