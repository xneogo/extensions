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
 @Time    : 2025/7/3 -- 11:20
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xother xother/device.go
*/

package xother

import (
	"github.com/xneogo/extensions/xcrypto"
	"strings"
)

type DeviceInfo struct {
	formated     bool
	OS           string            `json:"os,omitempty"`
	OSVersion    string            `json:"os_version,omitempty"`
	IDFA         string            `json:"idfa,omitempty"`
	IDFAMD5      string            `json:"idfa_md5,omitempty"`
	CAID         map[string]string `json:"caid,omitempty"`
	CAIDMD5      map[string]string `json:"caid_md5,omitempty"`
	GAID         string            `json:"gaid,omitempty"`
	GAIDMD5      string            `json:"gaid_md5,omitempty"`
	IMEI         string            `json:"imei,omitempty"`
	IMEIMD5      string            `json:"imei_md5,omitempty"`
	OAID         string            `json:"oaid,omitempty"`
	OAIDMD5      string            `json:"oaid_md5,omitempty"`
	AndroidID    string            `json:"android_id,omitempty"`
	AndroidIDMD5 string            `json:"android_id_md5,omitempty"`
	Mac          string            `json:"mac,omitempty"`
	MacMD5       string            `json:"mac_md5,omitempty"`
	Mac1         string            `json:"mac1,omitempty"`
	Mac1MD5      string            `json:"mac1_md5,omitempty"`
	IP           string            `json:"ip,omitempty"`
	IPMD5        string            `json:"ip_md5,omitempty"`
	IPv4         string            `json:"ipv4,omitempty"`
	IPv4MD5      string            `json:"ipv4_md5,omitempty"`
	IPv6         string            `json:"ipv6,omitempty"`
	IPv6MD5      string            `json:"ipv6_md5,omitempty"`
	UA           string            `json:"ua,omitempty"`
	Brand        string            `json:"brand,omitempty"`
	Model        string            `json:"model,omitempty"`
	NetType      string            `json:"net_type,omitempty"`
	Carrier      string            `json:"carrier,omitempty"`
}

const (
	OSAndroid = "android"
	OSiOS     = "ios"
	OSUnknown = "unknown"
)

