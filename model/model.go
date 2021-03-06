package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerRole struct {
	ID string `gorm:"primary_key" json:"id"`
}

type Customer struct {
	ID              string       `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time    `json:"createdAt"`
	UpdatedAt       time.Time    `json:"updatedAt"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Role            CustomerRole `json:"role"`
	RoleId          string       `json:"roleId"`
	WantsNewsletter bool         `json:"wantsNewsletter"`
	InNewsletter    bool         `json:"inNewsletter"`
}

type WineComment struct {
	ID         string    `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Customer   Customer  `json:"customer"`
	CustomerId string    `json:"customerId"`
	WineId     string    `json:"wineId"`
	WineYear   int       `json:"wineYear"`
	Comment    string    `json:"comment"`
}

func Initialize(db *gorm.DB) {
	db.SingularTable(true)
	db.Exec(SCHEMA)
}
