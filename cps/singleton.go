//单例模式

package main

import (
	"sync"
)

type Instance struct {
	Val int
}

var (
	inst *Instance
	once sync.Once
	wg   sync.WaitGroup
)

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(a int) {
			once.Do(func() {
				inst = &Instance{100}
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
}
