package main

import (
	"fmt"
	"runtime"
)

func main() {

	for i := 0; i < 10; i++ {

		go func(k int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("----------------")
	p := make([]runtime.StackRecord, 10000)

	fmt.Println(runtime.GoroutineProfile(p)) // 1 true
}
