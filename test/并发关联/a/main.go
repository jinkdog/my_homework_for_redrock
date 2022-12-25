package main

import (
	"fmt"
)

var ch1 = make(chan string, 1)

func Work(WorkName string) {
	ch1 <- WorkName

}

func main() {
	go Work("goroutine1")
	fmt.Println(<-ch1)
	go Work("goroutine2")
	fmt.Println(<-ch1)
	go Work("goroutine3")
	fmt.Println(<-ch1)
	fmt.Println("success")
}
