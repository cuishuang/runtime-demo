//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//	pc, file, line, ok := runtime.Caller(0)
//	fmt.Println(pc)   // 4336410771
//	fmt.Println(file) // /Users/fliter/runtime-demo/16Caller.go
//	fmt.Println(line) //9
//	fmt.Println(ok)   //true
//}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc, _, line, _ := runtime.Caller(1)
	fmt.Printf("main函数的pc:%d\n", pc)      // main函数的pc:4364609931
	fmt.Printf("main函数被调用的行数:%d\n", line) // main函数被调用的行数:250
	show()
}
func show() {
	pc, _, line, _ := runtime.Caller(1)
	fmt.Printf("show函数的pc:%d\n", pc)      // show函数的pc:4364974271
	fmt.Printf("show函数被调用的行数:%d\n", line) // show函数被调用的行数:27
	// 这个是main函数的栈
	pc, _, line, _ = runtime.Caller(2)
	fmt.Printf("show的上层函数的pc:%d\n", pc)      // show的上层函数的pc:4364609931
	fmt.Printf("show的上层函数被调用的行数:%d\n", line) // show的上层函数被调用的行数:250
	pc, _, _, _ = runtime.Caller(3)
	fmt.Println(pc) //4364778899
	pc, _, _, _ = runtime.Caller(4)
	fmt.Println(pc) // 0
}
