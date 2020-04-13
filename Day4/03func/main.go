package main

import "fmt"

func main() {
	r := sum(1, 3)
	fmt.Println(r)
	r = subtr(15, 6)
	fmt.Println(r)
	r, str := mult(5, 6)
	fmt.Println(r, str)
	r, b := divi(0, 15)
	if !b {
		fmt.Println("尝试除以0")
	}
	r, b = divi(15, 3)
	fmt.Println(r, b)
	r = sumAll(1, 2, 3, 4, 5, 6)
	fmt.Println(r)
	fmt.Println(sumAll(5, 6))
}
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}
func subtr(x int, y int) int {
	re := x - y
	return re
}
func mult(x int, y int) (int, string) { //多个返回值
	re := x * y
	return re, "乘法"
}
func divi(x, y int) (int, bool) { //连续多个相同类型，可以写一个
	if x == 0 {
		return x, false
	} else {
		re := x / y
		return re, true
	}
}
func sumAll(x, y int, o ...int) (ret int) {
	ret = x + y
	for _, a := range o {
		ret += a
	}
	fmt.Printf("%T\n", o)
	return
}
