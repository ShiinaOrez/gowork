package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard //one to one

	Emails []Email //one to many

	BillingAddress   Address //one to one foriegn key
	BillingAddressID sql.NullInt64

	ShippingAddress   Address
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string
	Address2 string
	Post     sql.NullString
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:idx_name_code"`
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
