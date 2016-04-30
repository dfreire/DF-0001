package main

import (
	"log"
	"strings"

	"labix.org/v2/mgo/bson"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	env := viper.Get("ENV").(string)

	viper.AddConfigPath(".")
	viper.SetConfigName(strings.Join([]string{"config", env}, "."))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config err: %+v", err)
	}
}

func main() {
	db, err := gorm.Open(
		viper.Get("database.dialect").(string),
		viper.Get("database.connectionString").(string),
	)
	if err != nil {
		log.Fatalf("open connection to the database err: %+v", err)
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
