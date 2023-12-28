package create

import (
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	// 生成唯一的时间订单Id
	for i := 0; i < 5; i++ {
		log.Println("CreateUniqueOrderNo:", CreateUniqueId("TEST"))
	}

	// 生成指定随机字符串：
	for i := 0; i < 5; i++ {
		log.Println("CreateRandomString:", CreateRandomString(32))
	}

	// 生成指定随机数字：
	for i := 0; i < 5; i++ {
		log.Println("CreateRandomNumStr:", CreateRandomNumStr(8))
	}

	money := "123.321"
	log.Println("CreateMoneyToChina:", money, CreateMoneyToChina(money))

	money = "0.321"
	log.Println("CreateMoneyToChina:", money, CreateMoneyToChina(money))

	money = "10000.00"
	log.Println("CreateMoneyToChina:", money, CreateMoneyToChina(money))

	money = "10000"
	log.Println("CreateMoneyToChina:", money, CreateMoneyToChina(money))

	str := "生成指定随机数字2131321321地方："
	log.Println("CreateNewStrWithStar:", str, CreateNewStrWithStar(str, ""))

}
