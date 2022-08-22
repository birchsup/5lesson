package main

import (
	"fmt"
)

func fibonacci() func() int {
	FibMap := make(map[int]int)
	n := 0
	return func() int {
		if n == 0 {
			FibMap[n] = 0
			n++
			return 0
		}
		if n == 1 {
			FibMap[n] = 1
			n++
			return 1
		}
		number := FibMap[n-1] + FibMap[n-2]
		FibMap[n] = number
		n++
		fmt.Println(FibMap)
		return number

	}
}

func main() {

	var n int
	fmt.Scanln(&n)

	f := fibonacci()

	for i := 0; i < n+1; i++ {
		fmt.Println(f())

	}
}
