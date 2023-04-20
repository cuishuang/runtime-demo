package main

import (
	"fmt"
	"runtime"
)

func main() {
	go print25()
	runtime.Gosched()
	fmt.Println("继续执行")
}
func print25() {
	fmt.Println("执行打印方法")
}
