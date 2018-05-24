package api

import (
	"github.com/kataras/iris"
	"gowork/muxibook/models"

	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type Login struct {
	Username string
}

var DB *gorm.DB

func init(){
	var err error
	DB,err=gorm.Open("sqlite3","test.db")
	if err != nil{
		log.Fatalln(err)
	}
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Kind{})
}

func login(ctx iris.Context){
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
		DB.Create(&usr)
	}else{
		ctx.JSON(
			map[string]string{
				"token":strconv.Itoa(usr.ID)+usr.Username
			}
		)
	}
	ctx.StatusCode(200)
}

func signup(ctx iris.Context){
	var data Login
	var usr models.User
	err:=ctx.ReadJSON(&data)
	if err != nil{
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	if DB.Where("username=?",data.Username).First(&usr).RecordNotFound(){
		usr.Password="muxiboook"
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

