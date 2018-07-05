package main

import "fmt"

func f1(){
	for{
		fmt.Println("call f1...")
	}
}

func f2(){
	fmt.Println("call f2...")
}

func main(){
	go f1()
	go f2()
	ch:=make(chan int)
	<-ch
	close(ch)
}