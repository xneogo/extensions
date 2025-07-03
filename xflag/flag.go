package xflag

import (
	"fmt"
	"strings"
)

type ArrayStringFlags []string

// Value ...
func (i *ArrayStringFlags) String() string {
	return fmt.Sprint(*i)
}

// Set 方法是flag.Value接口, 设置flag Value的方法.
// 通过多个flag指定的值， 所以我们追加到最终的数组上.
func (i *ArrayStringFlags) Set(value string) error {
	arr := strings.Split(value, ",")
	for _, s := range arr {
		s = strings.TrimSpace(s)
		if s != "" {
			*i = append(*i, strings.TrimSpace(s))
		}
	}
	return nil
}
