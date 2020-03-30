package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	s1 = []int{1, 2, 3}
	s2 = []string{"鞠博", "zhoulei", "huxuan"}
	fmt.Println(s1, s2)
	s3 := []int{1, 2, 4, 5, 6, 8}
	fmt.Println(s3)
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
	s4 := s3[1:3]
	fmt.Println(s4)
	s5 := s3[1:3]
	s6 := s3[:3]
	s7 := s3[:]
	fmt.Println(s5, s6, s7)
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7), cap(s7))
}
