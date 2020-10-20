package main

import (
	"fmt"
)

type Person struct {
	age   int
	class *Class
}

type Class struct {
	name string
}

func (p *Person) changeClass() {
	p.class.name = "class_three"
}

func (p *Person) printAge() {
	fmt.Println("print age") // console can print this sentence 
	fmt.Println(p.age) // panic runtime
}

func main() {
	var p *Person = nil
	p.printAge()
}
