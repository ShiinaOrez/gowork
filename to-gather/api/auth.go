package api

import (
	_ "fmt"
	"github.com/ShiinaOrez/gowork/to-gather/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
	"log"
	"strconv"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "to-gather.db")
	if err != nil {
		log.Fatalln(err)
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Activity{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Picker2Activity{})
}

func Login(ctx iris.Context) {
	var data LoginPostData
	var usr models.User
	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	if DB.Where("std_num=?", data.StdNum).First(&usr).RecordNotFound() {
		//add a new record into User
		usr.Name = data.Username
		usr.StdNum = data.StdNum
		DB.Create(&usr)
	}
	ReturnData := LoginReturnData{
		strconv.Itoa(usr.ID) + "?" + usr.Name,
		usr.ID,
		usr.StdNum,
	}
	ctx.StatusCode(200)
	ctx.JSON(map[string]LoginReturnData{
		"Data": ReturnData,
	})
	return
}
