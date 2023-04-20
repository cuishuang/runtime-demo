package main

import (
	"runtime"
	"time"
)

type Student struct {
	name string
}

func main() {
	var i *Student = new(Student)
	runtime.SetFinalizer(i, func(i interface{}) {
		println("垃圾回收了") // 垃圾回收了
	})
	runtime.GC() // 如果将这行注释，则上面不会输出
	time.Sleep(time.Second)
}
