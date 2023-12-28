package create

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/miya0315/funcs/changliang"
)

// CreateUniqueOrderNo 创建唯一Id:长度 len(prefix)+27
func CreateUniqueId(prefix string) string {
	strTime := time.Now().Format(changliang.DATETIMEINT + ".00000000")
	strTime = strings.Replace(strTime, ".", "", -1)
	randStr := ""
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(89999)
	randStr = fmt.Sprintf("%d", (randNum + 10000))
	return prefix + strTime + randStr
}

// CreateRandomString 随机长度字符串
func CreateRandomString(length int) string {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var charByte []byte
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(len(strByte))
		charByte = append(charByte, strByte[randNum])
	}
	return string(charByte)
}

// Random 生成随机数
func CreateRandomNumStr(length int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	var strByte []byte
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(len(table))
		strByte = append(strByte, table[randNum])
	}
	return string(strByte)
}

