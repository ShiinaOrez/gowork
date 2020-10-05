package main

func main() {
    var i interface{}
    var a int64 = 0
    i = a
    if i == 0 {
        println("equal")
        return
    }
    println("not equal")
    return
}
