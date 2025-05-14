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
 @Time    : 2025/5/14 -- 17:11
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xcrypto examples/xcrypto/main.go
*/

package main

import (
	"bytes"
	"context"
	"crypto"
	"encoding/hex"
	"fmt"
	"net/url"

	"github.com/xneogo/extensions/xcrypto"
)

func main() {
	// 示例1：哈希算法
	fmt.Println("示例1：哈希算法")
	text := "Hello, World!"
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("MD5: %s\n", xcrypto.MD5.DoString(text))
	fmt.Printf("SHA1: %s\n", xcrypto.SHA1.DoString(text))
	fmt.Printf("SHA256: %s\n", xcrypto.SHA256.DoString(text))
	fmt.Printf("SHA512: %s\n", xcrypto.SHA512.DoString(text))
	fmt.Printf("CRC32: %s\n", xcrypto.CRC32.DoString(text))

	// 示例2：HMAC
	fmt.Println("\n示例2：HMAC")
	key := []byte("secret")
	data := []byte(text)
	fmt.Printf("HMAC-MD5: %s\n", xcrypto.HMacMD5(key, data))
	fmt.Printf("HMAC-SHA1: %s\n", xcrypto.HMacSHA1(key, data))
	fmt.Printf("HMAC-SHA256: %s\n", xcrypto.HMacSHA256(key, data))

	// 示例3：签名器
	fmt.Println("\n示例3：签名器")
	// 创建签名器
	signer := xcrypto.New(
		xcrypto.WithMethod(xcrypto.NewHashMethod(crypto.SHA256)),
		xcrypto.WithEncoder(&xcrypto.DefaultEncoder{}),
	)

	// 签名URL参数
	values := url.Values{}
	values.Set("name", "test")
	values.Set("age", "18")
	values.Set("timestamp", "1234567890")

	signature, err := signer.SignValues(context.Background(), values)
	if err != nil {
		fmt.Printf("签名错误: %v\n", err)
		return
	}
	fmt.Printf("URL参数签名: %s\n", hex.EncodeToString(signature))

	// 验证签名
	err = signer.VerifyValues(context.Background(), values, signature)
	if err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println("验证成功")
	}

	// 签名字节数组
	data = []byte(text)
	signature, err = signer.SignBytes(context.Background(), data)
	if err != nil {
		fmt.Printf("签名错误: %v\n", err)
		return
	}
	fmt.Printf("字节数组签名: %s\n", hex.EncodeToString(signature))

	// 验证签名
	err = signer.VerifyBytes(context.Background(), data, signature)
	if err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println("验证成功")
	}

	// 示例4：随机数生成
	fmt.Println("\n示例4：随机数生成")
	nonce, err := xcrypto.GenerateNonce()
	if err != nil {
		fmt.Printf("生成随机数错误: %v\n", err)
		return
	}
	fmt.Printf("随机数: %s\n", nonce)

	// 示例5：编码器
	fmt.Println("\n示例5：编码器")
	encoder := &xcrypto.DefaultEncoder{}
	buffer := &bytes.Buffer{}

	// 编码URL参数
	values = url.Values{}
	values.Set("a", "1")
	values.Set("b", "2")
	values.Set("c", "3")

	encoded, err := encoder.EncodeValues(buffer, values, &xcrypto.SignOptions{
		Prefix: "prefix_",
		Suffix: "_suffix",
	})
	if err != nil {
		fmt.Printf("编码错误: %v\n", err)
		return
	}
	fmt.Printf("编码结果: %s\n", string(encoded))

	// 编码字节数组
	data = []byte(text)
	encoded, err = encoder.EncodeBytes(buffer, data, &xcrypto.SignOptions{
		Prefix: "prefix_",
		Suffix: "_suffix",
	})
	if err != nil {
		fmt.Printf("编码错误: %v\n", err)
		return
	}
	fmt.Printf("编码结果: %s\n", string(encoded))
}
