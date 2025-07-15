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
 @Time    : 2025/7/11 -- 13:36
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: syncxmap xsync/syncxmap/ttl.go
*/

package syncxmap

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// TTLMapSegment 分段存储结构
type TTLMapSegment struct {
	m             sync.Map
	mu            sync.RWMutex
	itemCount     int64
	memorySize    int64 // 估算内存使用
	lastCleanTime time.Time
}

// TTLMap TTLMap
type TTLMap struct {
	segments    []*TTLMapSegment
	segmentMask int // 分段掩码，用于快速hash
	config      *Config
	stopCh      chan struct{}
	cleanTicker *time.Ticker // 定期清理定时器
	closed      int32        // 关闭标志
}

// item 企业级存储项结构
type item struct {
	value       interface{}
	expiry      time.Time
	lastAccess  time.Time
	createTime  time.Time
	accessCount int64
}

// NewTTLMap 创建新的企业级TTLMap
func NewTTLMap(config *Config) (*TTLMap, error) {
	if config == nil {
		config = DefaultConfig()
	}

	// 验证配置
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	// 确保分段数量是2的幂，便于快速hash
	segmentCount := nextPowerOf2(config.SegmentCount)
	if segmentCount < 2 {
		segmentCount = 2
	}

	em := &TTLMap{
		segments:    make([]*TTLMapSegment, segmentCount),
		segmentMask: segmentCount - 1,
		config:      config,
		stopCh:      make(chan struct{}),
	}

	// 初始化所有分段
	for i := 0; i < segmentCount; i++ {
		em.segments[i] = &TTLMapSegment{
			lastCleanTime: time.Now(),
		}
	}

	// 启动后台清理任务
	go em.backgroundClean()

	return em, nil
}

// nextPowerOf2 返回大于等于n的最小2的幂
func nextPowerOf2(n int) int {
	if n <= 1 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}

// getSegmentIndex 根据key计算分段索引
func (em *TTLMap) getSegmentIndex(key interface{}) int {
	return int(hashKey(key)) & em.segmentMask
}

// hashKey 计算key的hash值
func hashKey(key interface{}) uint32 {
	switch k := key.(type) {
	case string:
		return stringHash(k)
	case int:
		return uint32(k)
	case int32:
		return uint32(k)
	case int64:
		return uint32(k) ^ uint32(k>>32)
	case uint32:
		return k
	case uint64:
		return uint32(k) ^ uint32(k>>32)
	default:
		// 对于其他类型，使用简单的字符串转换
		return stringHash(key.(string))
	}
}

// stringHash 计算字符串hash值 (FNV-1a)
func stringHash(s string) uint32 {
	const (
		offset32 = 2166136261
		prime32  = 16777619
	)
	hash := uint32(offset32)
	for i := 0; i < len(s); i++ {
		hash ^= uint32(s[i])
		hash *= prime32
	}
	return hash
}

// estimateItemSize 估算条目的内存大小
func (em *TTLMap) estimateItemSize(key, value interface{}) int64 {
	keySize := estimateSize(key)
	valueSize := estimateSize(value)
	// 加上item结构体的固定开销
	return keySize + valueSize + 64 // 64字节为item结构体和时间戳的估算开销
}

// estimateSize 估算对象的内存大小
func estimateSize(obj interface{}) int64 {
	switch v := obj.(type) {
	case string:
		return int64(len(v)) + 24 // 字符串头部开销
	case []byte:
		return int64(len(v)) + 24
	case int, int32, int64, uint, uint32, uint64:
		return 8
	case float32, float64:
		return 8
	case bool:
		return 1
	default:
		// 对于其他类型，使用保守估算
		return 64
	}
}

// backgroundClean 后台清理goroutine
func (em *TTLMap) backgroundClean() {
	em.cleanTicker = time.NewTicker(em.config.CleanInterval)
	defer em.cleanTicker.Stop()

	for {
		select {
		case <-em.cleanTicker.C:
			em.performCleanup()
		case <-em.stopCh:
			return
		}
	}
}

