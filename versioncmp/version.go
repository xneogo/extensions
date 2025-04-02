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
 @Time    : 2025/4/2 -- 15:31
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: versioncmp versioncmp/version.go
*/

package versioncmp

import (
	"fmt"
	"strings"
)

type VersionCmp struct {
	ver string
}

func New(ver string) *VersionCmp {
	v := &VersionCmp{}

	v.ver = v.fmtVer(ver)
	return v
}

func (m *VersionCmp) fmtVer(ver string) string {
	pvs := strings.Split(ver, ".")

	rv := ""
	for _, pv := range pvs {
		rv += fmt.Sprintf("%020s", pv)
	}

	return rv

}

func (m *VersionCmp) Min() string {
	return m.fmtVer("0")
}

func (m *VersionCmp) Max() string {
	return m.fmtVer("99999999999999999999")
}

func (m *VersionCmp) Lt(ver string) bool {
	return m.ver < m.fmtVer(ver)
}

func (m *VersionCmp) Lte(ver string) bool {
	return m.ver <= m.fmtVer(ver)
}

func (m *VersionCmp) Gt(ver string) bool {
	return m.ver > m.fmtVer(ver)
}

func (m *VersionCmp) Gte(ver string) bool {
	return m.ver >= m.fmtVer(ver)
}

func (m *VersionCmp) Eq(ver string) bool {
	return m.ver == m.fmtVer(ver)
}

func (m *VersionCmp) Ne(ver string) bool {
	return m.ver != m.fmtVer(ver)
}

func (m *VersionCmp) GetFormatVersion() string {
	return m.ver
}
