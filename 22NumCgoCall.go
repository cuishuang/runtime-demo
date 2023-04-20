package main

import (
	"runtime"
)

/*
#include <stdio.h>
*/
import "C"

func main() {
	println(runtime.NumCgoCall()) // 1
}
