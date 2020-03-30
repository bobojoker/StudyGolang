package main

import (
	"fmt"
	"strings"
)

func main() {
	//正常的使用是一样的
	//多了一种 `` ，~按键下的带格式的 支付串
	s := `jubo
	hua
	shenqi
			shenme	ss
	jubo`
	fmt.Println(s) //原样输出
	s2 := `D:\Project\Service\GoWorkSpace\bin`
	fmt.Println(s2) //原样输出
	s = "jubo"
	//s2 = "shenme"
	//拼接
	//+ 拼接
	//String.Join()
	fmt.Printf("%s%s", s, s2)
	s3 := fmt.Sprintf("%s%s", s, s2) //返回一个string
	fmt.Println(s3)
	//分割 注意这里是strings
	ret := strings.Split(s2, "\\")
	fmt.Println(ret)
	//包含，返回bool Contains
	//前缀:HasPrefix,后缀:HasSuffix
	//获取位置 Index，倒数位置 LastIndex
	//获取长度len（string）
}
