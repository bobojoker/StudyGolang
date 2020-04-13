package main

import (
	"fmt"
	"math/rand"
	"time"
)

//类似字典
func main() {
	m1 := make(map[string]int, 10)
	m1["jubo"] = 18
	m1["bobo"] = 35

	fmt.Println(m1)
	value, ok := m1["long"] //通用写法，ok作为bool
	if !ok {
		fmt.Println(value)
	}
	//遍历

	for k := range m1 { //遍历key
		fmt.Println(k)
	}
	for _, v := range m1 { //遍历value
		fmt.Println(v)
	}

	//删除
	delete(m1, "jubo")
	fmt.Println(m1)
	//随机数
	rand.Seed(time.Now().UnixNano())

	scoreMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu %02d ", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
}
