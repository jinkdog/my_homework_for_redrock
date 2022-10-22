package main

import (
	"math/rand"
	"time"
)

func bubblesort(A []int, N int) {
	var flag bool
	flag = true
	for i := 0; flag; i++ {
		flag = false
		for j := N - 1; j >= i+1; j-- {
			if A[j] < A[j-1] {
				temp := A[j]
				A[j] = A[j-1]
				A[j-1] = temp
				flag = true
			}
		}
	}
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Intn(100)
	for key := range r {
		
	}
}
