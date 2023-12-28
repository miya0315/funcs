package verify

import (
	"log"
	"testing"
)

func TestRegExp(t *testing.T) {
	// 16为卡号。卡密校验：
	str := "AC34-d5ff-DcbF-deD0"
	log.Println("RegExpCardPwd: ", str, RegExpCardPwd(str))

	// 校验字符串：数字、字母
	str = "fasfdse234ewr"
	log.Println("RegExpChar: ", str, RegExpChar(str, 5, 32))
	str = "fasfdse-234ewr"
	log.Println("RegExpChar: ", str, RegExpChar(str, 5, 32))

	// 校验数字：数字
	str = "12332"
	log.Println("RegExpNum: ", str, RegExpNum(str, 5, 32))
	str = "f23wr"
	log.Println("RegExpNum: ", str, RegExpNum(str, 5, 32))

	// RegExpIsUid 用户名：字母、数字、点、中横线、下划线
	str = "jfdsm123_3"
	log.Println("RegExpIsUid: ", str, RegExpIsUid(str, 5, 32))
	str = "f23wrfd-fds.!"
	log.Println("RegExpIsUid: ", str, RegExpIsUid(str, 5, 32))

	// RegExpIsMobile 手机号校验
	str = "18913444444"
	log.Println("RegExpIsMobile: ", str, RegExpIsMobile(str))
	str = "28123213213"
	log.Println("RegExpIsMobile: ", str, RegExpIsMobile(str))

	// RegExpIsMobiles 手机号校验
	str = "18913444444#13513523456"
	log.Println("RegExpIsMobiles: ", str, RegExpIsMobiles(str))
	str = "18123213213"
	log.Println("RegExpIsMobiles: ", str, RegExpIsMobiles(str))

	// RegExpIsEmail 邮箱格式校验；校验临时邮箱
	str = "189@189.cb"
	log.Println("RegExpIsEmail: ", str, RegExpIsEmail(str))
	str = "189@trbvm.com"
	log.Println("RegExpIsEmail: ", str, RegExpIsEmail(str))

	// RegExpNoChar 校验字符串不包含字符
	str = "189@189[.cb"
	log.Println("RegExpNoChar: ", str, "<>", RegExpNoChar(str, "<>", 5, 31))
	str = "189@trbvm.com"
	log.Println("RegExpNoChar: ", str, "@", RegExpNoChar(str, "@", 5, 31))

	// RegExpIdCard 身份证校验
	str = "41282819910320527X"
	log.Println("RegExpIdCard-18: ", str, RegExpIdCard(str))

	str = "412828199103205279"
	log.Println("RegExpIdCard-18: ", str, RegExpIdCard(str))

	str = "412828910320527"
	log.Println("RegExpIdCard-15: ", str, RegExpIdCard(str))

	// RegExpDate 日期校验
	str = "2023-10-10"
	log.Println("RegExpDate: ", str, RegExpDate(str))

	str = "2023-10-10 11:11:11"
	log.Println("RegExpDatetime: ", str, RegExpDatetime(str))

	// RegExpUrl url地址校验
	str = "https://github.com/miya0315/funcs"
	log.Println("RegExpUrl: ", str, RegExpUrl(str))

	// RegExpIp ip地址校验
	str = "127.0.0.1"
	log.Println("RegExpIp: ", str, RegExpIp(str))

	// RegExpHardPhone 固定号码校验：(0512-12321321 | (0512)12321323)
	str = "0512-12321321"
	log.Println("RegExpHardPhone: ", str, RegExpHardPhone(str))

	str = "(0512)12321323"
	log.Println("RegExpHardPhone: ", str, RegExpHardPhone(str))
}
