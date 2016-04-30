package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerRole struct {
	ID string `gorm:"primary_key"`
}

type Customer struct {
	ID              string `gorm:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	Email           string
	Role            CustomerRole
	RoleId          string
	WantsNewsletter bool
	InNewsletter    bool
}

type WineComment struct {
	ID         string `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Customer   Customer
	CustomerId string
	WineId     string
	WineYear   int
	Comment    string
}

func Initialize(db *gorm.DB) {
	db.SingularTable(true)
	db.Exec(SCHEMA)
}
