package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&n)
	var result = fn(a, b, n)
	fmt.Println(result)
}

func fn(a int, b int, n int) string {
	var result string
	var bb = b / n
	if b % n != 0 {
		bb += 1
	}
	if a > bb {
		result = "Yes"
	} else {
		result = "No"
	}
	return result
}

