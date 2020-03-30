package main

import "fmt"

//声明变量
// var a int
// var b bool
// var c string

//批量声明
//go 推荐使用小驼峰式命名： studentName
var (
	name string
	age  int
	isok bool
)

//外部定义一定要以，关键词开头： var const等
func main() {
	name = "jubo"
	age = 27
	isok = true
	//var heiheihei string
	//go语言很不推荐变量声明后不使用，在方法内声明的局部变量，如不使用会编译不通过。
	fmt.Printf("name:%s", name)
	fmt.Println(age)
	fmt.Print(isok)

	//声明变量同时赋值
	var a string = "jubo"
	fmt.Println(a)
	//类型推导，后面的值是可以知道类型的。
	var b = "jubo"
	fmt.Println(b)
	//短声明，仅函数内的变量声明可以使用
	c := "jubo"
	fmt.Println(c)
	//匿名变量： _    ,也叫哑元变量，可以重复声明，变量不会使用仅用来接收不需要的函数返回值。
}