// performCleanup 执行清理任务
func (em *TTLMap) performCleanup() {
	if atomic.LoadInt32(&em.closed) == 1 {
		return
	}

	// 智能清理策略
	triggers := em.evaluateCleanupTriggers()

	if !triggers.needsCleanup() {
		return
	}

	// 批处理清理每个分段
	for _, segment := range em.segments {
		em.batchCleanSegment(segment, triggers)
	}

	// 如果内存压力过大，执行内存保护
	if triggers.memoryPressure {
		em.enforceMemoryLimit()
	}
}

// cleanupTriggers 清理触发条件
type cleanupTriggers struct {
	memoryPressure bool // 内存压力触发
	itemCountLimit bool // 条目数限制触发
	timeInterval   bool // 时间间隔触发
	segmentLoad    bool // 分段负载触发
}

func (ct *cleanupTriggers) needsCleanup() bool {
	return ct.memoryPressure || ct.itemCountLimit || ct.timeInterval || ct.segmentLoad
}

// evaluateCleanupTriggers 评估清理触发条件
func (em *TTLMap) evaluateCleanupTriggers() *cleanupTriggers {
	triggers := &cleanupTriggers{}

	// 检查内存压力
	triggers.memoryPressure = em.isMemoryPressure()

	// 检查条目数限制
	totalItems := em.getTotalItemCount()
	triggers.itemCountLimit = totalItems > int64(float64(em.config.MaxItemCount)*0.8)

	// 检查时间间隔
	triggers.timeInterval = em.shouldCleanByTime()

	// 检查分段负载
	triggers.segmentLoad = em.hasHighLoadSegments()

	return triggers
}

// shouldCleanByTime 检查是否应该按时间清理
func (em *TTLMap) shouldCleanByTime() bool {
	for _, segment := range em.segments {
		if time.Since(segment.lastCleanTime) > em.config.CleanInterval {
			return true
		}
	}
	return false
}

// hasHighLoadSegments 检查是否有高负载分段
func (em *TTLMap) hasHighLoadSegments() bool {
	avgItemsPerSegment := em.getTotalItemCount() / int64(len(em.segments))
	threshold := avgItemsPerSegment * 2 // 2倍平均值作为阈值

	for _, segment := range em.segments {
		if atomic.LoadInt64(&segment.itemCount) > threshold {
			return true
		}
	}
	return false
}

// batchCleanSegment 批处理清理分段
func (em *TTLMap) batchCleanSegment(segment *TTLMapSegment, triggers *cleanupTriggers) {
	segment.mu.Lock()
	defer segment.mu.Unlock()

	// 跳过最近清理过的分段（除非有内存压力）
	if !triggers.memoryPressure && time.Since(segment.lastCleanTime) < em.config.CleanInterval/2 {
		return
	}

	now := time.Now()
	processed := 0
	batchSize := em.config.CleanBatchSize

	// 批处理清理
	segment.m.Range(func(key, value interface{}) bool {
		if processed >= batchSize {
			return false // 停止遍历，下次继续
		}

		item := value.(*item)
		shouldDelete := false

		// 检查过期
		if now.After(item.expiry) {
			shouldDelete = true
		}

		// 在内存压力下，检查是否应该LRU淘汰
		if !shouldDelete && triggers.memoryPressure && em.config.EnableLRU {
			// 淘汰超过一定时间未访问的数据
			if time.Since(item.lastAccess) > em.config.CleanInterval*2 {
				shouldDelete = true
			}
		}

		if shouldDelete {
			segment.m.Delete(key)
			itemSize := em.estimateItemSize(key, item.value)
			atomic.AddInt64(&segment.itemCount, -1)
			atomic.AddInt64(&segment.memorySize, -itemSize)
		}

		processed++
		return true
	})

	// 更新清理时间
	segment.lastCleanTime = now
}

