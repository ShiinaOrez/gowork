package main

import (
    "github.com/kataras/iris"
    "fmt"
)

type Company struct {
    Name  string
    City  string
    Other string
}

func MyHandler(ctx iris.Context){
    var c Company
    if err :=ctx.ReadJSON(&c);err!=nil{
        ctx.StatusCode(iris.StatusBadRequest)
        //same like ctx.StatusCode(400)
        ctx.WriteString(err.Error())
        return
    }
    fmt.Println("the name is :",c.Name)
    fmt.Println("the city is :",c.City)
    fmt.Println("the other is :",c.Other)
    ctx.Writef("Received: %#+v\n",c)
}

type Person struct{
    Name string `json:"name"`
    Age  string `json:"age"`
}

func MyHandler2(ctx iris.Context){
    var persons []Person
    err:=ctx.ReadJSON(&persons)
    if err!=nil{
        ctx.StatusCode(iris.StatusBadRequest)
        //same like ctx.StatusCode(400)
        ctx.WriteString(err.Error())
        return
    }
    ctx.Writef("Received:%#+v\n",persons)
}

func main(){
    app:=iris.New()
    app.Post("/",MyHandler)
    app.Post("/slice",MyHandler2)

    app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed),iris.WithOptimizations)
}
