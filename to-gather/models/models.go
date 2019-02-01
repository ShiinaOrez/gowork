package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "log"
    "time"
)

var DB *gorm.DB

func init() {
    var err error
    DB, err = gorm.Open("sqlite3", "to-gather.db")
    if err != nil {
        log.Fatalln(err)
    }
    DB.AutoMigrate(&User{})
    DB.AutoMigrate(&Activity{})
    DB.AutoMigrate(&Message{})
    DB.AutoMigrate(&Picker2Activity{})
}

type User struct {
    //tablename auto-create to be "users"
    ID        int        `gorm:"primary_key"`
    Name      string     `gorm:"size: 20"`
    StdNum    string     `gorm:"size: 11; unique"`

    PostActivities []Activity `gorm:"ForeignKey: PosterID"`
    PickActivities []Activity `gorm:"ForeignKey: PickerID"`
    Messages       []Message  `gorm:"ForeignKey: PickerID"`
}

type Activity struct {
    ID        int        `gorm:"primary_key"`
    Date      time.Time  `gorm:"the activity's date"`
    Time      string     `gorm:"size: 40"`
    Event     string     `gorm:"size: 120"`
    Location  string     `gorm:"size: 60"`
    Tel       string     `gorm:"size: 13"`
    QQ        string     `gorm:"size: 12"`
    Question  string     `gorm:"size: 120"`
    Pickable  bool       `gorm:"default: True"`
    Close     bool       `gorm:"default: False"`

    PosterID  int
    PickerID  int
}

type Message struct {
    //tablename auto-create to be "messages"
    ID        int        `gorm:"primary_key"`
    Time      time.Time  `gorm:"the message creat time"`
    Readed    bool       `gorm:"default: False"`
    Answer    string     `gorm:"size: 120"`

    AID       int
    PickerID  int
}

type Picker2Activity struct {
    ID        int        `gorm:"primary_key"`
    Answer    string     `gorm:"size: 120"`
    Waiting   bool       `gorm:"default: True"`
    Fail      bool       `gorm:"default: True"`

    AID       int
    PickerID  int
}

func (User) TableName() string {
    return "users"
}

func (Activity) TableName() string {
    return "activities"
}

func (Picker2Activity) TableName() string {
    return "picker2activities"
}

func (Message) TableName() string {
    return "messages"
}