// Store 存储带TTL的键值对（企业级版本）
func (em *TTLMap) Store(key, value interface{}, ttl time.Duration) error {
	if atomic.LoadInt32(&em.closed) == 1 {
		return ErrMapClosed
	}

	// 检查key和value大小限制
	if err := em.checkSizeLimits(key, value); err != nil {
		return err
	}

	// 获取分段
	segmentIndex := em.getSegmentIndex(key)
	segment := em.segments[segmentIndex]

	// 估算内存大小
	itemSize := em.estimateItemSize(key, value)

	// 检查内存和条目数限制
	if err := em.checkLimits(segment, itemSize); err != nil {
		return err
	}

	now := time.Now()
	newItem := &item{
		value:       value,
		expiry:      now.Add(ttl),
		lastAccess:  now,
		createTime:  now,
		accessCount: 1,
	}

	// 检查是否是新key还是更新已存在的key
	_, existed := segment.m.LoadOrStore(key, newItem)

	segment.mu.Lock()
	if !existed {
		// 新key，增加计数和内存使用
		atomic.AddInt64(&segment.itemCount, 1)
		atomic.AddInt64(&segment.memorySize, itemSize)
	} else {
		// 已存在的key，直接更新
		segment.m.Store(key, newItem)
		// 内存使用可能有变化，这里简化处理
	}
	segment.mu.Unlock()

	return nil
}

// checkSizeLimits 检查key和value的大小限制
func (em *TTLMap) checkSizeLimits(key, value interface{}) error {
	keySize := estimateSize(key)
	valueSize := estimateSize(value)

	if keySize > int64(em.config.MaxKeySize) {
		return ErrKeySizeLimit
	}

	if valueSize > int64(em.config.MaxValueSize) {
		return ErrValueSizeLimit
	}

	return nil
}

// checkLimits 检查内存和条目数限制
func (em *TTLMap) checkLimits(segment *TTLMapSegment, itemSize int64) error {
	// 检查条目数限制
	totalItems := em.getTotalItemCount()
	if totalItems >= em.config.MaxItemCount {
		return ErrItemCountLimit
	}

	// 检查内存限制
	totalMemoryMB := em.getTotalMemoryMB()
	if totalMemoryMB+(itemSize/1024/1024) > em.config.MaxMemoryMB {
		return ErrMemoryLimit
	}

	return nil
}

// getTotalItemCount 获取总条目数
func (em *TTLMap) getTotalItemCount() int64 {
	var total int64
	for _, segment := range em.segments {
		total += atomic.LoadInt64(&segment.itemCount)
	}
	return total
}

// getTotalMemoryMB 获取总内存使用量(MB)
func (em *TTLMap) getTotalMemoryMB() int64 {
	var total int64
	for _, segment := range em.segments {
		total += atomic.LoadInt64(&segment.memorySize)
	}
	return total / 1024 / 1024
}

// Load 加载指定key的值（企业级版本）
func (em *TTLMap) Load(key interface{}) (interface{}, bool) {
	if atomic.LoadInt32(&em.closed) == 1 {
		return nil, false
	}

	// 获取分段
	segmentIndex := em.getSegmentIndex(key)
	segment := em.segments[segmentIndex]

	value, ok := segment.m.Load(key)
	if !ok {
		return nil, false
	}

	item := value.(*item)
	now := time.Now()

	// 检查是否过期
	if now.After(item.expiry) {
		// 惰性删除过期项
		segment.m.Delete(key)
		segment.mu.Lock()
		atomic.AddInt64(&segment.itemCount, -1)
		atomic.AddInt64(&segment.memorySize, -em.estimateItemSize(key, item.value))
		segment.mu.Unlock()
		return nil, false
	}

	// 更新LRU信息
	item.lastAccess = now
	atomic.AddInt64(&item.accessCount, 1)

	return item.value, true
}

