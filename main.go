package main

import (
	"fmt"
	//m指代的不是目录cal，而是cal目录下所有go文件唯一的package:etc
	//一个目录下的go文件不能分属不同的package
	"github.com/j1mh0/ho/etc"
	m "github.com/j1mh0/ho/etc/cal"
	// "runtime/pprof"
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

	// f, err := os.Create("cpu")
	// if err == nil {
	// 	pprof.StartCPUProfile(f)
	// 	defer pprof.StopCPUProfile()
	// }

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

	//实验包名和目录名不同。import的是目录，而使用的是包名
	//一个目录下的go文件不能分属两个以上不同的package(子目录除外)
	display.SayHello()
	fmt.Println(m.Sum(1, 3))

}
