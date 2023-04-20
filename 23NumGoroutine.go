package main

import "runtime"

func main() {
	go print()
	print()
	println(runtime.NumGoroutine()) // 2
}
func print2() {

}
