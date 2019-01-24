package api

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/ShiinaOrez/gowork/to-gather/models"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"strconv"
	_"fmt"
)

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
	if DB.Where("stdnum=?", data.Username).First(&usr).RecordNotFound() {
		//add a new record into User
		usr.Name = data.Username
		usr.StdNum = data.StdNum
		DB.Create(&usr)

	}
	ReturnData := data_structure2.LoginReturnData {
		strconv.Itoa(usr.ID) + usr.Name,
		usr.ID,
		usr.StdNum,
	}
	ctx.StatusCode(200)
	ctx.JSON(map[string]LoginReturnData {
		"Data": ReturnData,
	})
	return
}