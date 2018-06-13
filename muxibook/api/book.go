package api

import (
	"github.com/kataras/iris"
	"gowork/muxibook/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	_ "strconv"
	_ "fmt"
	"strings"
	"time"
	_ "encoding/json"
	_ "fmt"
	"strconv"
)

type MyBook struct {
	Username string
}

type AddBook struct {
	Kind int    //the kind of this book
	Book string //book's name
	No   string //book's number
}

type SearchBook struct {
	Partten  string
	Page     int
}

type LendBook struct {
	Book     string  //book name
	Username string
	Realname string
}

type ReturnBook struct {
	No         string
	Username   string
}

type Monomer struct {
	Book       string    `json:"book"`
	Kind       int       `json:"kind"`
	Available  uint       `json:"available"`
	Who        string    `json:"who"`
	When       time.Time `json:"when"`
	Realname   string    `json:"realname"`
}

type LResponse struct {
	Book       string    `json:"book"`
	Kind       int       `json:"kind"`
	Available  int       `json:"available"`
	Who        string    `json:"who"`
	When       time.Time `json:"when"`
	Realname   string    `json:"realname"`
}

type SResponse struct {
	No         string    `json:"no"`
	BookName   string       `json:"bookname"`
	ReturnTime time.Time `json:"returntime"`
}

type MResponse struct {
	Lend     []SResponse `json:"lend"`
}

type BResponse struct {
	Num        int        `json:"num"`
	Page       int        `json:"page"`
	Books      []Monomer  `json:"books"`
}

func init(){
	var err error
	DB,err=gorm.Open("sqlite3","muxibook.db")
	if err!= nil{
		log.Fatalln(err)
	}
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Kind{})
}

func Book (ctx iris.Context){
	var data AddBook
	var bok models.Book
	var t string
	err:=ctx.ReadJSON(&data)
	if err !=nil{
		ctx.StatusCode(iris.StatusBadRequest)//bad request 400
		ctx.WriteString(err.Error())
		return
	}
	t=ctx.GetHeader("token")
	if t == "" {
		ctx.StatusCode(401) //confirm fail
		ctx.JSON(map[string]string{
			"msg": "can't confirm!",
		})
		return
	}else{
		if DB.Where("bookname=?",data.Book).First(&bok).RecordNotFound() && DB.Where("booknum=?",data.No).First(&bok).RecordNotFound(){
			bok.Bookname = data.Book
			bok.Booknum = data.No
			bok.KindID = data.Kind
			bok.Available=1
			DB.Create(&bok)
			ctx.JSON(map[string]string{
				"msg": "add successful!",
			})
			return
		}
	}
	ctx.StatusCode(402)
	ctx.JSON(map[string]string{
		"msg": "can't re-add!",
	})
	return
}

func BookLend (ctx iris.Context){
	var data LendBook
	var usr models.User
	var bok models.Book
	err:=ctx.ReadJSON(&data)
	if err !=nil{
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	t:=ctx.GetHeader("token")
	if !strings.Contains(t,data.Username) {
		ctx.StatusCode(401) //confirm fail
		ctx.JSON(map[string]string{
			"msg": "can't confirm!",
		})
		return
	}
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound() || DB.Where("bookname=?",data.Book).First(&bok).RecordNotFound(){
		ctx.StatusCode(402)
		ctx.JSON(map[string]string{
			"msg": "user or book not found!",
		})
		return
	}
	if usr.BookCount>=5{
		ctx.StatusCode(403)
		ctx.JSON(map[string]string{
			"msg": "user's bookcount >= 5!",
		})
		return
	}
	if bok.Available==0{
		ctx.StatusCode(407)
		ctx.JSON(map[string]string{
			"msg": "book not available!",
		})
	}
	bok.Available=0
	bok.UserID=usr.ID
//	DB.Model(&usr).Related(&usr.Books)
	bok.Realname=data.Realname
	usr.BookCount++
	DB.Save(&bok)
	DB.Save(&usr)
	lendtime:=bok.UpdatedAt
	mm,err:=time.ParseDuration("1h")
	returntime:=lendtime.Add(720*mm)
	bok.LendTime=lendtime
	bok.ReturnTime=returntime
	DB.Save(&bok)
	response:=LResponse{data.Book,bok.KindID,0,data.Username,returntime,data.Realname}
//	lresponse,err:=json.Marshal(response)

	ctx.JSON(response)
//	fmt.Print(bok)
	return
}

func BookReturn(ctx iris.Context){
	var data ReturnBook
	var usr models.User
	var bok models.Book
	err:=ctx.ReadJSON(&data)
	if err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.WriteString(err.Error())
	}
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound() || DB.Where("booknum=?",data.No).First(&bok).RecordNotFound(){
		ctx.StatusCode(402)
		ctx.JSON(map[string]string{
			"msg": "user or book not found!",
		})
		return
	}
	if bok.Available == 1{
		ctx.StatusCode(403)
		ctx.JSON(map[string]string{
			"msg": "this book is available!",
		})
		return
	}
	if bok.UserID != usr.ID {
		ctx.StatusCode(407)
		ctx.JSON(map[string]string{
			"msg": "you can't return other user's book!",
		})
	}
	usr.BookCount-=1
	bok.UserID=-1
//	DB.Model(&usr).Related(&usr.Books)
	bok.Available=1
	bok.ReturnTime=time.Time{}
	bok.LendTime=time.Time{}
	DB.Save(&usr)
	DB.Save(&bok)
	ctx.JSON(map[string]string{
        "msg": "successful",
	})
	return 
}

