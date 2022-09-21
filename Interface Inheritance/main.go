package main

import (
	"fmt"
)

type AAA interface{
	A(string)
	B(string)
}

type BBB interface{
	C(string)
	D(string)
}

type Car struct {

}

func NewCar() interface{} {
	return &Car{}
}

func (c *Car) A(i string){
	fmt.Println(i)
}

func (c *Car) B(i string){
	fmt.Println(i)
}

func (c *Car) C(i string){
	fmt.Println(i)
}

func (c *Car) D(i string){
	fmt.Println(i)
}

func main() {
	c := NewCar()

	c.(AAA).A("Hello")
	c.(AAA).B("World")
	c.(BBB).C("Hello")
	c.(BBB).D("World")
}