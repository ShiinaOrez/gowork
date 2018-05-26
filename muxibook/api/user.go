package api

import (
	"github.com/kataras/iris"
	"gowork/muxibook/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"strconv"
	_ "fmt"
	"fmt"
)

type Login struct {
	Username string
}

var DB *gorm.DB

func init(){
	var err error
	DB,err=gorm.Open("sqlite3","muxibook.db")
	if err != nil{
		log.Fatalln(err)
	}
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Kind{})
}

func Signin(ctx iris.Context){
	var data Login
	var usr models.User
	err:=ctx.ReadJSON(&data)
	if err != nil{
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound(){
		usr.Username=data.Username
		usr.Password="muxibook"
		usr.BookCount=0
		DB.Create(&usr)
	}else{
		ctx.JSON(map[string]string{
				"token":strconv.Itoa(usr.ID)+usr.Username,
		})
	}
	ctx.StatusCode(200)
}

func Signup(ctx iris.Context){
	var data Login
	var usr models.User
	err:=ctx.ReadJSON(&data)
	if err != nil{
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound(){
		usr.Username=data.Username
		usr.Password="muxibook"
		usr.BookCount=0
		fmt.Println(usr.Username)
		DB.Create(&usr)
		ctx.JSON(map[string]string{
			"msg": "successful",
		})
		ctx.StatusCode(200)
	}else{
		ctx.StatusCode(401)
		return
	}
}

