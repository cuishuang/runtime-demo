package main

import (
	"runtime"
)

func main() {

	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs)
	for _, pc := range pcs[:i] {
		println(runtime.FuncForPC(pc))
	}
}
