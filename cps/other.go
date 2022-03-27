package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("starting")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("hello")
		}()
	}

	for {
		time.Sleep(time.Second)
	}
	fmt.Println("end")
}
