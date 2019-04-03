package main

import "fmt"

type Type int

type Set map[Type]bool

func (set *Set)AddItem(item Type) error {
    _, ok := *set[item]
    if ok {
        return errors.New("Item already existed!")
    } else {
        *set[item] = true
    }
    return nil
}

func (set *Set)Contain(item Type) bool {
    return _, ok = *set[item]; ok
}

func (set *Set)Delete(item Type) error {
    if _, ok = *set[item]; ok {
        delete(*set, item)
    } else {
        return errors.New("Item not found!")
    }
    return nil
}

func (set *Set)Size() int {
    return len(*set)
}

func (set *Set)Items() ([]Type, error) {
    var res []Type
    for key, _ := range(*set) {
        res = append(res, key)
    }
    return res, nil
}

func Add(setA, setB Set) Set {
    set := make(Set)
    for k, _ := range setA {
        set[k] = true
    }
    for k, _ := range setB {
        if _, ok := set[k], !ok {
            set[k] = true
        }
    }
    return set
}

func main() {
    fmt.Println()
}
