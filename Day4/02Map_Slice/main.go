package main

import "fmt"

func main() {
	s1 := make([]map[int]string, 5, 10)

	s1[0] = make(map[int]string, 1)
	s1[0][5] = "long"
	fmt.Println(s1)

	m1 := make(map[string][]int, 10)
	m1["北京"] = []int{15, 20, 100}
	fmt.Println(m1)
}
