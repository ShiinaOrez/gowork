package main

import (
	"log"
)

// 声明一个函数,
func add(a int, c chan int) {
	log.Println("准备往管道添加:", a)
	c <- a
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	// 声明了一个管道, 管道长度为0
	c := make(chan int)
	// 做两个循环, 开启多个协程
	for _, v := range a[:len(a)/2] {
		go add(v, c)
	}
	for _, v := range a[len(a)/2:] {
		go add(v, c)
	}
	//    defer close(c)
	// 上面的chan是阻塞的, 每当这里取出一个值, 上面的协程才会继续再执行一次
	var x, n int
	ok := true
	for ok {
		n, ok = <-c
		log.Println("从管道取出:", n, "并加入到x中", "Ok:", ok)
		x += n
	}
	log.Println("Sum:", x)
}
