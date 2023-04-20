package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	start := time.Now()
	go func() {
		time.Sleep(3e9)
		println("123")
	}()

	defer fmt.Println(time.Since(start))
	runtime.Goexit()

}
