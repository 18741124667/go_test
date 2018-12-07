package main

import (
	"time"
	"runtime"
	"fmt"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				if i >= 10 {
					fmt.Println(">=10")
				}
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