func BookRenew(ctx iris.Context){
	var data ReturnBook
	var usr models.User
	var bok models.Book
	err:=ctx.ReadJSON(&data)
	if err != nil{
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.WriteString(err.Error())
	}
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound() || DB.Where("booknum=?",data.No).First(&bok).RecordNotFound(){
		ctx.StatusCode(402)
		ctx.JSON(map[string]string{
            "msg": "user or book not found!",
		})
		return 
	}
	if bok.LendTime.Before(time.Now()) && bok.ReturnTime.After(time.Now()) {
		bok.LendTime=time.Now()
		DB.Save(&bok)
		lendtime:=bok.UpdatedAt
	    mm,_:=time.ParseDuration("1h")
	    returntime:=lendtime.Add(720*mm)
	    bok.ReturnTime=returntime
	    DB.Save(&bok)
	    ctx.JSON(map[string]string{
            "msg": "successful!",
	    })
	    return
	}else{
		ctx.StatusCode(401)
		ctx.JSON(map[string]string{
            "msg": "you aren't reach the time!",
		})
		return
	}
}

func BookMy(ctx iris.Context){
	var data MyBook
	var usr models.User
	var response MResponse
	t:=ctx.GetHeader("token")
	if !strings.Contains(t,data.Username) {
		ctx.StatusCode(401) //confirm fail
		ctx.JSON(map[string]string{
			"msg": "can't confirm!",
		})
		return
	}
	ctx.ReadJSON(&data)
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound(){
		ctx.StatusCode(402)
		ctx.JSON(map[string]string{
			"msg": "user not found!",
		})
	}
	var boks[] models.Book
	DB.Model(&usr).Related(&boks)
    for _,bok:=range boks{
    	one:=SResponse{bok.Booknum,bok.Bookname,bok.ReturnTime}
    	response.Lend=append(response.Lend,one)
	}
//	fmt.Print(usr.Books)
//    fmt.Print(usr)
	ctx.JSON(response)
    return
}


func BookGet(ctx iris.Context){
	var response BResponse
	knd:=ctx.URLParam("kind")
	pge:=ctx.URLParam("page")
	page,_:=strconv.Atoi(pge)
	kind,_:=strconv.Atoi(knd)
	response.Page=page
    var kind_books[] models.Book
    U:=models.User{}
    DB.Where(&models.Book{KindID:kind}).Find(&kind_books)
	response.Num=len(kind_books)
	start:=(page-1)*10
	now:=start
	var usr models.User
	if kind_books==nil {
		ctx.StatusCode(401)
		ctx.JSON(map[string]string{
			"msg": "No matching book!",
		})
		return
	}
	for now<start+10 && now<len(kind_books){
		if DB.Where(&models.User{ID:kind_books[now].UserID}).First(&usr).RecordNotFound(){
			usr=U
		}
		one:=Monomer{kind_books[now].Bookname,kind,kind_books[now].Available,usr.Username,kind_books[now].ReturnTime,kind_books[now].Realname}
		response.Books=append(response.Books,one)
		now+=1
	}
	ctx.JSON(response)
	return
}

func Search(ctx iris.Context){
	var response BResponse
	data:=SearchBook{}
	ctx.ReadJSON(&data)
	var books[] models.Book
	DB.Where("bookname LIKE ?","%"+data.Partten+"%").Find(&books)
	if books==nil {
		ctx.StatusCode(401)
		ctx.JSON(map[string]string{
			"msg": "No matching book!",
		})
		return
	}
	var usr models.User
	U:=models.User{}
	response.Num=len(books)
	start:=(data.Page-1)*10
	now:=start
	for now<start+10 && now<len(books){
		if DB.Where(&models.User{ID:books[now].UserID}).First(&usr).RecordNotFound(){
			usr=U
		}
		one:=Monomer{books[now].Bookname,books[now].KindID,books[now].Available,usr.Username,books[now].ReturnTime,books[now].Realname}
		response.Books=append(response.Books,one)
		now+=1
	}
	ctx.JSON(response)
	return
}




