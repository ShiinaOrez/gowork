package main

import(
	"muxibook_model"
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
)

type AddBook struct{
	Kind  uint
	Bokname string
	Boknum  string
}

func check (e error){
	if e!= nil{
		panic(e)
	}
}

func main(){
	muxibook:=iris.New()
	db, err := gorm.Open("sqlite3", "muxibook.db")
	check(err)
	defer db.Close()
	muxibook.Post("/book",func (ctx iris.Context){
		addbook:=AddBook{}
		err:=ctx.ReadJSON(&addbook)
		check(err)
		kind:=addbook.Kind
		bokname:=addbook.Bokname
		boknum:=addbook.Boknum
		user:=muxibook_model.User{}
		db.Where("bookname = ?",bokname).First(&user)

	})
}