/* 限制并发数目*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Glimit struct {
	n int
	c chan struct{}
}

func NewGlimit(n int) (g *Glimit) {
	g = &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}

	return g
}

func (g *Glimit) Run(f func(), wg *sync.WaitGroup) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
		wg.Done()
	}()
}

func main() {
	number := 10
	g := NewGlimit(3)
	var wg = sync.WaitGroup{}
	wg.Add(number)
	for i := 0; i < number; i++ {
		value := i
		g.Run(func() {
			fmt.Println("HELLO,run func：", value)
			time.Sleep(time.Duration(10) * time.Second)
		}, &wg)
	}
	wg.Wait()
}
