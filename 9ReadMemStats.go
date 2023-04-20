package main

import (
	"fmt"
	"runtime"
	"time"
)

type Student2 struct {
	name string
}

func main() {
	var list = make([]*Student2, 0)
	for i := 0; i < 100000; i++ {
		var s *Student2 = new(Student2)
		list = append(list, s)
	}
	memStatus := runtime.MemStats{}
	runtime.ReadMemStats(&memStatus)
	fmt.Printf("申请的内存:%d\n", memStatus.Mallocs) // 申请的内存:100250

	fmt.Printf("释放的内存次数:%d\n", memStatus.Frees) // 释放的内存次数:45

	time.Sleep(time.Second)
}