func (d *DeviceInfo) Format() {
	if d == nil || d.formated {
		return
	}

	if !MacroValid(d.OS) || d.OS == "" {
		d.OS = OSUnknown
	}
	if d.OS != OSAndroid && d.OS != OSiOS {
		d.OS = OSUnknown
	}

	if !MacroValid(d.OSVersion) {
		d.OSVersion = ""
	}

	d.IDFA = strings.ToUpper(d.IDFA)
	d.IDFAMD5 = strings.ToLower(d.IDFAMD5)
	if !MacroValid(d.IDFA) || !IDFAValid(d.IDFA) {
		d.IDFA = ""
	} else {
		d.IDFAMD5 = xcrypto.MD5.DoBytes([]byte(d.IDFA))
	}
	if !MD5Valid(d.IDFAMD5) {
		d.IDFAMD5 = ""
	}

	for k, v := range d.CAID {
		if k == "" || v == "" {
			delete(d.CAID, k)
			continue
		}
		v = strings.ToLower(v)
		if MD5Valid(v) {
			d.CAID[k] = v
			d.CAIDMD5[k] = xcrypto.MD5.DoBytes([]byte(v))
		} else {
			delete(d.CAID, k)
		}
	}
	for k, v := range d.CAIDMD5 {
		if k == "" || v == "" {
			delete(d.CAIDMD5, k)
			continue
		}
		v = strings.ToLower(v)
		if MD5Valid(v) {
			d.CAIDMD5[k] = v
		} else {
			delete(d.CAIDMD5, k)
		}
	}

	d.GAID = strings.ToLower(d.GAID)
	d.GAIDMD5 = strings.ToLower(d.GAIDMD5)
	if !MacroValid(d.GAID) || !GAIDValid(d.GAID) {
		d.GAID = ""
	} else {
		d.GAIDMD5 = xcrypto.MD5.DoBytes([]byte(d.GAID))
	}
	if !MD5Valid(d.GAIDMD5) {
		d.GAIDMD5 = ""
	}

	d.IMEIMD5 = strings.ToLower(d.IMEIMD5)
	if !MacroValid(d.IMEI) || !IMEIValid(d.IMEI) {
		d.IMEI = ""
	} else {
		d.IMEIMD5 = xcrypto.MD5.DoBytes([]byte(d.IMEI))
	}
	if !MD5Valid(d.IMEIMD5) {
		d.IMEIMD5 = ""
	}

	d.OAIDMD5 = strings.ToLower(d.OAIDMD5)
	if !MacroValid(d.OAID) || !OAIDValid(d.OAID) {
		d.OAID = ""
	} else {
		d.OAIDMD5 = xcrypto.MD5.DoBytes([]byte(d.OAID))
	}
	if !MD5Valid(d.OAIDMD5) {
		d.OAIDMD5 = ""
	}

	if d.OS == OSUnknown {
		if d.IDFA != "" || d.IDFAMD5 != "" || len(d.CAID) > 0 || len(d.CAIDMD5) > 0 {
			d.OS = OSiOS
		} else if d.GAID != "" || d.GAIDMD5 != "" || d.IMEI != "" || d.IMEIMD5 != "" || d.OAID != "" || d.OAIDMD5 != "" {
			d.OS = OSAndroid
		}
	}

	d.AndroidIDMD5 = strings.ToLower(d.AndroidIDMD5)
	if !MacroValid(d.AndroidID) {
		d.AndroidID = ""
	} else {
		d.AndroidIDMD5 = xcrypto.MD5.DoBytes([]byte(d.AndroidID))
	}
	if !MD5Valid(d.AndroidIDMD5) {
		d.AndroidIDMD5 = ""
	}

	if !MacroValid(d.Mac) {
		d.Mac = ""
	}
	if !MD5Valid(d.MacMD5) {
		d.MacMD5 = ""
	}
	if !MacroValid(d.Mac1) {
		d.Mac1 = ""
	}
	if !MD5Valid(d.Mac1MD5) {
		d.Mac1MD5 = ""
	}

	if !MacroValid(d.IP) {
		d.IP = ""
	}
	if !MD5Valid(d.IPMD5) {
		d.IPMD5 = ""
	}
	if !MacroValid(d.IPv4) {
		d.IPv4 = ""
	}
	if !MD5Valid(d.IPv4MD5) {
		d.IPv4MD5 = ""
	}
	if !MacroValid(d.IPv6) {
		d.IPv6 = ""
	}
	if !MD5Valid(d.IPv6MD5) {
		d.IPv6MD5 = ""
	}

	if !MacroValid(d.UA) {
		d.UA = ""
	}
	if !MacroValid(d.Brand) {
		d.Brand = ""
	}
	if !MacroValid(d.Model) {
		d.Model = ""
	}
	if !MacroValid(d.NetType) {
		d.NetType = ""
	}
	if !MacroValid(d.Carrier) {
		d.Carrier = ""
	}
}

func (d *DeviceInfo) Reset() {
	if d == nil {
		return
	}
	d.OS = ""
	d.OSVersion = ""
	d.IDFA = ""
	d.IDFAMD5 = ""
	d.CAID = map[string]string{}
	d.CAIDMD5 = map[string]string{}
	d.GAID = ""
	d.GAIDMD5 = ""
	d.IMEI = ""
	d.IMEIMD5 = ""
	d.OAID = ""
	d.OAIDMD5 = ""
	d.AndroidID = ""
	d.AndroidIDMD5 = ""
	d.Mac = ""
	d.MacMD5 = ""
	d.Mac1 = ""
	d.Mac1MD5 = ""
	d.IP = ""
	d.IPMD5 = ""
	d.IPv4 = ""
	d.IPv4MD5 = ""
	d.IPv6 = ""
	d.IPv6MD5 = ""
	d.UA = ""
	d.Brand = ""
	d.Model = ""
	d.NetType = ""
	d.Carrier = ""
}

func (d *DeviceInfo) IsSet() bool {
	if d == nil {
		return false
	}
	d.Format()

	if d.IDFA == "" && d.IDFAMD5 == "" && d.IMEI == "" && d.IMEIMD5 == "" && d.OAID == "" && d.OAIDMD5 == "" {
		return false
	}
	return true
}

func (d *DeviceInfo) Key() string {
	if d == nil {
		return ""
	}
	d.Format()

	if d.OAIDMD5 != "" {
		return d.OAIDMD5
	} else if d.IMEIMD5 != "" {
		return d.IMEIMD5
	} else if d.IDFAMD5 != "" {
		return d.IDFAMD5
	}
	return ""
}
