package main

import "fmt"

//对于ascii表中的字符可以使用byte
//而其他的使用 rune

func main() {
	s := "wo shi zhong guo 人"
	n := len(s)
	fmt.Println(n)

	for i := 0; i < n; i++ {
		fmt.Printf("%c\n", s[i])
		fmt.Printf("%T\n", s[i])
	}
	// for _, sr := range s {
	// 	fmt.Printf("%c\t", sr)

	// }

	//类型转换 T（value）
	a := 78
	h := byte(a)
	fmt.Printf("%c\n", h)
}
