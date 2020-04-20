package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//获取汉字的数量
	s1 := "Hello 鞠博"
	count := 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println("一共汉字： ", count, " 个")

	//单词出现的次数(切分字符串)
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10) //某个单词的出现个数
	for _, w := range s3 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		// fmt.Println(w)
		// m1[w]++
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
