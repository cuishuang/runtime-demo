package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go calcSum1()
	go calcSum2()
	time.Sleep(time.Second * 10)
}

func calcSum1() {
	runtime.LockOSThread()
	start := time.Now()
	count := 0
	for i := 0; i < 10000000000; i++ {
		count += i
	}
	end := time.Now()
	fmt.Println("calcSum1耗时")
	fmt.Println(end.Sub(start))
	defer runtime.UnlockOSThread()
}

func calcSum2() {
	start := time.Now()
	count := 0
	for i := 0; i < 10000000000; i++ {
		count += i
	}
	end := time.Now()
	fmt.Println("calcSum2耗时")
	fmt.Println(end.Sub(start))
}