// LoadOrStore 如果key存在则返回现有值，否则存储并返回给定值（企业级版本）
func (em *TTLMap) LoadOrStore(key, value interface{}, ttl time.Duration) (interface{}, bool, error) {
	if atomic.LoadInt32(&em.closed) == 1 {
		return nil, false, ErrMapClosed
	}

	// 先尝试加载
	if actual, loaded := em.Load(key); loaded {
		return actual, true, nil
	}

	// 不存在则存储
	err := em.Store(key, value, ttl)
	if err != nil {
		return nil, false, err
	}

	return value, false, nil
}

// Delete 删除指定key的值（企业级版本）
func (em *TTLMap) Delete(key interface{}) bool {
	if atomic.LoadInt32(&em.closed) == 1 {
		return false
	}

	// 获取分段
	segmentIndex := em.getSegmentIndex(key)
	segment := em.segments[segmentIndex]

	value, existed := segment.m.LoadAndDelete(key)
	if existed {
		item := value.(*item)
		segment.mu.Lock()
		atomic.AddInt64(&segment.itemCount, -1)
		atomic.AddInt64(&segment.memorySize, -em.estimateItemSize(key, item.value))
		segment.mu.Unlock()
		return true
	}

	return false
}

// Len 返回当前有效数据量的估算值（企业级版本）
func (em *TTLMap) Len() int {
	if atomic.LoadInt32(&em.closed) == 1 {
		return 0
	}

	return int(em.getTotalItemCount())
}

// enforceMemoryLimit 执行内存限制保护
func (em *TTLMap) enforceMemoryLimit() {
	if atomic.LoadInt32(&em.closed) == 1 {
		return
	}

	totalMemoryMB := em.getTotalMemoryMB()
	if totalMemoryMB <= em.config.MaxMemoryMB {
		return
	}

	// 内存超限，需要清理
	targetMemoryMB := int64(float64(em.config.MaxMemoryMB) * 0.8) // 清理到80%
	freedMemoryMB := int64(0)

	// 先清理过期数据
	for _, segment := range em.segments {
		if freedMemoryMB >= (totalMemoryMB - targetMemoryMB) {
			break
		}
		freedMemoryMB += em.cleanExpiredInSegment(segment)
	}

	// 如果还需要更多空间，进行LRU淘汰
	if freedMemoryMB < (totalMemoryMB-targetMemoryMB) && em.config.EnableLRU {
		for _, segment := range em.segments {
			if freedMemoryMB >= (totalMemoryMB - targetMemoryMB) {
				break
			}
			freedMemoryMB += em.performLRUEviction(segment, (totalMemoryMB - targetMemoryMB - freedMemoryMB))
		}
	}
}

// cleanExpiredInSegment 清理分段中的过期数据
func (em *TTLMap) cleanExpiredInSegment(segment *TTLMapSegment) int64 {
	var freedMemory int64
	now := time.Now()

	segment.m.Range(func(key, value interface{}) bool {
		item := value.(*item)
		if now.After(item.expiry) {
			segment.m.Delete(key)
			itemSize := em.estimateItemSize(key, item.value)
			atomic.AddInt64(&segment.itemCount, -1)
			atomic.AddInt64(&segment.memorySize, -itemSize)
			freedMemory += itemSize
		}
		return true
	})

	return freedMemory / 1024 / 1024 // 转换为MB
}

