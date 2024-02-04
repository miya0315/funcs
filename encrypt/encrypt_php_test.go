package encrypt

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/miya0315/funcs/php"
)

var secretKey string = "103c3eea0bddbf186ce8ea125279935e"

func TestEncryptPhp(t *testing.T) {
	log.Println()
	log.Println()
	log.Println()

	var token string = "2Hn0al0oaXlZydXylTAj75IT6SMcK7B42SjrvBwN"
	strSid2 := LaravelEncode(token)
	log.Println("加密数据：" + strSid2)

	destr := LaravelDecrypt(strSid2)
	log.Println("==", destr)
}


// LaravelDecrypt laravel cookie 解密
func LaravelDecrypt(strSid string) interface{} {
	str64, _ := Base64Decode(strSid)
	// 解析sid
	var payload map[string]string
	json.Unmarshal([]byte(str64), &payload)

	if validMac(payload) {
		log.Println("Mac 校验成功")
		strSour, _ := Base64Decode(payload["value"])
		iv, _ := Base64Decode(payload["iv"])

		secStr, err := CBCDecrypt([]byte(strSour), []byte(secretKey), []byte(iv))

		log.Println("解密数据：", string(secStr), err)
		str, _ := php.Unserialize(string(secStr))

		return str
	}
	return ""
}

// validMac 校验mac是否合法
func validMac(payload map[string]string) bool {
	bytes, _ := RandomBytes(16)
	bytesStr := Base64Encode(string(bytes))

	// 通过key 加密计算mac
	hasMac := ToXString(Hash256Mac(payload["iv"]+payload["value"], secretKey))
	hasMac2 := ToXString(Hash256Mac(string(hasMac), bytesStr))
	// paylodMac
	payloadMac := ToXString(Hash256Mac(payload["mac"], bytesStr))

	return HashEquals([]byte(payloadMac), []byte(hasMac2))
}

// cbcSidEncrypt 加密
func LaravelEncode(str string) string {
	bytes, _ := RandomBytes(16)

	str, _ = php.Serialize(str)
	enVal, _ := CBCEncrypt([]byte(str), []byte(secretKey), bytes)
	value := Base64Encode(string(enVal)) // php openssl_encrypt
	iv := Base64Encode(string(bytes))

	macByte := ToXString(Hash256Mac(iv+value, secretKey))

	result := map[string]string{
		"iv":    iv,
		"value": value,
		"mac":   string(macByte),
	}
	jsonStr, _ := json.Marshal(result)
	return Base64Encode(string(jsonStr))
}
