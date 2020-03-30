package main

import "fmt"

func main() {
	s := "jubo"
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
	}
	for _, a := range s {
		fmt.Println(a)
	}
	i := 0
	for i < 10 {

		fmt.Println(i)
		i++
	}
	// for {
	// 	//死循环
	// }

}
