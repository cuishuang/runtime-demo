package main

import (
	"fmt"
	"runtime"
)

func main() {
	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs)
	fmt.Println(pcs[:i]) // [4311883569 4311525404 4311694372]
}
