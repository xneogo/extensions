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
 @Time    : 2024/11/4 -- 18:25
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: device_no.go
*/

package xother

import (
	"strings"
)

func MacroValid(s string) bool {
	if s == "" ||
		strings.EqualFold(s, "null") ||
		strings.EqualFold(s, "nil") ||
		(strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")) ||
		(strings.HasPrefix(s, "_") && strings.HasSuffix(s, "_")) ||
		(strings.HasPrefix(s, "$") && strings.HasSuffix(s, "$")) {
		return false
	}
	return true
}

func DiDValid(no string) bool {
	if no == "" {
		return false
	}
	if no == "NULL" || no == "null" || no == "nul" || no == "nil" ||
		strings.Contains(no, "{") || strings.Contains(no, "}") || strings.Contains(no, "_") ||
		strings.Contains(no, "！") || strings.Contains(no, "@") || strings.Contains(no, "#") || strings.Contains(no, "$") || strings.Contains(no, "%") || strings.Contains(no, "&") || strings.Contains(no, "*") {
		return false
	}
	return true
}

func allZero(s string) bool {
	for _, v := range s {
		if (v >= '1' && v <= '9') || (v >= 'a' && v <= 'e') || (v >= 'A' && v <= 'E') {
			return false
		}
	}
	return true
}

func MD5ValidCi(no string) bool {
	if len(no) != 32 {
		return false
	}

	for _, v := range no {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'f') || (v >= 'A' && v <= 'F') {
			// nothing
		} else {
			return false
		}
	}

	if allZero(no) {
		return false
	}

	return true
}

func MD5Valid(no string) bool {
	if len(no) != 32 {
		return false
	}

	for _, v := range no {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'f') {
			// nothing
		} else {
			return false
		}
	}

	if allZero(no) {
		return false
	}

	return true
}

func IDFAValid(no string) bool {
	if len(no) != 36 {
		return false
	}

	var cnt int
	for _, v := range no {
		if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'F') || v == '-' {
			// nothing
		} else {
			return false
		}
		if v == '-' {
			cnt++
		}
	}
	if cnt != 4 {
		return false
	}

	if allZero(no) {
		return false
	}

	return true
}

func GAIDValid(no string) bool {
	if len(no) != 36 {
		return false
	}

	var cnt int
	for _, v := range no {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'f') || v == '-' {
			// nothing
		} else {
			return false
		}
		if v == '-' {
			cnt++
		}
	}
	if cnt != 4 {
		return false
	}

	if allZero(no) {
		return false
	}

	return true
}

func IMEIValid(no string) bool {
	if len(no) != 14 && len(no) != 15 {
		return false
	}

	for _, v := range no {
		if v >= '0' && v <= '9' {
			// nothing
		} else {
			return false
		}
	}

	if allZero(no) {
		return false
	}

	return true
}

func OAIDValid(no string) bool {
	if len(no) != 16 && len(no) != 32 && len(no) != 64 && len(no) != 36 {
		return false
	}

	for _, v := range no {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'f') || (v >= 'A' && v <= 'F') || v == '-' {
			// nothing
		} else {
			return false
		}
	}

	if allZero(no) {
		return false
	}

	return true
}

func IPValid(ip string) bool {
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}

	for _, v := range ip {
		if (v >= '0' && v <= '9') || v == '.' {
			// nothing
		} else {
			return false
		}
	}

	if allZero(ip) {
		return false
	}

	return true
}
