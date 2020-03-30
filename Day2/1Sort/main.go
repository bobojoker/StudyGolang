package main

import "fmt"

func main() {
	a := []int{5, 4, 7, 6, 3}

	temp := 0
	for j := len(a) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if a[i] > a[i+1] {
				temp = a[i]
				a[i] = a[i+1]
				a[i+1] = temp
			}
		}
	}

	for _, t := range a {
		fmt.Printf("%d\n", t)
	}

}
