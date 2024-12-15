package main

import "fmt"

type WasRun struct {
	name   string
	wasRun bool
}

func NewWasRun(name string) WasRun {
	return WasRun{name: name, wasRun: false}
}

func (w *WasRun) Run() {
	w.TestMethod()
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}

func main() {
	test := NewWasRun("testMethod")
	fmt.Println(test.wasRun)
	test.Run()
	fmt.Println(test.wasRun)
}
