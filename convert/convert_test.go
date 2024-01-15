package convert

import (
	"log"
	"testing"
)

func TestConvert(t *testing.T) {
	by := 1.024
	log.Println("ByteSizeToStr: ", by, ByteSizeToStr(by))

	by *= 1000
	log.Println("ByteSizeToStr: ", by, ByteSizeToStr(by))

	by *= 1000
	log.Println("ByteSizeToStr: ", by, ByteSizeToStr(by))

	by *= 1000
	log.Println("ByteSizeToStr: ", by, ByteSizeToStr(by))

	by *= 1000
	log.Println("ByteSizeToStr: ", by, ByteSizeToStr(by))
}
