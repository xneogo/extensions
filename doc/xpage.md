# 分页工具 (xpage)

## 简介

简单实用的分页计算工具，支持：
- 页码转换：页码和偏移量的双向转换
- 数据库分页：直接适用于SQL LIMIT查询
- 边界处理：自动处理边界条件和异常情况
- 灵活配置：支持不同的分页大小设置

## 使用示例

```go
import "github.com/xneogo/extensions/xpage"

// 页码转换为偏移量和限制
page := 2
size := 10
offset, limit := xpage.PageSize2OffsetLimit(page, size)
fmt.Printf("第%d页，每页%d条 -> 偏移量: %d, 限制: %d\n", page, size, offset, limit)

// 偏移量转换为页码
offset = 20
limit = 10
page, size = xpage.OffsetLimit2PageSize(offset, limit)
fmt.Printf("偏移量: %d, 限制: %d -> 第%d页，每页%d条\n", offset, limit, page, size)

// 实际使用示例
func getUsers(page, size int) {
    offset, limit := xpage.PageSize2OffsetLimit(page, size)
    // 使用offset和limit进行数据库查询
    // SELECT * FROM users LIMIT offset, limit
    fmt.Printf("查询用户: LIMIT %d, %d\n", offset, limit)
}
``` 