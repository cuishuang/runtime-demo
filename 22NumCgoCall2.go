package main

import (
	"runtime"
)

/*
   #include <stdio.h>
   // 自定义一个c语言的方法
   static void myPrint(const char* msg) {
     printf("myPrint: %s", msg);
   }
*/
import "C"

func main() {
	// 调用c方法
	C.myPrint(C.CString("Hello,C\n"))
	println(runtime.NumCgoCall())
}
