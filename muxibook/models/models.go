package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	//tablename auto-create to be "users"
	ID            int         `gorm:"primary_key"`
	//Model include : ID CreateAt UpdateAt DeleteAt
	Username      string      `gorm:"size:20;unique"`//username=db.Column(db.String(20),unique=True)
	Password      string      `gorm:"size:20"`
//	PasswordHash  string  	  `gorm:"size:128"`//column name is password_hash
	Confirmed     bool        `gorm:"default:False"`
	BookCount     int         `gorm:"default:False"`

	Books         [] Book      `gorm:"ForeignKey:UserID"`
}

type Kind struct {
	//tablename auto-create to be "kinds"
	ID            uint        `gorm:"primary_key"`
	Books         [] *Book      `gorm:"ForeignKey:KindID"`
}

type Book struct {
	//tablename auto-create to be "books"
	gorm.Model
	Bookname      string     `gorm:"size:30"`
	Booknum       string     `gorm:"size:20;unique"`
	Available     uint       `gorm:"default:1"`
	Realname      string     `gorm:"size:20"`
	LendTime      time.Time  //think gorm.Model.xx later
	ReturnTime    time.Time

	KindID int
//	Kind   Kind
	UserID int
//	User   User
}

