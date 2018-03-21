package main

import (
	"fmt"
)

//T 一个链表结构
type T struct {
	name string
	next *T
}

//TT 是T的别名，Golang1.9支持的新语法
type TT = T

//返回函数中分配的内存
func retMen() *[2]int {
	a := [2]int{1, 2}
	fmt.Printf("in p is: %p\n", &a)
	return &a
}

func main() {
	fmt.Println("Hello,Wrold!")
	//修改函数中返回的内存空间是否会报错
	a := retMen()
	//居然不报错。查了Golang官方FAQ，Golang会自动确定内存分配在Heap或者Stack上。这里显然是在Stack上，和C不同
	(*a)[0] = 100
	fmt.Printf("out p is: %p\n", a)
	fmt.Println(*a)

	//----------------------
	tail := TT{"tail", nil}
	head := TT{"head", &tail}

	for ptr := &head; ptr != nil; ptr = ptr.next {
		fmt.Println(ptr.name)
	}

}
