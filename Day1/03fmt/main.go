package main

import "fmt"

func main() {
	n := 511
	fmt.Printf("%T\n", n) //type	类型
	fmt.Printf("%v\n", n) //value	值
	fmt.Printf("%b\n", n) //bit		二进制
	fmt.Printf("%d\n", n) //decimal	十进制
	fmt.Printf("%o\n", n) //Octet 	八进制
	fmt.Printf("%x\n", n) //Hex 	十六进制
	s := "jubo"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
}
