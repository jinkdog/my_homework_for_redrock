package main

import "fmt"

func main() {
	N := 1
	var result = make([]int, 0, N)
	for {
		var a int
		_, _ = fmt.Scanln(&a)
		var cmd string
		_, _ = fmt.Scanln(&cmd)
		var b int
		_, _ = fmt.Scanln(&b)
		if a+b == a+b {
			result = append(result, a+b)
			fmt.Println(result)
		} else if a-b == a-b {
			result = append(result, a-b)
			fmt.Println(result)
		} else if a*b == a*b {
			result = append(result, a*b)
			fmt.Println(result)
		} else if a/b == a/b {
			result = append(result, a/b)
			fmt.Println(result)
		}
		N++
	}
}
