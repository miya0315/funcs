package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// ECBDecrypt ECB模式解密
func ECBDecrypt(crypted, key []byte) ([]byte, error) {
	if !validKey(key) {
		return nil, fmt.Errorf("秘钥长度错误,当前传入长度为 %d", len(key))
	}
	if len(crypted) < 1 {
		return nil, fmt.Errorf("源数据长度不能为0")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypted)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("源数据长度必须是 %d 的整数倍，当前长度为：%d", block.BlockSize(), len(crypted))
	}
	var dst []byte
	tmpData := make([]byte, block.BlockSize())

	for index := 0; index < len(crypted); index += block.BlockSize() {
		block.Decrypt(tmpData, crypted[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}
	dst, err = PKCS5UnPadding(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// ECBEncrypt ECB模式加密
func ECBEncrypt(src, key []byte) ([]byte, error) {
	if !validKey(key) {
		return nil, fmt.Errorf("秘钥长度错误, 当前传入长度为 %d", len(key))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src) < 1 {
		return nil, fmt.Errorf("源数据长度不能为0")
	}
	src = PKCS5Padding(src, block.BlockSize())
	if len(src)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("源数据长度必须是 %d 的整数倍，当前长度为：%d", block.BlockSize(), len(src))
	}
	var dst []byte
	tmpData := make([]byte, block.BlockSize())
	for index := 0; index < len(src); index += block.BlockSize() {
		block.Encrypt(tmpData, src[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}
	return dst, nil
}

// PKCS5Padding PKCS5填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding去除PKCS5填充
func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])

	if length < unpadding {
		return nil, fmt.Errorf("invalid unpadding length")
	}
	return origData[:(length - unpadding)], nil
}

// validKey 秘钥长度验证
func validKey(key []byte) bool {
	k := len(key)
	switch k {
	default:
		return false
	case 16, 24, 32:
		return true
	}
}

// Md5 计算md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Encode base64加密
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Base64Decode base64解密
func Base64Decode(input string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

// Sha256Encrypt sha256 不加盐
func Sha256Encrypt(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256SaltEncrypt sha256 加盐加密
func Sha256SaltEncrypt(str,salt string) string {
	if salt == "" {
		salt = "sha-256-salt-encrypt:2023-12-28" //自定义加盐字符串
	}
	hash := sha256.New()
	hash.Write([]byte(str + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// GenRsaKey RSA公钥私钥产生 filePaht :./ 需要以 / 结尾
func GenRsaKey(bits int, filePath string) (private string, public string, err error) {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", errors.New("生成私钥异常：" + err.Error())
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	private = string(pem.EncodeToMemory(block))

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", errors.New("生成公钥钥异常：" + err.Error())
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	public = string(pem.EncodeToMemory(block))

	if filePath != "" {
		_, err = os.Stat(filePath)
		if err == nil {
			// 写入文件可以用下面
			file, err := os.Create(filePath + "private.pem")
			if err != nil {
				return "", "", errors.New("秘钥文件初始化失败：" + err.Error())
			}
			_, err = file.WriteString(private)
			if err != nil {
				return "", "", errors.New("秘钥写入文件失败：" + err.Error())
			}

			file, err = os.Create(filePath + "public.pem")
			if err != nil {
				return "", "", errors.New("公钥文件初始化失败：" + err.Error())
			}
			_, err = file.WriteString(public)

			if err != nil {
				return "", "", errors.New("公钥写入失败：：" + err.Error())
			}
		} else {
			return "", "", errors.New("生成秘钥文件目录不存在：" + err.Error())
		}
	}

	return
}

// RsaPubKeyEncrypt 公钥加密
func RsaPubKeyEncrypt(str, pubKey string, base64 bool) (encryptStr string, err error) {

	block, _ := pem.Decode([]byte(pubKey))

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err = errors.New("公钥秘钥异常")
		return
	}
	pk := publicKey.(*rsa.PublicKey)
	en, err := rsa.EncryptPKCS1v15(rand.Reader, pk, []byte(str))
	if err != nil {
		err = errors.New("公钥加密异常")
		return
	}
	encryptStr = string(en)
	if base64 {
		encryptStr = Base64Encode(encryptStr)
	}
	return
}

// RsaPrivateKeyDecrypt 公钥解密
func RsaPrivateKeyDecrypt(str, privateKey string, base64 bool) (decryptStr string, err error) {

	if base64 {
		str, err = Base64Decode(str)
		if err != nil {
			return "", err
		}
	}

	block, _ := pem.Decode([]byte(privateKey))

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err = errors.New("秘钥解密异常")
		return
	}
	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, []byte(str))

	if err != nil {
		err = errors.New("签名解密失败")
		return
	}

	decryptStr = string(decryptBytes)
	return
}
