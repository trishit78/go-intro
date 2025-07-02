package main

import "fmt"

type complexOps interface {
	add()
	subtract()
}

type complex struct {
	real int
	img  int
}

func newComplexNo(real int, img int) *complex {
	c := complex{
		real: real,
		img:  img,
	}
	return &c

}

func printOps(num1 complexOps, num2 complexOps) {
	addNum := num1.add
	fmt.Println("addition:", addNum)

	subNum := num.subtract()
	fmt.Println("substraction:", subNum)

}

func (n *complex) add() {

}

func main() {
	fmt.Println("hello")

	num := newComplexNo(3, 4)
	fmt.Println("real no", num.real)
	fmt.Println("img no", num.img)
}

//build systems
