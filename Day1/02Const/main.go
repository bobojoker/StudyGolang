package main

import "fmt"

//常量
//定义之后不能修改
//程序运行期间不会修改
const pi = 3.14159265359

const (
	a = "jubo"
	b = 15
	c = true
)

//不赋值的情况下，下一个常量和上一个值一样
const (
	n1 = 100
	n2
	n3
)

//iota,初始使用为0，没次声明一行变量加一
const (
	m1 = iota //0
	m2        //1
	m3        //2
)

//跳位
const (
	a1 = iota
	a2
	_
	a3
)

//插队
const (
	b1 = iota //0
	b2 = 100  //100
	b3 = iota //2  是有一行变量声明就是+1
	b4        //3
)

//多个常量声明在一行
const (
	c1, c2 = iota + 1, iota + 2 //1 2
	c3, c4 = iota + 1, iota + 2 //2 3
)

//定义数量级(iota的简单应用)
const (
	_  = iota
	kb = 1 << (10 * iota)
	mb
	gb
	tb
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)
	fmt.Println(tb)
}
