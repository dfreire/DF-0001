package main

import (
	"log"
	"strings"

	"github.com/dfreire/df0001/commands"
	"github.com/dfreire/df0001/model"
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

	db.LogMode(true)
	db.SingularTable(true)

	model.Initialize(db)

	tx := db.Begin()

	err = commands.SignupCustomerWithNewsletter(tx, commands.SignupCustomerWithNewsletterRequestData{
		Name:   "Joe Doe",
		Email:  "joe.doe+1@mailinator.com",
		RoleId: "wine_lover",
	})

	if err != nil {
		tx.Rollback()
	}

	tx.Commit()
}
