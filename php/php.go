package php

import (
	"net"
	"strings"
)

// UniqueStrSlice 字符串切片去重
func UniqueStrSlice(strSlice []string) []string {
	temp := map[string]string{}

	for _, item := range strSlice {
		temp[item] = item
	}

	unique := []string{}
	for _, itemU := range temp {
		unique = append(unique, itemU)
	}

	return unique
}

// UniqueIntSlice int 切片去重
func UniqueIntSlice(strSlice []int) []int {
	temp := map[int]int{}

	for _, item := range strSlice {
		temp[item] = item
	}

	unique := []int{}
	for _, itemU := range temp {
		unique = append(unique, itemU)
	}

	return unique
}

// UniqueInt64Slice int64 切片去重
func UniqueInt64Slice(strSlice []int64) []int64 {
	temp := map[int64]int64{}

	for _, item := range strSlice {
		temp[item] = item
	}

	unique := []int64{}
	for _, itemU := range temp {
		unique = append(unique, itemU)
	}

	return unique
}

// ExistStrSlice 查找字符是否在数组中
func ExistStrSlice(strSlice []string, str string) bool {

	var exist bool
	for _, item := range strSlice {
		if item == str {
			exist = true
			break
		}
	}

	return exist
}

// ExistIntSlice 数字是否在切片中
func ExistIntSlice(intSlice []int, id int) bool {

	var exist bool
	for _, item := range intSlice {
		if item == id {
			exist = true
			break
		}
	}

	return exist
}

// ExistInt64Slice 数字是否在切片中
func ExistInt64Slice(intSlice []int64, id int64) bool {

	var exist bool
	for _, item := range intSlice {
		if item == id {
			exist = true
			break
		}
	}

	return exist
}

// HtmlSpecialChars Html实体特殊字符转移
func HtmlSpecialChars(str string) string {
	chars := []string{"&", "\"", "'", "<", ">"}
	htmls := []string{"&amp;", "&quot;", "'", "&lt;", "&gt;"}

	for i, ch := range chars {
		str = strings.Replace(str, ch, htmls[i], -1)
	}

	return str
}

func LocalIp() string {
	localIP := "0.0.0.0"
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, address := range addrs {
			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					localIP = ipnet.IP.String()
				}
			}
		}
	}

	return localIP
}
