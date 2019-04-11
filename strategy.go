package main

import "fmt"

// 策略模式
type behaviour interface {
	Do()
}

type fly struct {
}

func (f fly) Do() {
	fmt.Println("Fly")
}

type run struct{}

func (r run) Do() {
	fmt.Println("Run")
}

type animal struct {
	b behaviour
}

func (a *animal) setBehaviour(b behaviour) {
	a.b = b
}
