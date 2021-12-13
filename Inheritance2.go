package main

import "fmt"

type iBase interface {
	say()
}
type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("Hi Monika")
}

type child struct {
	base  //embedding
	style string
}

func check(b iBase) {
	b.say()
}
func main() {
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
	fmt.Println("The color is " + child.color)
	check(child)
}
