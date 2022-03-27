package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producer(factor int, out chan<- int) {
	for i := 1; ; i++ {
		out <- factor * i
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	mes := make(chan int, 50)
	go producer(3, mes)
	go producer(5, mes)
	go consumer(mes)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit:%v\n", <-sig)
}
