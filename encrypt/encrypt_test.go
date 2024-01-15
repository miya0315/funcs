package encrypt

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {

	// 生成Rsa对称秘钥
	filePath := "../create/"
	filePath = ""
	private, public, err := GenRsaKey(2048, filePath)
	log.Println("GenRsaKey: ", private, public, err)

	// 公钥加密
	str := "RsaPubKeyEncrypt"
	encryptStr, err := RsaPubKeyEncrypt(str, public, true)
	log.Println("RsaPubKeyEncrypt: ", str, encryptStr, err)

	// 私钥解密
	desStr, err := RsaPrivateKeyDecrypt(encryptStr, private, true)
	log.Println("RsaPrivateKeyDecrypt: ", desStr, err)

	// ECB PKCS5Padding 加密
	strByte := []byte("ECBEncrypt")
	key := []byte("abcdefghijklmnjh")
	enStr, err := ECBEncrypt(strByte, key)
	log.Println("ECBEncrypt: ", string(strByte), string(enStr), err)

	// ECB PKCS5Padding 解密
	dsStr, err := ECBDecrypt(enStr, key)
	log.Println("ECBDecrypt: ", string(key), string(dsStr), err)

	// Sha256Encrypt
	str = "Sha256Encrypt"
	log.Println("Sha256Encrypt: ", Sha256Encrypt(str))

	// Sha256SalfEncrypt
	log.Println("Sha256SaltEncrypt: ", Sha256SaltEncrypt(str, "Sha256SaltEncrypt"))

	// Base64UrlEncode
	log.Println("Base64UrlEncode: ", Base64UrlEncode(str))
}
