package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)
	startTime := time.Now()
	var s1 chan int64 = make(chan int64)
	var s2 chan int64 = make(chan int64)
	var s3 chan int64 = make(chan int64)
	var s4 chan int64 = make(chan int64)
	go calc(s1)
	go calc(s2)
	go calc(s3)
	go calc(s4)
	<-s1
	<-s2
	<-s3
	<-s4
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime)) // 第一行注释掉 耗时: 386.954625ms； 取消注释 耗时: 1.34715s

}
func calc(s chan int64) {
	var count int64 = 0
	for i := 0; i < 1000000000; i++ {
		count += int64(i)
	}
	s <- count
}
