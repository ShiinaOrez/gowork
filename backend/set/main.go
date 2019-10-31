package main

import (
	"errors"
	"fmt"
)

type Type int

type Set map[Type]bool

func (set *Set) AddItem(item Type) error {
	_, ok := (*set)[item]
	if ok {
		return errors.New("Item already existed!")
	} else {
		(*set)[item] = true
	}
	return nil
}

func (set *Set) Contain(item Type) bool {
	_, ok := (*set)[item]
	return ok
}

func (set *Set) Delete(item Type) error {
	if _, ok := (*set)[item]; ok {
		delete(*set, item)
	} else {
		return errors.New("Item not found!")
	}
	return nil
}

func (set *Set) Size() int {
	return len(*set)
}

func (set *Set) Items() ([]Type, error) {
	var res []Type
	for key, _ := range *set {
		res = append(res, key)
	}
	return res, nil
}

func And(setA, setB Set) Set {
	set := make(Set)
	for k, _ := range setA {
		set[k] = true
	}
	for k, _ := range setB {
		if _, ok := set[k]; !ok {
			set[k] = true
		}
	}
	return set
}

func Or(setA, setB Set) Set {
	set := make(Set)
	for k, _ := range setA {
		if _, ok := setB[k]; ok {
			set[k] = true
		}
	}
	return set
}

func (setA *Set) Sub(setB Set) Set {
	set := *setA
	for k, _ := range setB {
		if _, ok := (*setA)[k]; ok {
			set.Delete(k)
		}
	}
	return set
}

func test() {
	seta := make(Set)
	setb := make(Set)
	seta.AddItem(1)
	seta.AddItem(3)
	seta.AddItem(5)
	setb.AddItem(2)
	setb.AddItem(4)
	setb.AddItem(6)
	fmt.Println("seta:", seta)
	fmt.Println("setb:", setb)
	fmt.Println("seta and setb:", And(seta, setb))
	fmt.Println("seta or setb:", Or(seta, setb))
	fmt.Println("seta sub setb:", seta.Sub(setb))
}

func main() {
	test()
}
