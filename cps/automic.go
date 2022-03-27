//原子操作

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Singleton struct{}

var (
	instance      *Singleton
	mx            sync.Mutex
	initialzation uint32
)

func Instance() *Singleton {
	if atomic.LoadUint32(&initialzation) == 1 {
		return instance
	}

	mx.Lock()
	defer mx.Unlock()
	if nil == instance {
		defer atomic.StoreUint32(&initialzation, 1)
		instance = &Singleton{}
	}

	return instance

}

func main() {
	exit := make(chan interface{})
	inst := Instance()
	fmt.Println(inst)
	exit <- 1
}
