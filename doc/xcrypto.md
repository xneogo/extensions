# 加密工具 (xcrypto)

## 简介

完整的加密工具库，包含：
- 哈希算法：MD5、SHA1、SHA256、SHA512、CRC32
- HMAC签名：支持多种哈希算法的HMAC
- 对称加密：AES加密，支持ECB、CBC等模式
- 编码解码：Base64、Base62、Base58等编码
- 填充方案：PKCS5、PKCS7、Zero padding等
- 证书处理：PEM格式证书加载和验证

## 安装

```bash
go get github.com/xneogo/extensions/xcrypto
```

## 使用示例

### 哈希算法

```go
import "github.com/xneogo/extensions/xcrypto"

// 哈希计算
text := "Hello, World!"
md5 := xcrypto.MD5.DoString(text)
sha256 := xcrypto.SHA256.DoString(text)
crc32 := xcrypto.CRC32.DoString(text)

// 字节数组哈希
data := []byte(text)
md5Bytes := xcrypto.MD5.Do(data)
sha256Bytes := xcrypto.SHA256.Do(data)

fmt.Printf("MD5: %s\n", md5)
fmt.Printf("SHA256: %s\n", sha256)
fmt.Printf("CRC32: %s\n", crc32)
```

### HMAC签名

```go
// HMAC签名
key := []byte("secret")
data := []byte(text)
hmacSHA256 := xcrypto.HMacSHA256(key, data)
hmacMD5 := xcrypto.HMacMD5(key, data)

fmt.Printf("HMAC-SHA256: %x\n", hmacSHA256)
fmt.Printf("HMAC-MD5: %x\n", hmacMD5)

// 验证HMAC
isValid := xcrypto.VerifyHMacSHA256(key, data, hmacSHA256)
fmt.Printf("HMAC验证: %v\n", isValid)
```

### AES加密

```go
import "github.com/xneogo/extensions/xcrypto/aescipher"

ctx := context.Background()
plaintext := []byte("sensitive data")
secretKey := []byte("1234567890123456") // 16字节密钥
iv := []byte("1234567890123456")        // 16字节IV

// ECB模式加密
encrypted, err := aescipher.ECB.Encrypt(ctx, plaintext, secretKey, iv, nil)
if err != nil {
    fmt.Printf("加密失败: %v\n", err)
    return
}

decrypted, err := aescipher.ECB.Decrypt(ctx, encrypted, secretKey, iv, nil)
if err != nil {
    fmt.Printf("解密失败: %v\n", err)
    return
}

fmt.Printf("原文: %s\n", string(plaintext))
fmt.Printf("密文: %x\n", encrypted)
fmt.Printf("解密: %s\n", string(decrypted))
```

### CBC模式加密

```go
// CBC模式加密
encrypted, err := aescipher.CBC.Encrypt(ctx, plaintext, secretKey, iv, nil)
if err != nil {
    fmt.Printf("CBC加密失败: %v\n", err)
    return
}

decrypted, err := aescipher.CBC.Decrypt(ctx, encrypted, secretKey, iv, nil)
if err != nil {
    fmt.Printf("CBC解密失败: %v\n", err)
    return
}

fmt.Printf("CBC加密结果: %x\n", encrypted)
fmt.Printf("CBC解密结果: %s\n", string(decrypted))
```

### GCM模式加密

```go
// GCM模式加密（带认证）
additionalData := []byte("additional auth data")
encrypted, err := aescipher.GCM.Encrypt(ctx, plaintext, secretKey, iv, additionalData)
if err != nil {
    fmt.Printf("GCM加密失败: %v\n", err)
    return
}

decrypted, err := aescipher.GCM.Decrypt(ctx, encrypted, secretKey, iv, additionalData)
if err != nil {
    fmt.Printf("GCM解密失败: %v\n", err)
    return
}

fmt.Printf("GCM加密结果: %x\n", encrypted)
fmt.Printf("GCM解密结果: %s\n", string(decrypted))
```

### Base编码

```go
import "github.com/xneogo/extensions/xcrypto/base"

text := "hello world"

// Base64编码
encoded64 := base.Base64.SEncode(text)
decoded64, err := base.Base64.SDecode(encoded64)
fmt.Printf("Base64编码: %s\n", encoded64)
fmt.Printf("Base64解码: %s\n", decoded64)

// Base62编码
encoded62 := base.Base62.SEncode(text)
decoded62, err := base.Base62.SDecode(encoded62)
fmt.Printf("Base62编码: %s\n", encoded62)
fmt.Printf("Base62解码: %s\n", decoded62)

// 字节数组编码
data := []byte(text)
encoded64Bytes := base.Base64.Encode(data)
decoded64Bytes, err := base.Base64.Decode(encoded64Bytes)
```

