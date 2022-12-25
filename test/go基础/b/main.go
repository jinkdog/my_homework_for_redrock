package main

import "fmt"

// 输出1的匿名函数延迟执行
// 由于a恒为true导致输出2的函数立即执行在return完成前执行输出1的函数
// 输出3的函数在return之后开始执行defer但由于主goroutine结束导致不再继续执行输出3的函数
func main() {
	var a = true
	defer func() {
		fmt.Println("并发关联")
	}()

	if a {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
