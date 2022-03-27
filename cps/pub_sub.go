/* 在发布订阅模型中，每条消息都会传送给多个订阅者。发布者通常不会知道、
也不关心哪一个订阅者正在接收主题消息。订阅者和发布者可以在运行时动态添加，是一种松散的耦合关系*/

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber   chan interface{}       //订阅者
	topicfileter func(interface{}) bool //主体过滤器是一个函数
)

type Publisher struct {
	buffer        int //缓存数目
	sub           map[subscriber]topicfileter
	pulishtimeout time.Duration //超时时间
	mux           sync.Mutex
}

func NewPublisher(n int, publishtimeout time.Duration) (p *Publisher) {
	p = &Publisher{
		buffer:        n,
		pulishtimeout: publishtimeout,
		sub:           make(map[subscriber]topicfileter),
	}

	return p
}

func (p *Publisher) Subscribe() subscriber {
	return p.SubscribeTopic(nil)
}

//订阅主题
func (p *Publisher) SubscribeTopic(filter topicfileter) subscriber {
	ch := make(subscriber, p.buffer)
	p.mux.Lock()
	defer p.mux.Unlock()
	p.sub[ch] = filter
	return ch
}

//退出订阅
func (p *Publisher) Evict(sub subscriber) {
	p.mux.Lock()
	defer p.mux.Unlock()
	delete(p.sub, sub)
	close(sub)
}

//发布
func (p *Publisher) Publish(v interface{}) {
	//fmt.Println("pulish :", v)
	p.mux.Lock()
	defer p.mux.Unlock()
	var wg sync.WaitGroup
	for s, topic := range p.sub {
		wg.Add(1)
		go p.sendTopic(s, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendTopic(sub subscriber, topic topicfileter, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
		time.Sleep(time.Duration(5) * time.Second)
	case <-time.After(p.pulishtimeout):
	}
}

// 关闭发布者对象，同时关闭所有的订阅者管道。
func (p *Publisher) Close() {
	p.mux.Lock()
	defer p.mux.Unlock()

	for sub, _ := range p.sub {
		delete(p.sub, sub)
		close(sub)
	}
}

func main() {

	p := NewPublisher(10, 100*time.Millisecond)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(10 * time.Second)

}
