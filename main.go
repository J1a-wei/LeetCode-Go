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
	fmt.Println(p.age)       // panic runtime
}

// 访问闭包元素，例如slice，是否影响外层
// s 的引用没有问题
//func main() {
//	s := []int{}
//
//	var test func(int)
//	test = func(i int) {
//		if i == 5 {
//			return
//		}
//		s = append(s, i)
//		test(i + 1)
//	}
//
//	test(0)
//	fmt.Println(s)
//
//	// [0 1 2 3 4]
//
//}

//func main() {
//	x := 1
//
//	var change func(int)
//	change = func(i int) {
//		x = i
//	}
//	change(3)
//	fmt.Println(x)
//	3
//
//}

//func main() {
//	s := []int{}
//	fmt.Printf("%p\n", &s)
//	var change func([]int)
//	change = func(x []int) {
//		fmt.Printf("%p\n", &x)
//		x = append(x, 1, 2, 3, 4, 5) // case1 外层slice 并没有添加成功
//		fmt.Printf("%p\n", &x)
//		//s = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // case2 外层slice添加成功了
//	}
//	change(s)
//	//fmt.Println(s)
//	fmt.Printf("%p\n", &s)
//}

func main() {
	s := []int{}
	fmt.Printf("%p\n", &s) // 与下面的相等
	testAppend(s)
	fmt.Printf("%p\n", &s) // 与上面的相等
	fmt.Println(s)
}

func testAppend(x []int) {
	fmt.Printf("%p\n", &x) // 带了取地址符号是相等的，如果不带取地址符号，后面append重新allocate一个数组
	x = append(x, 0)
	fmt.Printf("%p\n", &x)
}
