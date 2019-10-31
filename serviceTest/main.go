package main

import (
	"fmt"
	"sync"
)

type List []int

func (list *List) GetItem(index int) int {
	return (*list)[index]
}

type Response struct {
	Lock *sync.Mutex
	List []int
}

func main() {
	list := List([]int{1, 2, 3, 4, 5, 6, 7})
	indexs := []int{0, 1, 2, 3, 4, 5, 6}

	wg := sync.WaitGroup{}
	finished := make(chan bool, 1)
	response := Response{
		Lock: new(sync.Mutex),
		List: []int{},
	}

	for _, index := range indexs {
		wg.Add(1)
		fmt.Println("2")
		go func(id int) {
			fmt.Println("Go")
			defer wg.Done()

			v := list.GetItem(id)
			fmt.Println(v)
			response.Lock.Lock()
			defer response.Lock.Unlock()

			response.List = append(response.List, v)
		}(index)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	fmt.Println(list)
	fmt.Println(response.List)
	return
}
