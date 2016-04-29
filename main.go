package main

import (
	"log"

	"labix.org/v2/mgo/bson"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatalf("error: +%v", err)
	}

	db.SingularTable(true)

	type User struct {
		ID    string `gorm:"primary_key"`
		Name  string
		Email string
	}

	db.AutoMigrate(&User{})

	user := User{
		ID:    bson.NewObjectId().Hex(),
		Name:  "joe.doe",
		Email: "joe.doe@example.com",
	}

	db.Create(&user)
}
