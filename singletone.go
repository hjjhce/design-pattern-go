package main

import "sync"

// Singletone 单例模式
type Singletone struct {
}

var singletone *Singletone
var once sync.Once

func getInstance() *Singletone {

	once.Do(func() { // Do当且仅当第一次调用时才执行
		if singletone == nil {
			singletone = &Singletone{}
		}
	})

	return singletone
}
