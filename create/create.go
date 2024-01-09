package create

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/miya0315/funcs/changliang"
	"github.com/xuri/excelize/v2"
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

// CreateMoneyToChina 数字金额转化为中文
func CreateMoneyToChina(str string) string {
	var (
		digits     []string
		radices    []string
		bigRadices []string
		decimals   []string
		cnDollar   string
		cnInteger  string
		outStr     string
		intStr     string
		floatStr   string
	)
	digits = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	radices = []string{"", "拾", "佰", "仟", "万", "亿"}
	bigRadices = []string{"", "万", "亿", "万"}
	decimals = []string{"角", "分", "厘"}
	cnDollar = "元"
	cnInteger = "整"

	money := strings.Split(str, ".")
	intStr = money[0]
	if len(money) > 1 {
		floatStr = money[1]
	}
	log.Println("money[]:", money)
	if intStr != "0" {
		intNumLen := len(intStr)
		zeroCount := 0

		for i := 0; i < intNumLen; i++ {
			p := intNumLen - i - 1
			d, _ := strconv.Atoi(string(intStr[i]))

			quotient := p / 4
			modulus := p % 4

			if d == 0 {
				zeroCount++
			} else {
				if zeroCount > 0 {
					outStr += bigRadices[0]
				}
				zeroCount = 0
				outStr += digits[d] + radices[modulus]
			}

			if modulus == 0 && zeroCount < 4 {
				outStr += bigRadices[quotient]
				zeroCount = 0
			}

		}
		outStr += cnDollar
	}

	if floatStr != "" && floatStr != "0" {
		floadLen := len(floatStr)
		for i := 0; i < floadLen; i++ {
			d, _ := strconv.Atoi(string(floatStr[i]))
			if d != 0 {
				outStr += digits[d] + decimals[i]
			}
		}
	}

	if outStr == "" {
		outStr = digits[0] + cnDollar
	}

	if floatStr != "" {
		d, _ := strconv.Atoi(string(floatStr[0]))
		if d != 0 {
			outStr += cnInteger
		}
	}

	return outStr
}

// CreateNewStrWithStar 替换字符串中间1/3字符串为*
func CreateNewStrWithStar(str, symbol string) string {
	if symbol == "" {
		symbol = "*"
	}

	strLen := utf8.RuneCountInString(str)
	starLen := int(strLen / 3)
	strSlice := []rune(str)
	star := strings.Repeat(symbol, starLen+1)

	return string(strSlice[:starLen]) + star + string(strSlice[2*starLen:])
}

// CreateExcelFile 导出excel
func CreateExcelFile(fileName string, data [][]interface{}) error {
	f := excelize.NewFile()

	// 创建一个新的工作表
	index, _ := f.NewSheet("Sheet1")

	// 为工作表设置单元格的值
	for i, row := range data {
		for j, value := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue("Sheet1", cell, value)
		}
	}

	// 设置工作表为活动状态
	f.SetActiveSheet(index)

	// 保存文件
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}
