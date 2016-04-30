package main

import (
	"log"
	"strings"

	"github.com/dfreire/df0001/handlers"
	"github.com/dfreire/df0001/middleware"
	"github.com/dfreire/df0001/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
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
	env := viper.Get("ENV").(string)

	db, err := gorm.Open(
		viper.Get("database.dialect").(string),
		viper.Get("database.connectionString").(string),
	)
	if err != nil {
		log.Fatalf("open connection to the database err: %+v", err)
	}

	model.Initialize(db)

	e := echo.New()

	e.Use(echomiddleware.Gzip())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())

	// withDatabase := middleware.WithDatabase(db)
	withTransaction := middleware.WithTransaction(db)
	withErrorLogging := middleware.ErrorLogging()

	e.Post("/signup-customer-with-wine-comments", handlers.SignupCustomerWithWineComments, withErrorLogging, withTransaction)
	e.Post("/signup-customer-with-newsletter", handlers.SignupCustomerWithNewsletter, withErrorLogging, withTransaction)

	if env == "development" {
		db.LogMode(true)
		e.SetDebug(true)
	}

	port := viper.Get("httpServer.port").(string)
	log.Printf("Running on port %s", port)
	e.Run(standard.New(port))
}