### RSA加密

```go
import "github.com/xneogo/extensions/xcrypto"

// 生成RSA密钥对
privateKey, publicKey, err := xcrypto.GenerateRSAKeyPair(2048)
if err != nil {
    fmt.Printf("生成密钥失败: %v\n", err)
    return
}

// RSA加密
plaintext := []byte("RSA encryption test")
encrypted, err := xcrypto.RSAEncrypt(publicKey, plaintext)
if err != nil {
    fmt.Printf("RSA加密失败: %v\n", err)
    return
}

// RSA解密
decrypted, err := xcrypto.RSADecrypt(privateKey, encrypted)
if err != nil {
    fmt.Printf("RSA解密失败: %v\n", err)
    return
}

fmt.Printf("RSA原文: %s\n", string(plaintext))
fmt.Printf("RSA密文: %x\n", encrypted)
fmt.Printf("RSA解密: %s\n", string(decrypted))
```

### 数字签名

```go
// RSA签名
message := []byte("message to sign")
signature, err := xcrypto.RSASign(privateKey, message)
if err != nil {
    fmt.Printf("签名失败: %v\n", err)
    return
}

// 验证签名
isValid := xcrypto.RSAVerify(publicKey, message, signature)
fmt.Printf("签名验证: %v\n", isValid)
```

### 随机数生成

```go
// 生成随机字节
randomBytes, err := xcrypto.GenerateRandomBytes(32)
if err != nil {
    fmt.Printf("生成随机字节失败: %v\n", err)
    return
}

// 生成随机字符串
randomString := xcrypto.GenerateRandomString(16)
fmt.Printf("随机字节: %x\n", randomBytes)
fmt.Printf("随机字符串: %s\n", randomString)

// 生成Nonce
nonce := xcrypto.GenerateNonce(12)
fmt.Printf("Nonce: %x\n", nonce)
```

## API 参考

### 哈希算法

```go
// 哈希器接口
type Hasher interface {
    Do(data []byte) []byte
    DoString(data string) string
}

// 预定义哈希器
var (
    MD5    Hasher
    SHA1   Hasher
    SHA256 Hasher
    SHA512 Hasher
    CRC32  Hasher
)
```

### HMAC函数

```go
func HMacSHA256(key, data []byte) []byte
func HMacMD5(key, data []byte) []byte
func VerifyHMacSHA256(key, data, signature []byte) bool
func VerifyHMacMD5(key, data, signature []byte) bool
```

### AES加密

```go
// 加密器接口
type Cipher interface {
    Encrypt(ctx context.Context, plaintext, key, iv, additional []byte) ([]byte, error)
    Decrypt(ctx context.Context, ciphertext, key, iv, additional []byte) ([]byte, error)
}

// AES加密模式
var (
    ECB Cipher // ECB模式
    CBC Cipher // CBC模式
    CFB Cipher // CFB模式
    GCM Cipher // GCM模式
)
```

### RSA加密

```go
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)
func RSAEncrypt(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error)
func RSADecrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error)
func RSASign(privateKey *rsa.PrivateKey, message []byte) ([]byte, error)
func RSAVerify(publicKey *rsa.PublicKey, message, signature []byte) bool
```

### Base编码

```go
// 编码器接口
type Encoder interface {
    Encode(data []byte) []byte
    Decode(data []byte) ([]byte, error)
    SEncode(data string) string
    SDecode(data string) (string, error)
}

// 预定义编码器
var (
    Base64 Encoder
    Base62 Encoder
    Base58 Encoder
)
```

## 最佳实践

1. **密钥管理**：
   - 使用足够长的密钥（AES-256，RSA-2048以上）
   - 安全存储密钥，避免硬编码
   - 定期轮换密钥

2. **初始化向量(IV)**：
   - 每次加密使用不同的随机IV
   - IV不需要保密，但必须随机
   - GCM模式中IV不能重复使用

3. **填充攻击防护**：
   - 优先使用认证加密模式（GCM）
   - 验证消息完整性
   - 防范时序攻击

4. **哈希使用**：
   - 存储密码时使用bcrypt或scrypt
   - 数据完整性验证使用SHA-256或更强算法
   - 避免使用MD5和SHA-1（除非兼容性要求）

5. **随机数生成**：
   - 使用加密安全的随机数生成器
   - 足够的熵源保证随机性

## 安全注意事项

- **密钥安全**：永远不要在代码中硬编码密钥
- **算法选择**：避免使用已知有漏洞的算法
- **实现安全**：防范侧信道攻击和时序攻击
- **数据清理**：使用完毕后清理内存中的敏感数据
- **错误处理**：不要泄露加密过程中的错误信息 