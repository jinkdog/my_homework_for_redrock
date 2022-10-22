package main

import (
	"fmt"
)

//	func add(a int, b int) {
//		fmt.Println(a + b)
//	}
//
//	func sub(a, b int) int {
//		return a - b
//	}
//
//	func mul(a, b int) int {
//		return a * b
//	}
//
//	func div(a, b int) int {
//		return a / b
//	}
func cal() {

	//switch cmd {
	//case "+":
	//	ret := a + b
	//	fmt.Println(ret)
	//	//result = append(result, a+b)
	//	//fmt.Println(result)
	//case "-":
	//	ret := a - b
	//	fmt.Println(ret)
	//	//result = append(result, a-b)
	//	//fmt.Println(result)
	//
	//case "*":
	//	ret := a * b
	//	fmt.Println(ret)
	//	//result = append(result, a*b)
	//	//fmt.Println(result)
	//
	//case "/":
	//	ret := a / b
	//	fmt.Println(ret)
	//	//result = append(result, a/b)
	//	//fmt.Println(result)
	//
	//default:
	//	fmt.Println("请输入正确表达式")
	//
	//}
	N := 1
	var result = make([]int, 0, N)
	for { //for i := N; i <= N; i++
		var a int
		_, _ = fmt.Scanln(&a)
		var cmd string
		_, _ = fmt.Scanln(&cmd)
		var b int
		_, _ = fmt.Scanln(&b)
		//_, _ = fmt.Scanf("%d%s%d", &a, &cmd, &b) //使用这种方式输入数据导致一直出现“请输入正确表达式”
		//c := cal(a, cmd, b)
		//result = append(result, c)
		//fmt.Println(result)
		//c := cal(a, cmd, b)
		//fmt.Println(c)
		//var num []int
		//num = make([]int, 100, 100)
		//num2 := []int{c}
		//num = append(num2, c)
		//fmt.Println(num2)
		//num0 := [...]int{}
		//slice := num0[:]
		//slice = append(slice, c)
		//fmt.Println(slice)
		switch cmd {
		case "+":
			//ret := a + b
			//fmt.Println(ret)
			result = append(result, a+b)
			fmt.Println(result)
		case "-":
			//ret := a - b
			//fmt.Println(ret)
			result = append(result, a-b)
			fmt.Println(result)

		case "*":
			//ret := a * b
			//fmt.Println(ret)
			result = append(result, a*b)
			fmt.Println(result)

		case "/":
			//ret := a / b
			//fmt.Println(ret)
			result = append(result, a/b)
			fmt.Println(result)

		default:
			fmt.Println("请输入正确表达式")
			break
		}
		//fmt.Println(result)
		N++
	}
}

// 它要输入一个运算符来表示a+b等，目前还不清楚该如何表示
// switch的基本语法的使用也不清晰
func main() {
	//N := 1
	//var result = make([]int, 0, N)
	//for { //for i := N; i <= N; i++
	//	var a int
	//	_, _ = fmt.Scanln(&a)
	//	var cmd string
	//	_, _ = fmt.Scanln(&cmd)
	//	var b int
	//	_, _ = fmt.Scanln(&b)
	//	//_, _ = fmt.Scanf("%d%s%d", &a, &cmd, &b) //使用这种方式输入数据导致一直出现“请输入正确表达式”
	//	//c := cal(a, cmd, b)
	//	//result = append(result, c)
	//	//fmt.Println(result)
	//	//c := cal(a, cmd, b)
	//	//fmt.Println(c)
	//	//var num []int
	//	//num = make([]int, 100, 100)
	//	//num2 := []int{c}
	//	//num = append(num2, c)
	//	//fmt.Println(num2)
	//	//num0 := [...]int{}
	//	//slice := num0[:]
	//	//slice = append(slice, c)
	//	//fmt.Println(slice)
	//	switch cmd {
	//	case "+":
	//		//ret := a + b
	//		//fmt.Println(ret)
	//		result = append(result, a+b)
	//		fmt.Println(result)
	//	case "-":
	//		//ret := a - b
	//		//fmt.Println(ret)
	//		result = append(result, a-b)
	//		fmt.Println(result)
	//
	//	case "*":
	//		//ret := a * b
	//		//fmt.Println(ret)
	//		result = append(result, a*b)
	//		fmt.Println(result)
	//
	//	case "/":
	//		//ret := a / b
	//		//fmt.Println(ret)
	//		result = append(result, a/b)
	//		fmt.Println(result)
	//
	//	default:
	//		fmt.Println("请输入正确表达式")
	//		break
	//	}
	//	//fmt.Println(result)
	//	N++
	//}
	cal()
}
