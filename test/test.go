package main

import "fmt"

type Eyes  string // 眼睛类型
type Mouth string // 嘴类型
type Nose  string // 鼻子类型
type Ears  string // 耳朵类型

func (ears *Ears) Hear(s string) {
    fmt.Println(s)
}

func (eyes *Eyes) Hear(s string) {
    fmt.Println(s)
}

type Face struct { // 脸类型，由各个部分拼接而成
    Eyes
    Mouth
    Nose
    Ears
}

func main() {
    face := Face{}
    face.Hear("Golang is the best!")
}
