package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	//m指代的不是目录cal，而是cal目录下所有go文件唯一的package:etc
	//一个目录下的go文件不能分属不同的package
	"github.com/j1mh0/ho/etc"
	m "github.com/j1mh0/ho/etc/cal"
	// "runtime/pprof"
	"flag"
)

//Singleton 单例模式
type Singleton struct{}

var s *Singleton
var once sync.Once

func getSingleton() *Singleton {
	once.Do(func() {
		if s == nil {
			s = &Singleton{}
			fmt.Println("Singleton Created.")
		}
	})
	return s
}

//T 一个链表结构
type T struct {
	name string
	next *T
}

//Int is int type
type Int int

//SayHi is type Int's function
func (i *Int) SayHi() {
	fmt.Println("say hi in Int")
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

	//测试输入参数的解析
	config := flag.String("config", "./", "app name")
	dir := flag.String("dir", ".", "cureent dirctory")
	flag.Parse()
	fmt.Println(*config)
	fmt.Println(*dir)
	fmt.Println("Hello,Visual Studio Code.")

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
	//计算100个素数的函数，函数定义在prime.go里。这里注释是因为atom编辑器无法加载同包的另一个文件，go install没有问题
	//doPrime()

	//单例模式
	n := getSingleton()
	fmt.Printf("1st singleton return pointer is %p.\n", n)
	p := getSingleton()
	fmt.Printf("2nd singleton return pointer is %p.\n", p)

	//go 里的i变量很可能是无序的，因为没有值copy
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("inner count:%v\n", i)
		}()
	}

	//空指针的方法调用测试
	var iInt *Int = nil
	iInt.SayHi()

	//------------------------
	sa := make([]int, 5)
	sa = append(sa, 1, 2, 3)
	fmt.Println(sa)

	//非阻塞带缓冲的Channel
	tSig := make(chan int, 1)
	tSig <- 1
	for i := 0; i < 100; i++ {
		select {
		case c := <-tSig:
			fmt.Printf("signal is %v\n", c)
		default: //有default子句就不会阻塞
			fmt.Printf("loop #%v\n", i)
		}

	}

	//捕捉Ctrl+C后退出
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	select {
	case <-sig:
		fmt.Println("main program exit.")
	}

}
