# 队列工具 (xqueue)

## 简介

线程安全的内存队列，包含：
- FIFO队列：先进先出的消息队列
- 阻塞操作：支持阻塞式的Push和Pop操作
- 消息ID：自动生成唯一的消息标识
- 队列状态：支持队列大小查询和关闭操作
- 并发安全：完全线程安全的队列实现

## 使用示例

```go
import "github.com/xneogo/extensions/xqueue"

// 创建队列
queue := xqueue.NewQueue()

// 生产者
go func() {
    for i := 0; i < 5; i++ {
        err := queue.Push(fmt.Sprintf("消息-%d", i))
        if err != nil {
            fmt.Printf("推送失败: %v\n", err)
        } else {
            fmt.Printf("推送消息: 消息-%d\n", i)
        }
        time.Sleep(100 * time.Millisecond)
    }
}()

// 消费者
go func() {
    for {
        msg, err := queue.Pop()
        if err != nil {
            if err == xqueue.ErrQueueClosed {
                fmt.Println("队列已关闭")
                break
            }
            fmt.Printf("弹出失败: %v\n", err)
            continue
        }
        fmt.Printf("消费消息: %s (ID: %s)\n", msg.Body, msg.ID)
    }
}()

// 队列大小
size := queue.Size()
fmt.Printf("队列大小: %d\n", size)

// 关闭队列
time.Sleep(1 * time.Second)
queue.Close()
``` 