// performLRUEviction 执行LRU淘汰
func (em *TTLMap) performLRUEviction(segment *TTLMapSegment, targetFreeMB int64) int64 {
	type keyWithTime struct {
		key        interface{}
		lastAccess time.Time
		size       int64
	}

	var candidates []keyWithTime

	// 收集候选项
	segment.m.Range(func(key, value interface{}) bool {
		item := value.(*item)
		itemSize := em.estimateItemSize(key, item.value)
		candidates = append(candidates, keyWithTime{
			key:        key,
			lastAccess: item.lastAccess,
			size:       itemSize,
		})
		return true
	})

	// 按最后访问时间排序（最老的在前）
	for i := 0; i < len(candidates)-1; i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[i].lastAccess.After(candidates[j].lastAccess) {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}

	// 淘汰最老的数据
	var freedMemory int64
	evictCount := int(float64(len(candidates)) * em.config.LRUEvictPercent)
	if evictCount == 0 {
		evictCount = 1
	}

	for i := 0; i < evictCount && i < len(candidates); i++ {
		candidate := candidates[i]
		segment.m.Delete(candidate.key)
		atomic.AddInt64(&segment.itemCount, -1)
		atomic.AddInt64(&segment.memorySize, -candidate.size)
		freedMemory += candidate.size

		if freedMemory/1024/1024 >= targetFreeMB {
			break
		}
	}

	return freedMemory / 1024 / 1024 // 转换为MB
}

// isMemoryPressure 检查是否存在内存压力
func (em *TTLMap) isMemoryPressure() bool {
	totalMemoryMB := em.getTotalMemoryMB()
	return totalMemoryMB > int64(float64(em.config.MaxMemoryMB)*0.9) // 90%阈值
}

// Stop 停止企业级TTLMap，清理资源
func (em *TTLMap) Stop() {
	if !atomic.CompareAndSwapInt32(&em.closed, 0, 1) {
		return // 已经关闭
	}

	// 关闭后台清理goroutine
	if em.stopCh != nil {
		close(em.stopCh)
	}

	// 停止定期清理定时器
	if em.cleanTicker != nil {
		em.cleanTicker.Stop()
	}

	// 清理所有分段数据（可选）
	for _, segment := range em.segments {
		segment.mu.Lock()
		// 清空数据
		segment.m.Range(func(key, value interface{}) bool {
			segment.m.Delete(key)
			return true
		})
		atomic.StoreInt64(&segment.itemCount, 0)
		atomic.StoreInt64(&segment.memorySize, 0)
		segment.mu.Unlock()
	}
}

// IsClosed 检查是否已关闭
func (em *TTLMap) IsClosed() bool {
	return atomic.LoadInt32(&em.closed) == 1
}

// Stats 获取统计信息
func (em *TTLMap) Stats() map[string]interface{} {
	if atomic.LoadInt32(&em.closed) == 1 {
		return map[string]interface{}{
			"status": "closed",
		}
	}

	stats := make(map[string]interface{})
	stats["total_items"] = em.getTotalItemCount()
	stats["total_memory_mb"] = em.getTotalMemoryMB()
	stats["segment_count"] = len(em.segments)
	stats["memory_limit_mb"] = em.config.MaxMemoryMB
	stats["item_count_limit"] = em.config.MaxItemCount
	stats["status"] = "running"

	// 分段统计
	segmentStats := make([]map[string]interface{}, len(em.segments))
	for i, segment := range em.segments {
		segmentStats[i] = map[string]interface{}{
			"items":      atomic.LoadInt64(&segment.itemCount),
			"memory_mb":  atomic.LoadInt64(&segment.memorySize) / 1024 / 1024,
			"last_clean": segment.lastCleanTime.Format("2006-01-02 15:04:05"),
		}
	}
	stats["segments"] = segmentStats

	return stats
}

// validateConfig 验证配置
func validateConfig(config *Config) error {
	if config.SegmentCount <= 0 {
		return errors.New("segment count must be positive")
	}
	if config.MaxMemoryMB <= 0 {
		return errors.New("max memory must be positive")
	}
	if config.MaxItemCount <= 0 {
		return errors.New("max item count must be positive")
	}
	if config.CleanInterval <= 0 {
		return errors.New("clean interval must be positive")
	}
	if config.CleanBatchSize <= 0 {
		return errors.New("clean batch size must be positive")
	}
	if config.LRUEvictPercent <= 0 || config.LRUEvictPercent > 1 {
		return errors.New("LRU evict percent must be between 0 and 1")
	}
	return nil
}
