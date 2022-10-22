package main

import (
	"fmt"
)

func main() {

	//var a = [...]string{}
	////var b string
	////_, _ = fmt.Scanln(&b)
	//_, _ = fmt.Scanln(&a)
	////fmt.Println(a)
	var a string
	_, _ = fmt.Scanln(&a)
	//fmt.Println(len(a))
	s1 := make([]rune, 0, len(a)) //len(a)得出字符串a的长度
	for _, h := range a {         //遍历a中的字符储存到切片s1中//这里为什么要用空白标识符？
		s1 = append(s1, h)
	}
	//fmt.Println(len(a))1个中文有3个长度
	//l := utf8.RuneCountInString(string(s1)) - 1
	l := len(a)/3 - 1
	//fmt.Println(l)
	//fmt.Println(s1)
	var b int
	for i := 0; i < l/2; i++ {
		//var j=[i]bool()
		if s1[i] != s1[l-i] { //如果l长度不-1此处S1[l-i]就会超过切片的范围
			fmt.Println("false")
			break //break //return此处使用break和return没有区别
		} else {
			b++
		}
	}
	if b == l/2 {
		fmt.Println(string(s1[:len(a)/6]))
	}
}

//要求全部判断完后再确定是对是错，不能判断一个，输出一个
//题目没有要求输出true，false只要有一个就停止
//想做出一个布尔型的数组，数组内的每个数储存一次循环后的true和false，如果全部是true则输出true，否则输出false
//实际上记录判断成功次数即可不需要再建立布尔类型数组
