package php

import (
	"log"
	"testing"
)

func TestPhpSerializeFunc(t *testing.T) {
	str := `s:40:"T6SMcK7B42Sjr2Hn0al0oaXlZydXylTAj75IvBwN"`
	strUn, _ := Unserialize(str)
	log.Println(strUn)

	strMap := map[string]interface{}{
		"token":"T6SMcK7B42Sjr2Hn0al0oaXlZydXylTAj75IvBwN",
		"id":"string12",
	}
	strS, _ := Serialize(strMap)
	log.Println("---",strS)

	strM,_:=Unserialize(strS)
	log.Println("---",strM)

}
