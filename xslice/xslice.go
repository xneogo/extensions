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
 @Time    : 2024/7/15 -- 14:43
 @Author  : bishop ❤️ MONEY
 @Description: 对 go/src/slice 的补充 https://github.com/golang/go/blob/master/src/slices/slices.go
*/

package xslice

import (
	"github.com/xneogo/matrix/mgeneric"
	"math/rand"
	"slices"
	"time"
)

func Distinct[S ~[]E, E comparable](src S) (dst S) {
	tempMap := make(map[E]struct{})
	for _, v := range src {
		l := len(tempMap)
		tempMap[v] = struct{}{}
		if len(tempMap) != l {
			dst = append(dst, v)
		}
	}
	return
}

func DistinctFunc[S ~[]E, E any](src S, eq func(a, b E) bool) (dst S) {
	for _, v := range src {
		if slices.ContainsFunc(dst, func(e E) bool {
			return eq(v, e)
		}) {
			dst = append(dst, v)
		}
	}
	return
}

// Except 返回在 left 但是不在 right 中的 left 元素
func Except[S ~[]E, E comparable](left S, right S) (dst S) {
	tempMap := make(map[E]struct{}, len(right))
	for _, r := range right {
		tempMap[r] = struct{}{}
	}
	for _, l := range left {
		if _, ok := tempMap[l]; !ok {
			dst = append(dst, l)
		}
	}
	return
}

func ExceptFunc[S ~[]E, E any](left S, right S, eq func(a, b E) bool) (dst S) {
	// 在 left
	for _, v := range left {
		// 不在 right
		if !slices.ContainsFunc(right, func(e E) bool {
			return eq(v, e)
		}) {
			// 加入结果集
			dst = append(dst, v)
		}
	}
	return dst
}

func RPop[S ~[]E, E any](src S) (last E) {
	if len(src) == 0 {
		return
	}
	last = src[len(src)-1]
	src = src[:len(src)-1]
	return
}

func RPopN[S ~[]E, E any](src S, n int) S {
	res := make(S, n)
	for n > 0 {
		res = append(res, RPop(src))
		n--
	}
	return res
}

func LPop[S ~[]E, E any](src S) (first E) {
	if len(src) == 0 {
		return
	}
	first = src[0]
	src = src[1:]
	return
}

func LPopN[S ~[]E, E any](src S, n int) S {
	res := make(S, n)
	for n > 0 {
		res = append(res, LPop(src))
		n--
	}
	return res
}

func LPushN[S ~[]E, E any](src S, added S) S {
	return append(added, src...)
}

func RPushN[S ~[]E, E any](src S, added S) S {
	return append(src, added...)
}

func LPush[S ~[]E, E any](src S, added E) S {
	return append([]E{added}, src...)
}

func RPush[S ~[]E, E any](src S, added E) S {
	return append(src, added)
}

func Shuffle(n int, swap func(i, j int)) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if n < 0 {
		panic("invalid argument to Shuffle")
	}

	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// Shuffle really ought not be called with n that doesn't fit in 32 bits.
	// Not only will it take a very long time, but with 2³¹! possible permutations,
	// there's no way that any PRNG can have a big enough internal state to
	// generate even a minuscule percentage of the possible permutations.
	// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

// Cut 左开右闭
func Cut[S ~[]E, E any](src S, start, end int) S {
	if start < 0 {
		start = 0
	}
	if end > len(src) {
		end = len(src)
	}
	if !RangeSafe(len(src), start, end) {
		return make(S, 0)
	}
	return src[start:end]
}

func RangeSafe(length, start, end int) bool {
	return length > start && length >= end && start < end
}

func Sum[Elem mgeneric.Number](src []Elem) Elem {
	var sum Elem
	for _, s := range src {
		sum += s
	}
	return sum
}
