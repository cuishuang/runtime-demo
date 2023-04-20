package main

import (
	"fmt"
	"runtime"
)

func main() {
	print()
	fmt.Println("继续执行")
}
func print() {
	fmt.Println("准备结束go协程")
	defer fmt.Println("结束了--会输出出来")
	runtime.Goexit()
	//defer fmt.Println("结束了")
}
