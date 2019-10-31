package main

import "fmt"

type RouterGroup struct {
	engine *Engine
}

type Engine struct {
	RouterGroup
}

func main() {
	rg := RouterGroup{}
	fmt.Println(rg.engine)
}
