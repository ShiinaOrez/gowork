package main

import (
	"fmt"
)

func main(){
	fmt.Println("bi(a)tch")
	ch:=make(chan int,1)
	ch <- 2
	x:=<-ch
	<-ch
	close(ch)
	fmt.Println(x)
}