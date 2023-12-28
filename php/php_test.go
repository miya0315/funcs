package php

import (
	"log"
	"testing"
)

func TestPhpFunc(t *testing.T) {
	// 字符串切片去重
	str := []string{"a", "b", "c", "d", "e", "f", "g", "c", "d", "e", "f"}
	log.Println("UniqueStrSlice:", UniqueStrSlice(str), str)

	// int数字去重
	num := []int{1, 2, 3, 4, 5, 6, 7, 2, 1, 4, 5}
	log.Println("UniqueIntSlice:", UniqueIntSlice(num), num)

	// int64数字去重
	num64 := []int64{1, 2, 3, 4, 5, 6, 7, 2, 1, 4, 5}
	log.Println("UniqueInt64Slice:", UniqueInt64Slice(num64), num64)

	// int64数字去重
	strE := []string{"a", "b", "c", "d", "e", "f", "g", "c", "d", "e", "f"}
	log.Println("ExistStrSlice:", ExistStrSlice(strE, "b"), strE)

	// int64数字去重
	numE := []int{1, 2, 3, 4, 5, 6, 7, 2, 1, 4, 5}
	log.Println("ExistIntSlice:", ExistIntSlice(numE, 5), numE)

	// int64数字去重
	num64E := []int64{1, 2, 3, 4, 5, 6, 7, 2, 1, 4, 5}
	log.Println("ExistInt64Slice:", ExistInt64Slice(num64E, 9), num64E)

	// int64数字去重
	strStr := "<div class=\"div\">测试</div>"
	log.Println("HtmlSpecialChars:", HtmlSpecialChars(strStr), strStr)
}
