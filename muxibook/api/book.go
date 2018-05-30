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
)

type AddBook struct {
	Kind int    //the kind of this book
	Book string //book's name
	No   string //book's number
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

type LResponse struct {
	Book       string    `json:"book"`
	Kind       int       `json:"kind"`
	Available  int       `json:"available"`
	Who        string    `json:"who"`
	When       time.Time `json:"when"`
	Realname   string    `json:"realname"`
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
	if t=="" {
		ctx.StatusCode(402) //means not confirm
		ctx.JSON(map[string]string{
			"msg": "who are you?",
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
	ctx.StatusCode(401)
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
		ctx.StatusCode(405)
		ctx.JSON(map[string]string{
			"msg": "book not available!",
		})
	}
	bok.Available=0
	bok.UserID=usr.ID
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
	usr.BookCount-=1
	bok.UserID=0
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


