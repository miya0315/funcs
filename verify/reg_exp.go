package verify

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/miya0315/funcs/changliang"
)

// RegExpCardPwd 积分卡卡密规则 ****-****-****-**** 0-9/a-f
func RegExpCardPwd(str string) bool {
	resStr := "^[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}$"
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpNumChar 校验指定长度的数组、字母组合
func RegExpChar(str string, min, max int) bool {
	resStr := fmt.Sprintf("^[a-zA-Z0-9]{%d,%d}$", min, max)
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpNum 校验指定长度的数字
func RegExpNum(str string, min, max int) bool {
	resStr := fmt.Sprintf("^[0-9]{%d,%d}$", min, max)
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpIsUid 用户名：字母、数字、点、中横线、下划线 
func RegExpIsUid(str string, min, max int) bool {
	resStr := fmt.Sprintf("^[a-zA-Z0-9._\\-]{%d,%d}$", min, max)
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpIsMobile 手机号校验
func RegExpIsMobile(str string) bool {
	resStr := "^1[0-9]{10}$"
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpIsMobiles 手机号校验
func RegExpIsMobiles(str string) bool {
	resStr := "^1[3-9]{1}[0-9]{9}(#+1[3-9]{1}[0-9]{9}#{0,}){0,}$"
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpIsEmail 邮箱校验
func RegExpIsEmail(str string) bool {
	// 判断是否是临时邮箱
	if isTempEmail(str) {
		return false
	}
	resStr := "^[0-9a-zA-Z]+([-._]?[0-9a-zA-Z]+)*@[0-9a-zA-Z]+([-._]?[0-9a-zA-Z]+)*(.[a-zA-Z]{2,3})+$"
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// isTempEmail 临时邮箱后缀校验
func isTempEmail(str string) bool {
	tempEmail := []string{
		"@trbvm.com",
		"@chacuo.net",
		"@027168.com",
		"@mail.bccto.me",
		"@bccto.me",
		"@yopmail.com",
		"@guerrillamail",
		"@nowmymail",
		"@fakeinbox.com",
		"@mailnesia.com",
		"@mailinator",
		"@maildrop.cc",
		"@trashymail.com",
		"@mt2015.com",
		"@maildu.de",
		"@sharklasers.com",
		"@www.bccto.me",
		"bccto.me",
	}
	var isTemp bool
	for _, val := range tempEmail {
		if strings.Contains(str, val) {
			isTemp = true
		}
	}
	return isTemp
}

// RegExpChinaName 中文姓名:字母、数字、点、#、中文点
func RegExpChinaName(str string, min, max int) bool {
	resStr := fmt.Sprintf("^[0-9·\\.A-Za-z\u4e00-\u9fa5#]{%d,%d}$", min, max)

	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpOnlyChina 仅限中文名：中文、少数民族点
func RegExpOnlyChina(str string, min, max int) bool {

	resStr := fmt.Sprintf("^[\u4e00-\u9fa5·\\.]{%d,%d}$", min, max)

	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegChinaString 中文名称，地址啥的
func RegExpChinaString(str string, min, max int) bool {

	resStr := fmt.Sprintf("^[0-9A-Za-z&\u4e00-\u9fa5\\-\\/_\\(\\)\\[\\]（）\\.:,!。、，：【】！？\\s\\?#]{%d,%d}$", min, max)

	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegChinaString 中文名称，地址啥的
func RegExpNoChar(str, noChar string, min, max int) bool {

	resStr := fmt.Sprintf("^[^%s]{%d,%d}$", noChar, min, max)

	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpIdCard 身份证格式校验
func RegExpIdCard(idCard string) bool {
	idCard = strings.ToUpper(idCard)

	// 身份证基本格式校验
	resStr := `(^[0-9]{15}$)|(^[0-9]{17}([0-9]|X)$)`
	reg := regexp.MustCompile(resStr)
	if !reg.MatchString(idCard) {
		return false
	}

	// 15位老格式的身份证号码
	if len(idCard) == 15 {
		// 身份证基本格式校验
		resStr := `^(\d{6})(\d{2})(\d{2})(\d{2})(\d{3})$`
		reg := regexp.MustCompile(resStr)
		matches := reg.FindStringSubmatch(idCard)
		if len(matches) == 6 {
			birth := fmt.Sprintf("19%s-%s-%s", matches[2], matches[3], matches[4])
			if _, err := time.Parse(changliang.DATE, birth); err == nil {
				return true
			}
		}
	} else {
		// 18位新格式的身份证号码
		regStr18 := `^(\d{6})(\d{4})(\d{2})(\d{2})(\d{3})([0-9]|X)$`
		reg18 := regexp.MustCompile(regStr18)
		matches := reg18.FindStringSubmatch(idCard)
		birth := fmt.Sprintf("%s-%s-%s", matches[2], matches[3], matches[4])
		if _, err := time.Parse(changliang.DATE, birth); err == nil {
			// 检验18位身份证的校验码是否正确。
			// 校验位按照ISO 7064:1983.MOD 11-2的规定生成，X可以认为是数字10。
			intSlice := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
			strSlice := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
			var sign int
			for i := 0; i < 17; i++ {
				a, _ := strconv.Atoi(string(idCard[i]))
				sign += a * intSlice[i]
			}

			n := sign % 11
			if strSlice[n] == string(idCard[17]) {
				return true
			}
		}
	}

	return false
}

// RegExpMoney 金额校验
func RegExpMoney(str string, demic int) bool {

	resStr := fmt.Sprintf("^(-)?[0-9]{1,10}(.[0-9]{1,%d})?$", demic)
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpDate 日期格式校验
func RegExpDate(str string) bool {
	resStr := `^(20|19)\d{2}-((0[1-9])|(1[0-2]))-(([012][0-9])|(3[01]))$`
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpDatetime 日期时间格式校验
func RegExpDatetime(str string) bool {
	resStr := `^(20|19)\d{2}-((0[1-9])|(1[0-2]))-(([012][0-9])|(3[01]))\s((([01]\d)|(2[0-3])):[0-5]\d:[0-5]\d)$`
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpUrl url地址校验
func RegExpUrl(str string) bool {
	resStr := `^(http[s]?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?\'\\+&%$#=~_-]+))*$`
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpIp ip地址校验
func RegExpIp(str string) bool {
	resStr := `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}

// RegExpHardPhone 固定电话号码校验(0512-12321321 | (0512)12321323)
func RegExpHardPhone(str string) bool {
	resStr := `^(\(\d{3,4}\)|\d{3,4}-|\s)?\d{7,14}$`
	reg := regexp.MustCompile(resStr)
	return reg.MatchString(str)
}
