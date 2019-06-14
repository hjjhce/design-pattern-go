package main

// 装饰者模式

type rice struct{}

func (f *rice) price() float64 {
	return 1
}

func (f *rice) getDesc() string {
	return "rice"
}

type eggRice struct {
	*rice
}

func (e *eggRice) price() float64 {
	return e.rice.price() + 2
}

func (e *eggRice) getDesc() string {
	return e.rice.getDesc() + ", add egg"
}

type beefEggRice struct {
	*eggRice
}

func (b *beefEggRice) price() float64 {
	return b.eggRice.price() + 10
}

func (b *beefEggRice) getDesc() string {
	return b.eggRice.getDesc() + ", add beef"
}
