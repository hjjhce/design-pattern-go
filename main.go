package main

import "fmt"

func main() {

	// 策略模式
	/*
		f := fly{}
		r := run{}

		a := &animal{}
		a.setBehaviour(f)
		a.b.Do()
		a.setBehaviour(r)
		a.b.Do()
	*/

	// 观察者模式
	/*
		a := &subscribeA{}
		b := &subscribeB{}

		o := &observer{}
		o.observerList = make(map[string]subscribe)
		o.addObserver("subscribeA", a)
		o.addObserver("subscribeB", b)
		o.updateInfo()
		o.removeObserver("subscribeB")
		o.updateInfo()
	*/

	// 装饰者模式
	r := &rice{}
	e := &eggRice{rice: r}
	b := &beefEggRice{eggRice: e}

	fmt.Println(r.price())
	fmt.Println(r.getDesc())

	fmt.Println(e.price())
	fmt.Println(e.getDesc())

	fmt.Println(b.price())
	fmt.Println(b.getDesc())

}
