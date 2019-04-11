package main

import "fmt"

// 观察者模式

// 订阅者
type subscribe interface {
	Update()
}

type subscribeA struct{}

func (a *subscribeA) Update() {
	fmt.Println("subscribeA update")
}

type subscribeB struct{}

func (b *subscribeB) Update() {
	fmt.Println("subscribeB update")
}

// 生产者
type observer struct {
	observerList map[string]subscribe
}

func (o *observer) addObserver(key string, val subscribe) {
	o.observerList[key] = val
}

func (o *observer) removeObserver(key string) {
	delete(o.observerList, key)
}

func (o *observer) notify() {
	for _, item := range o.observerList {
		item.Update()
	}
}

func (o *observer) updateInfo() {
	o.notify()
}
