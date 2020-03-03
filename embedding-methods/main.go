package main

type T interface {
    GetName() string
    E
}

type E interface {
    GetName() string
}

func main() {
    var _ T
}
