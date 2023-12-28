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
}
